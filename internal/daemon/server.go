package daemon

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/agentcluster"
	"github.com/kimsemi-home/myhome-jarvis/internal/audit"
	"github.com/kimsemi-home/myhome-jarvis/internal/auth"
	"github.com/kimsemi-home/myhome-jarvis/internal/authority"
	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
	"github.com/kimsemi-home/myhome-jarvis/internal/confidence"
	"github.com/kimsemi-home/myhome-jarvis/internal/connectors"
	"github.com/kimsemi-home/myhome-jarvis/internal/controlplane"
	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
	"github.com/kimsemi-home/myhome-jarvis/internal/evidence"
	"github.com/kimsemi-home/myhome-jarvis/internal/evidencequality"
	"github.com/kimsemi-home/myhome-jarvis/internal/incidents"
	"github.com/kimsemi-home/myhome-jarvis/internal/learning"
	"github.com/kimsemi-home/myhome-jarvis/internal/linear"
	"github.com/kimsemi-home/myhome-jarvis/internal/planner"
	"github.com/kimsemi-home/myhome-jarvis/internal/qualitylog"
	"github.com/kimsemi-home/myhome-jarvis/internal/repo"
	"github.com/kimsemi-home/myhome-jarvis/internal/scheduler"
	"github.com/kimsemi-home/myhome-jarvis/internal/security"
	"github.com/kimsemi-home/myhome-jarvis/internal/supervisor"
	"github.com/kimsemi-home/myhome-jarvis/internal/translation"
)

const (
	defaultReadHeaderTimeout = 5 * time.Second
	defaultReadTimeout       = 15 * time.Second
	defaultWriteTimeout      = 30 * time.Second
	defaultIdleTimeout       = 60 * time.Second
	defaultMaxHeaderBytes    = 1 << 20
)

type Config struct {
	Root              string
	Host              string
	Port              int
	Execute           bool
	AllowLANBind      bool
	Version           string
	CommandPlatform   string
	CommandRunner     commands.Runner
	ReadHeaderTimeout time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	MaxHeaderBytes    int
}

type Server struct {
	config   Config
	started  time.Time
	requests atomic.Uint64
	events   *eventLog
}

func DefaultConfig(root string, version string) Config {
	config := Config{
		Root:    root,
		Host:    envString("MYHOME_BIND_HOST", "127.0.0.1"),
		Port:    envInt("MYHOME_BIND_PORT", 3888),
		Execute: os.Getenv("MYHOME_EXECUTE") == "true",
		Version: version,
	}
	return withDefaultResourceBounds(config)
}

func withDefaultResourceBounds(config Config) Config {
	if config.ReadHeaderTimeout <= 0 {
		config.ReadHeaderTimeout = defaultReadHeaderTimeout
	}
	if config.ReadTimeout <= 0 {
		config.ReadTimeout = defaultReadTimeout
	}
	if config.WriteTimeout <= 0 {
		config.WriteTimeout = defaultWriteTimeout
	}
	if config.IdleTimeout <= 0 {
		config.IdleTimeout = defaultIdleTimeout
	}
	if config.MaxHeaderBytes <= 0 {
		config.MaxHeaderBytes = defaultMaxHeaderBytes
	}
	return config
}

func New(config Config) (*Server, error) {
	config = withDefaultResourceBounds(config)
	if config.Root == "" {
		return nil, errors.New("root is required")
	}
	if config.Host == "" {
		config.Host = "127.0.0.1"
	}
	if config.Port <= 0 || config.Port > 65535 {
		return nil, fmt.Errorf("invalid port %d", config.Port)
	}
	if isWildcardHost(config.Host) && !config.AllowLANBind {
		return nil, errors.New("wildcard or LAN bind requires explicit allow-lan flag")
	}
	return &Server{config: config, started: time.Now().UTC(), events: newEventLog(maxRequestEvents)}, nil
}

func (server *Server) ListenAndServe() error {
	mux := server.Routes()
	address := net.JoinHostPort(server.config.Host, strconv.Itoa(server.config.Port))
	httpServer := server.httpServer(address, mux)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	state, err := supervisor.NewDaemonState(
		server.config.Root,
		server.config.Host,
		server.config.Port,
		server.config.Version,
		server.config.Execute,
		server.config.AllowLANBind,
	)
	if err != nil {
		_ = listener.Close()
		return err
	}
	if _, err := supervisor.WriteDaemonState(server.config.Root, state); err != nil {
		_ = listener.Close()
		return err
	}
	return httpServer.Serve(listener)
}

