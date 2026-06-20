package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMonetizationStatusReturnsRedactedSummary(t *testing.T) {
	root := t.TempDir()
	copyTestFile(t, repoRoot(t), root, "generated/monetization.generated.json")
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/monetization/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"policy_path": "generated/monetization.generated.json"`,
		`"ledger_path": "data/private/monetization/experiments.jsonl"`,
		`"experiment_count": 0`,
		`"decision_count": 0`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{"private_revenue_notes", "evidence_refs", "secret"} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("monetization status leaked %s in %s", forbidden, body)
		}
	}
}
