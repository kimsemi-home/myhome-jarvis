package daemon

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/auth"
	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
	"github.com/kimsemi-home/myhome-jarvis/internal/qualitylog"
)

func TestIntentReturnsDryRunPlan(t *testing.T) {
	server, err := New(DefaultConfig(t.TempDir(), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodPost, "/intent", bytes.NewBufferString(`{"command":"open-youtube","payload":{}}`))
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	if !bytes.Contains(recorder.Body.Bytes(), []byte(`"dry_run": true`)) {
		t.Fatalf("expected dry-run response, got %s", recorder.Body.String())
	}
}

func TestLinearStatusReturnsRedactedSummary(t *testing.T) {
	server, err := New(DefaultConfig(t.TempDir(), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/linear/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"queue_path": "data/private/linear-offline-queue.jsonl"`,
		`"viewer_configured": false`,
		`"team_count": 0`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{`"token_source"`, `"viewer"`, `"teams"`, `"id"`, `"queue_path": "/"`, `"queue_path": "\\"`} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("linear status leaked %s in %s", forbidden, body)
		}
	}
}

func TestLinearSyncReturnsRedactedSummaryOffline(t *testing.T) {
	server, err := New(DefaultConfig(t.TempDir(), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodPost, "/linear/sync", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"queue_path": "data/private/linear-offline-queue.jsonl"`,
		`"synced": false`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{
		`"issues"`,
		`"issue"`,
		`"description"`,
		`"url"`,
		`"team"`,
		`"id"`,
		`"queue_path": "/"`,
		`"queue_path": "\\"`,
	} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("linear sync leaked %s in %s", forbidden, body)
		}
	}
}

func TestConnectorsStatusReturnsPublicSafeCatalog(t *testing.T) {
	server, err := New(DefaultConfig(repoRoot(t), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/connectors/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"fixture_only": true`,
		`"real_credentials_allowed": false`,
		`"external_api_calls_allowed": false`,
		`"generated_path": "generated/connectors.generated.json"`,
		`"key": "mydata"`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{
		`"token":`,
		`"secret":`,
		`"cookie":`,
		`"account_id":`,
		`"card_number":`,
		`"local_path":`,
		`"generated_path": "/"`,
		`"generated_path": "\\"`,
	} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("connector status leaked %s in %s", forbidden, body)
		}
	}
}

func TestAgentClusterStatusReturnsPublicSafePolicy(t *testing.T) {
	server, err := New(DefaultConfig(repoRoot(t), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/agent-cluster/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"context": "AgentCluster"`,
		`"public_safe": true`,
		`"external_agent_execution_allowed": false`,
		`"self_approval_allowed": false`,
		`"authority_gate_required": true`,
		`"generated_path": "generated/agent_cluster.generated.json"`,
		`"key": "evidence_first"`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{
		`"token":`,
		`"secret":`,
		`"credential":`,
		`"cookie":`,
		`"generated_path": "/"`,
		`"generated_path": "\\"`,
	} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("agent cluster status leaked %s in %s", forbidden, body)
		}
	}
}

func TestLearningStatusReturnsRedactedPrivateLedgerSummary(t *testing.T) {
	root := t.TempDir()
	copyTestFile(t, repoRoot(t), root, "generated/learning.generated.json")
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/learning/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"path": "data/private/learning/observations.jsonl"`,
		`"policy_path": "generated/learning.generated.json"`,
		`"exists": false`,
		`"count": 0`,
		`"open_count": 0`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{
		`"summary":`,
		`"next_action":`,
		`"evidence_refs":`,
		`"token":`,
		`"secret":`,
		`"path": "/"`,
		`"path": "\\"`,
	} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("learning status leaked %s in %s", forbidden, body)
		}
	}
}

func TestIntentExecuteRequiresDaemonExecuteMode(t *testing.T) {
	config := DefaultConfig(t.TempDir(), "test")
	config.CommandPlatform = "darwin"
	config.CommandRunner = func(runnerContext context.Context, invocation commands.Invocation) commands.Execution {
		t.Fatal("runner must not be called when daemon execute mode is disabled")
		return commands.Execution{}
	}
	server, err := New(config)
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodPost, "/intent", bytes.NewBufferString(`{"command":"volume-set","payload":{"level":30},"execute":true}`))
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"dry_run": true`,
		`"execute_allowed": false`,
		`"execute was requested but daemon execute mode is disabled"`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
}

func copyTestFile(t *testing.T, sourceRoot string, targetRoot string, rel string) {
	t.Helper()
	source := filepath.Join(sourceRoot, filepath.FromSlash(rel))
	body, err := os.ReadFile(source)
	if err != nil {
		t.Fatal(err)
	}
	target := filepath.Join(targetRoot, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(target, body, 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestIntentExecutesWithExplicitDaemonGate(t *testing.T) {
	config := DefaultConfig(t.TempDir(), "test")
	config.Execute = true
	config.CommandPlatform = "darwin"
	calls := 0
	config.CommandRunner = func(runnerContext context.Context, invocation commands.Invocation) commands.Execution {
		calls++
		if invocation.Argv[0] != "osascript" {
			t.Fatalf("argv = %#v", invocation.Argv)
		}
		return commands.Execution{Executed: true, ExitCode: 0}
	}
	server, err := New(config)
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodPost, "/intent", bytes.NewBufferString(`{"command":"volume-set","payload":{"level":30},"execute":true}`))
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	if calls != 1 {
		t.Fatalf("runner calls = %d", calls)
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"dry_run": false`,
		`"execute_allowed": true`,
		`"executed": true`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
}

func TestWildcardBindRequiresExplicitAllow(t *testing.T) {
	config := DefaultConfig(t.TempDir(), "test")
	config.Host = "0.0.0.0"
	if _, err := New(config); err == nil {
		t.Fatal("expected wildcard bind to require explicit allow")
	}
}

func TestHTTPServerUsesBoundedResourceTimeouts(t *testing.T) {
	server, err := New(Config{Root: t.TempDir(), Port: 3888, Version: "test"})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := server.httpServer("127.0.0.1:0", server.Routes())

	if httpServer.ReadHeaderTimeout != 5*time.Second {
		t.Fatalf("read header timeout = %s", httpServer.ReadHeaderTimeout)
	}
	if httpServer.ReadTimeout != 15*time.Second {
		t.Fatalf("read timeout = %s", httpServer.ReadTimeout)
	}
	if httpServer.WriteTimeout != 30*time.Second {
		t.Fatalf("write timeout = %s", httpServer.WriteTimeout)
	}
	if httpServer.IdleTimeout != 60*time.Second {
		t.Fatalf("idle timeout = %s", httpServer.IdleTimeout)
	}
	if httpServer.MaxHeaderBytes != 1<<20 {
		t.Fatalf("max header bytes = %d", httpServer.MaxHeaderBytes)
	}
}

func TestMetricsReturnsRuntimeCounters(t *testing.T) {
	server, err := New(DefaultConfig(t.TempDir(), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"requests": 1`,
		`"event_count": 0`,
		`"goroutine_count":`,
		`"heap_alloc_bytes":`,
		`"heap_sys_bytes":`,
		`"stack_inuse_bytes":`,
		`"gc_count":`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	localPathMarker := filepath.Join(string(os.PathSeparator), "Users")
	for _, forbidden := range []string{"root", "token", localPathMarker, "\\"} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("metrics leaked %q in %s", forbidden, body)
		}
	}
}

func TestNonLocalhostRequestsRequireLocalToken(t *testing.T) {
	root := t.TempDir()
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	request.RemoteAddr = "192.168.1.20:4567"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
}

func TestNonLocalhostRequestsAcceptBearerLocalToken(t *testing.T) {
	root := t.TempDir()
	token, err := auth.Create(root, false)
	if err != nil {
		t.Fatal(err)
	}
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	request.RemoteAddr = "192.168.1.20:4567"
	request.Header.Set("Authorization", "Bearer "+token.Token)
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
}

func TestAuthStatusReturnsConfiguredStateWithoutToken(t *testing.T) {
	root := t.TempDir()
	token, err := auth.Create(root, false)
	if err != nil {
		t.Fatal(err)
	}
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/auth/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"configured": true`,
		`"path": "data/private/local-token.txt"`,
		`"mode": "-rw-------"`,
		`"message": "local LAN token is configured"`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	if bytes.Contains([]byte(body), []byte(token.Token)) {
		t.Fatalf("auth status leaked token in %s", body)
	}
}

func TestAuthStatusReturnsMissingState(t *testing.T) {
	server, err := New(DefaultConfig(t.TempDir(), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/auth/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"configured": false`,
		`"path": "data/private/local-token.txt"`,
		`"message": "local LAN token is not configured"`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
}

func TestEventsReturnRecentRequestsWithoutQueryData(t *testing.T) {
	server, err := New(DefaultConfig(t.TempDir(), "test"))
	if err != nil {
		t.Fatal(err)
	}
	routes := server.Routes()
	request := httptest.NewRequest(http.MethodGet, "/health?local_token=redacted-value", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	routes.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	eventsRequest := httptest.NewRequest(http.MethodGet, "/events", nil)
	eventsRequest.RemoteAddr = "127.0.0.1:1234"
	eventsRecorder := httptest.NewRecorder()

	routes.ServeHTTP(eventsRecorder, eventsRequest)

	if eventsRecorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", eventsRecorder.Code, eventsRecorder.Body.String())
	}
	body := eventsRecorder.Body.String()
	for _, expected := range []string{
		`"count": 1`,
		`"method": "GET"`,
		`"path": "/health"`,
		`"status": 200`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{"local_token", "redacted-value"} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("event log leaked query data %q in %s", forbidden, body)
		}
	}
}

func TestEventsKeepBoundedRecentBuffer(t *testing.T) {
	server, err := New(DefaultConfig(t.TempDir(), "test"))
	if err != nil {
		t.Fatal(err)
	}
	routes := server.Routes()
	for i := 0; i < maxRequestEvents+5; i++ {
		request := httptest.NewRequest(http.MethodGet, "/health", nil)
		request.RemoteAddr = "127.0.0.1:1234"
		recorder := httptest.NewRecorder()
		routes.ServeHTTP(recorder, request)
		if recorder.Code != http.StatusOK {
			t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
		}
	}
	eventsRequest := httptest.NewRequest(http.MethodGet, "/events", nil)
	eventsRequest.RemoteAddr = "127.0.0.1:1234"
	eventsRecorder := httptest.NewRecorder()

	routes.ServeHTTP(eventsRecorder, eventsRequest)

	if eventsRecorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", eventsRecorder.Code, eventsRecorder.Body.String())
	}
	if !bytes.Contains(eventsRecorder.Body.Bytes(), []byte(`"count": 100`)) {
		t.Fatalf("expected bounded event count in %s", eventsRecorder.Body.String())
	}
}