func (server *Server) httpServer(address string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              address,
		Handler:           handler,
		ReadHeaderTimeout: server.config.ReadHeaderTimeout,
		ReadTimeout:       server.config.ReadTimeout,
		WriteTimeout:      server.config.WriteTimeout,
		IdleTimeout:       server.config.IdleTimeout,
		MaxHeaderBytes:    server.config.MaxHeaderBytes,
	}
}

func (server *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", server.wrap(server.handleHealth))
	mux.HandleFunc("GET /version", server.wrap(server.handleVersion))
	mux.HandleFunc("GET /auth/status", server.wrap(server.handleAuthStatus))
	mux.HandleFunc("GET /commands", server.wrap(server.handleCommands))
	mux.HandleFunc("POST /intent", server.wrap(server.handleIntent))
	mux.HandleFunc("POST /harness/run", server.wrap(server.handleHarnessRun))
	mux.HandleFunc("GET /linear/status", server.wrap(server.handleLinearStatus))
	mux.HandleFunc("POST /linear/sync", server.wrap(server.handleLinearSync))
	mux.HandleFunc("GET /repo/status", server.wrap(server.handleRepoStatus))
	mux.HandleFunc("GET /security/status", server.wrap(server.handleSecurityStatus))
	mux.HandleFunc("GET /loop/status", server.wrap(server.handleLoopStatus))
	mux.HandleFunc("GET /domain/summary", server.wrap(server.handleDomainSummary))
	mux.HandleFunc("GET /connectors/status", server.wrap(server.handleConnectorsStatus))
	mux.HandleFunc("GET /agent-cluster/status", server.wrap(server.handleAgentClusterStatus))
	mux.HandleFunc("GET /learning/status", server.wrap(server.handleLearningStatus))
	mux.HandleFunc("GET /evidence/status", server.wrap(server.handleEvidenceStatus))
	mux.HandleFunc("GET /confidence/status", server.wrap(server.handleConfidenceStatus))
	mux.HandleFunc("GET /translation/status", server.wrap(server.handleTranslationStatus))
	mux.HandleFunc("GET /control-plane/status", server.wrap(server.handleControlPlaneStatus))
	mux.HandleFunc("GET /incidents/status", server.wrap(server.handleIncidentsStatus))
	mux.HandleFunc("GET /evidence-quality/status", server.wrap(server.handleEvidenceQualityStatus))
	mux.HandleFunc("GET /authority/status", server.wrap(server.handleAuthorityStatus))
	mux.HandleFunc("GET /household/summary", server.wrap(server.handleHouseholdSummary))
	mux.HandleFunc("GET /recommendations/summary", server.wrap(server.handleRecommendationsSummary))
	mux.HandleFunc("GET /metrics", server.wrap(server.handleMetrics))
	mux.HandleFunc("GET /events", server.wrap(server.handleEvents))
	mux.HandleFunc("GET /supervisor/status", server.wrap(server.handleSupervisorStatus))
	mux.HandleFunc("GET /audit/status", server.wrap(server.handleAuditStatus))
	mux.HandleFunc("GET /quality/status", server.wrap(server.handleQualityStatus))
	mux.HandleFunc("GET /planner/status", server.wrap(server.handlePlannerStatus))
	return mux
}

type handlerFunc func(http.ResponseWriter, *http.Request) error

func (server *Server) wrap(next handlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		started := time.Now()
		recorder := &statusRecorder{ResponseWriter: writer}
		var handlerErr error
		defer func() {
			server.recordRequestEvent(request, recorder.statusCode(), started, handlerErr)
		}()

		server.requests.Add(1)
		if err := server.authorize(request); err != nil {
			handlerErr = err
			writeError(recorder, http.StatusUnauthorized, err)
			return
		}
		if err := next(recorder, request); err != nil {
			handlerErr = err
			writeError(recorder, http.StatusBadRequest, err)
		}
	}
}

