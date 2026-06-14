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
	server, err := New(DefaultConfig(repoRoot(t), "test"))
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
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
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
		`"ready_count": 5`,
		`"blocked_external_write_count": 1`,
		`"id": "repo_safety"`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
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
