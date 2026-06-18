package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIncidentsStatusReturnsRedactedLifecycleSummary(t *testing.T) {
	root := t.TempDir()
	copyTestFile(t, repoRoot(t), root, "generated/incidents.generated.json")
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/incidents/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"policy_path": "generated/incidents.generated.json"`,
		`"ledger_path": "data/private/incidents/incidents.jsonl"`,
		`"exists": false`,
		`"count": 0`,
		`"incident_debt_count": 0`,
		`"quarantine_stale_after_hours": 168`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{
		`"summary":`,
		`"root_cause":`,
		`"evidence_refs":`,
		`"raw_prompt":`,
		`"raw_transcript":`,
		`"token":`,
		`"secret":`,
		`"credential":`,
		`"ledger_path": "/"`,
	} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("incidents status leaked %s in %s", forbidden, body)
		}
	}
}