func TestEventsRecordHandlerErrors(t *testing.T) {
	server, err := New(DefaultConfig(t.TempDir(), "test"))
	if err != nil {
		t.Fatal(err)
	}
	routes := server.Routes()
	request := httptest.NewRequest(http.MethodPost, "/intent", bytes.NewBufferString(`{"command":"unknown","payload":{}}`))
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	routes.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	eventsRequest := httptest.NewRequest(http.MethodGet, "/events", nil)
	eventsRequest.RemoteAddr = "127.0.0.1:1234"
	eventsRecorder := httptest.NewRecorder()

	routes.ServeHTTP(eventsRecorder, eventsRequest)

	if eventsRecorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", eventsRecorder.Code, eventsRecorder.Body.String())
	}
	body := eventsRecorder.Body.String()
	for _, expected := range []string{
		`"path": "/intent"`,
		`"status": 400`,
		`"error": "bad_request"`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
}

func TestSupervisorStatusReturnsPrivateStatePath(t *testing.T) {
	server, err := New(DefaultConfig(t.TempDir(), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/supervisor/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"recorded": false`,
		`"state_path": "data/private/supervisor/daemon-state.json"`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
}

func TestIntentWritesRedactedAuditJournal(t *testing.T) {
	root := t.TempDir()
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodPost, "/intent", bytes.NewBufferString(`{"command":"open-url","payload":{"url":"https://example.invalid/private"},"execute":false}`))
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	statusRequest := httptest.NewRequest(http.MethodGet, "/audit/status", nil)
	statusRequest.RemoteAddr = "127.0.0.1:1234"
	statusRecorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(statusRecorder, statusRequest)

	if statusRecorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", statusRecorder.Code, statusRecorder.Body.String())
	}
	statusBody := statusRecorder.Body.String()
	for _, expected := range []string{
		`"path": "data/private/audit/command-intents.jsonl"`,
		`"count": 1`,
		`"command": "open_url"`,
		`"source": "daemon"`,
		`"success": true`,
	} {
		if !bytes.Contains([]byte(statusBody), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, statusBody)
		}
	}
	data, err := os.ReadFile(filepath.Join(root, "data", "private", "audit", "command-intents.jsonl"))
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{"payload", "argv", "https://example.invalid/private"} {
		if bytes.Contains(data, []byte(forbidden)) {
			t.Fatalf("audit journal leaked %q in %s", forbidden, data)
		}
	}
}

func TestIntentAuditRecordsCommandErrors(t *testing.T) {
	server, err := New(DefaultConfig(t.TempDir(), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodPost, "/intent", bytes.NewBufferString(`{"command":"unknown","payload":{}}`))
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	statusRequest := httptest.NewRequest(http.MethodGet, "/audit/status", nil)
	statusRequest.RemoteAddr = "127.0.0.1:1234"
	statusRecorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(statusRecorder, statusRequest)

	if statusRecorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", statusRecorder.Code, statusRecorder.Body.String())
	}
	statusBody := statusRecorder.Body.String()
	for _, expected := range []string{
		`"count": 1`,
		`"command": "unknown"`,
		`"success": false`,
		`"error_category": "unknown_command"`,
	} {
		if !bytes.Contains([]byte(statusBody), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, statusBody)
		}
	}
}

func TestQualityStatusReturnsPrivateEvidencePath(t *testing.T) {
	server, err := New(DefaultConfig(t.TempDir(), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/quality/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"path": "data/private/quality/runs.jsonl"`,
		`"exists": false`,
		`"count": 0`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
}

func TestQualityStatusReturnsLastPrivateRun(t *testing.T) {
	root := t.TempDir()
	if err := qualitylog.AppendRun(root, qualitylog.NewRun(time.Now(), true, []qualitylog.Step{
		{Name: "go test", Status: "pass"},
		{Name: "flutter analyze", Status: "pass"},
	})); err != nil {
		t.Fatal(err)
	}
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/quality/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"exists": true`,
		`"count": 1`,
		`"ok": true`,
		`"pass_count": 2`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
}

func TestDomainSummaryReturnsFixtureStatus(t *testing.T) {
	root := repoRoot(t)
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/domain/summary", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"net_minor_units": 4346800`,
		`"recurring_candidate_count": 1`,
		`"recommendations"`,
		`"household"`,
		`"long_term_format": "parquet"`,
		`"private_root": "data/lake"`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	if bytes.Contains([]byte(body), []byte(root)) {
		t.Fatalf("domain summary leaked local root in %s", body)
	}
	home, err := os.UserHomeDir()
	if err == nil && home != "" && bytes.Contains([]byte(body), []byte(home)) {
		t.Fatalf("domain summary leaked local home in %s", body)
	}
	for _, forbidden := range []string{`"private_root": "/"`, `"private_root": "\\"`, `"private_root": "../"`} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("domain summary leaked unsafe storage root %s in %s", forbidden, body)
		}
	}
}