func (server *Server) authorize(request *http.Request) error {
	host, _, err := net.SplitHostPort(request.RemoteAddr)
	if err != nil {
		host = request.RemoteAddr
	}
	parsed := net.ParseIP(host)
	if parsed == nil || parsed.IsLoopback() {
		return nil
	}
	token, err := readLocalToken(server.config.Root)
	if err != nil {
		return errors.New("local token is required for non-localhost clients")
	}
	header := strings.TrimSpace(request.Header.Get("Authorization"))
	if header != "Bearer "+token {
		return errors.New("invalid local token")
	}
	return nil
}

func (server *Server) handleHealth(writer http.ResponseWriter, request *http.Request) error {
	return writeJSON(writer, http.StatusOK, map[string]any{
		"ok":       true,
		"mode":     "local",
		"dry_run":  true,
		"host":     server.config.Host,
		"started":  server.started.Format(time.RFC3339),
		"requests": server.requests.Load(),
	})
}

func (server *Server) handleVersion(writer http.ResponseWriter, request *http.Request) error {
	return writeJSON(writer, http.StatusOK, map[string]any{
		"name":    "myhome-jarvis",
		"version": server.config.Version,
	})
}

func (server *Server) handleAuthStatus(writer http.ResponseWriter, request *http.Request) error {
	return writeJSON(writer, http.StatusOK, auth.Status(server.config.Root))
}

func (server *Server) handleCommands(writer http.ResponseWriter, request *http.Request) error {
	return writeJSON(writer, http.StatusOK, commands.Specs())
}

type intentRequest struct {
	Command string          `json:"command"`
	Payload json.RawMessage `json:"payload"`
	Execute bool            `json:"execute"`
}

func (server *Server) handleIntent(writer http.ResponseWriter, request *http.Request) error {
	var body intentRequest
	if err := decodeBody(request, &body); err != nil {
		_ = audit.AppendCommandIntent(server.config.Root, audit.CommandIntentFromPlan("daemon", "", false, commands.Plan{}, err))
		return err
	}
	if len(body.Payload) == 0 {
		body.Payload = []byte("{}")
	}
	plan, err := commands.Build(body.Command, body.Payload)
	if err != nil {
		_ = audit.AppendCommandIntent(server.config.Root, audit.CommandIntentFromPlan("daemon", body.Command, body.Execute, plan, err))
		return err
	}
	executeAllowed := body.Execute && server.config.Execute
	plan = commands.WithExecuteAllowed(plan, executeAllowed)
	if body.Execute && !server.config.Execute {
		plan.Warnings = append(plan.Warnings, "execute was requested but daemon execute mode is disabled")
	}
	if executeAllowed {
		var err error
		plan, err = commands.Execute(request.Context(), plan, commands.ExecuteOptions{
			Platform: server.config.CommandPlatform,
			Runner:   server.config.CommandRunner,
		})
		if err != nil {
			_ = audit.AppendCommandIntent(server.config.Root, audit.CommandIntentFromPlan("daemon", body.Command, body.Execute, plan, err))
			return err
		}
	}
	if err := audit.AppendCommandIntent(server.config.Root, audit.CommandIntentFromPlan("daemon", body.Command, body.Execute, plan, nil)); err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, plan)
}

type harnessRequest struct {
	Name string `json:"name"`
}

func (server *Server) handleHarnessRun(writer http.ResponseWriter, request *http.Request) error {
	var body harnessRequest
	if err := decodeBody(request, &body); err != nil {
		return err
	}
	switch strings.TrimSpace(strings.ToLower(body.Name)) {
	case "", "home":
		return server.writeHarnessReport(writer, commands.RunHomeHarness())
	case "finance":
		return server.writeHarnessReport(writer, commands.RunFinanceHarness(server.config.Root))
	case "commerce":
		return server.writeHarnessReport(writer, commands.RunCommerceHarness(server.config.Root))
	default:
		return fmt.Errorf("unknown harness %q", body.Name)
	}
}

func (server *Server) writeHarnessReport(writer http.ResponseWriter, report commands.HarnessReport) error {
	status := http.StatusOK
	if !report.Passed {
		status = http.StatusBadRequest
	}
	return writeJSON(writer, status, report)
}

func (server *Server) handleLinearStatus(writer http.ResponseWriter, request *http.Request) error {
	status := linear.StatusForRequest(request.Context(), server.config.Root, http.DefaultClient)
	return writeJSON(writer, http.StatusOK, linear.SummarizeStatus(status))
}

