package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCodexCostStatusReturnsRedactedUsageSummary(t *testing.T) {
	root := t.TempDir()
	copyTestFile(t, repoRoot(t), root, "generated/codex_cost.generated.json")
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/codex-cost/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"policy_path": "generated/codex_cost.generated.json"`,
		`"ledger_path": "data/private/codex-cost/usage.jsonl"`,
		`"exists": false`,
		`"record_count": 0`,
		`"budget_state": "ok"`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{
		`"raw_prompt":`,
		`"raw_transcript":`,
		`"private_notes":`,
		`"evidence_refs":`,
		`"token":`,
		`"secret":`,
		`"ledger_path": "/"`,
	} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("codex cost status leaked %s in %s", forbidden, body)
		}
	}
}