func TestHarnessRunReturnsFinanceAndCommerceReports(t *testing.T) {
	server, err := New(DefaultConfig(repoRoot(t), "test"))
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name     string
		expected string
	}{
		{name: "finance", expected: `"name": "finance"`},
		{name: "commerce", expected: `"name": "commerce"`},
	} {
		request := httptest.NewRequest(http.MethodPost, "/harness/run", bytes.NewBufferString(`{"name":"`+tc.name+`"}`))
		request.RemoteAddr = "127.0.0.1:1234"
		recorder := httptest.NewRecorder()

		server.Routes().ServeHTTP(recorder, request)

		if recorder.Code != http.StatusOK {
			t.Fatalf("%s status = %d body = %s", tc.name, recorder.Code, recorder.Body.String())
		}
		body := recorder.Body.String()
		for _, expected := range []string{
			tc.expected,
			`"passed": true`,
		} {
			if !bytes.Contains([]byte(body), []byte(expected)) {
				t.Fatalf("expected %s in %s", expected, body)
			}
		}
	}
}

func TestLoopStatusReturnsSchedulerPolicy(t *testing.T) {
	server, err := New(DefaultConfig(repoRoot(t), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/loop/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"name": "closed_loop"`,
		`"heartbeat_interval_seconds": 15`,
		`"min_backoff_seconds": 5`,
		`"checkpoint_every": 1`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
}

func TestPlannerStatusReturnsGeneratedTaskGraph(t *testing.T) {
	server, err := New(DefaultConfig(repoRoot(t), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/planner/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"task_count": 6`,
		`"ready_count": 0`,
		`"completed_count": 5`,
		`"blocked_external_write_count": 1`,
		`"blocked_external_write_tasks": [`,
		`"id": "linear_sync"`,
		`"external_write_gate": {`,
		`"mutation_success_required": true`,
		`"linear_write_evidence": {`,
		`"evidence_path": "data/private/linear-write-evidence.jsonl"`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	if bytes.Contains([]byte(body), []byte(`"next_task"`)) {
		t.Fatalf("unexpected next task in %s", body)
	}
}

func TestRepoStatusReturnsGitWorktreeState(t *testing.T) {
	root := initTempRepo(t)
	if err := os.WriteFile(filepath.Join(root, "tracked.txt"), []byte("changed\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "new.txt"), []byte("new\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/repo/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"branch": "main"`,
		`"worktree_clean": false`,
		`"path": "tracked.txt"`,
		`"new.txt"`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
}

func TestSecurityStatusReturnsRedactedPublicSafetyState(t *testing.T) {
	server, err := New(DefaultConfig(repoRoot(t), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/security/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"ok": true`,
		`"current_ok": true`,
		`"history_ok": true`,
		`"current_finding_count": 0`,
		`"history_finding_count": 0`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{"findings", "root"} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("security status leaked %q in %s", forbidden, body)
		}
	}
}

func TestHouseholdSummaryReturnsScopeSwitchingData(t *testing.T) {
	server, err := New(DefaultConfig(repoRoot(t), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/household/summary", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"scope": "user"`,
		`"scope": "spouse"`,
		`"scope": "household"`,
		`"finance_net_minor_units": 4346800`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
}

func initTempRepo(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	runGit(t, root, "init", "-b", "main")
	runGit(t, root, "config", "user.name", "Test User")
	runGit(t, root, "config", "user.email", "test@example.invalid")
	if err := os.WriteFile(filepath.Join(root, "tracked.txt"), []byte("initial\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	runGit(t, root, "add", "tracked.txt")
	runGit(t, root, "commit", "-m", "initial")
	return root
}

func runGit(t *testing.T, root string, args ...string) {
	t.Helper()
	command := append([]string{"-C", root}, args...)
	cmd := exec.Command("git", command...)
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git %v failed: %v\n%s", args, err, output)
	}
}

func TestRecommendationsSummaryReturnsLocalFixtureRecommendations(t *testing.T) {
	server, err := New(DefaultConfig(repoRoot(t), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/recommendations/summary", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"count": 4`,
		`"kind": "recurring_purchase_review"`,
		`"kind": "card_usage_review"`,
		`"kind": "subscription_review"`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
}

func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		next := filepath.Dir(dir)
		if next == dir {
			t.Fatal("could not find repo root")
		}
		dir = next
	}
}