func (server *Server) handleLinearSync(writer http.ResponseWriter, request *http.Request) error {
	result := linear.PullIssues(request.Context(), server.config.Root, http.DefaultClient)
	if !result.Synced {
		if err := linear.AppendOfflineEvent(server.config.Root, "linear_sync", result.Message); err != nil {
			return err
		}
	}
	return writeJSON(writer, http.StatusOK, linear.SummarizeOperation(result))
}

func (server *Server) handleRepoStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := repo.Inspect(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleSecurityStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := security.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleLoopStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := scheduler.Status(server.config.Root, scheduler.ClosedLoopPolicy())
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleDomainSummary(writer http.ResponseWriter, request *http.Request) error {
	summary, err := domain.BuildSummary(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, summary)
}

func (server *Server) handleConnectorsStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := connectors.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleAgentClusterStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := agentcluster.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleLearningStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := learning.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleEvidenceStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := evidence.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleConfidenceStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := confidence.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleTranslationStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := translation.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleControlPlaneStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := controlplane.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleIncidentsStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := incidents.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleEvidenceQualityStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := evidencequality.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleAuthorityStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := authority.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleHouseholdSummary(writer http.ResponseWriter, request *http.Request) error {
	summary, err := domain.BuildSummary(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, summary.Household)
}

func (server *Server) handleRecommendationsSummary(writer http.ResponseWriter, request *http.Request) error {
	summary, err := domain.BuildSummary(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, summary.Recommendations)
}

func (server *Server) handleMetrics(writer http.ResponseWriter, request *http.Request) error {
	events := server.events.snapshot()
	var memory runtime.MemStats
	runtime.ReadMemStats(&memory)
	return writeJSON(writer, http.StatusOK, map[string]any{
		"started":           server.started.Format(time.RFC3339),
		"uptime_seconds":    int64(time.Since(server.started).Seconds()),
		"requests":          server.requests.Load(),
		"event_count":       len(events),
		"goroutine_count":   runtime.NumGoroutine(),
		"heap_alloc_bytes":  memory.HeapAlloc,
		"heap_sys_bytes":    memory.HeapSys,
		"stack_inuse_bytes": memory.StackInuse,
		"gc_count":          memory.NumGC,
		"execute_enabled":   server.config.Execute,
		"bind_host":         server.config.Host,
		"linear_queue":      filepath.ToSlash(filepath.Join("data", "private", "linear-offline-queue.jsonl")),
		"dry_run_default":   true,
		"lan_bind_allowed":  server.config.AllowLANBind,
	})
}

func (server *Server) handleEvents(writer http.ResponseWriter, request *http.Request) error {
	events := server.events.snapshot()
	return writeJSON(writer, http.StatusOK, map[string]any{
		"count":  len(events),
		"events": events,
	})
}

func (server *Server) handleSupervisorStatus(writer http.ResponseWriter, request *http.Request) error {
	return writeJSON(writer, http.StatusOK, supervisor.Status(server.config.Root, nil))
}

func (server *Server) handleAuditStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := audit.CommandIntentStatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handleQualityStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := qualitylog.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func (server *Server) handlePlannerStatus(writer http.ResponseWriter, request *http.Request) error {
	status, err := planner.StatusForRoot(server.config.Root)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, status)
}

func decodeBody(request *http.Request, target any) error {
	defer request.Body.Close()
	decoder := json.NewDecoder(io.LimitReader(request.Body, 1<<20))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return fmt.Errorf("invalid json body: %w", err)
	}
	return nil
}

func writeJSON(writer http.ResponseWriter, status int, value any) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	return encoder.Encode(value)
}

func writeError(writer http.ResponseWriter, status int, err error) {
	_ = writeJSON(writer, status, map[string]any{
		"ok":    false,
		"error": err.Error(),
	})
}

func readLocalToken(root string) (string, error) {
	return auth.Read(root)
}

func envString(name string, fallback string) string {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return fallback
	}
	return value
}

func envInt(name string, fallback int) int {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}

func isWildcardHost(host string) bool {
	return host == "0.0.0.0" || host == "::" || host == ""
}
