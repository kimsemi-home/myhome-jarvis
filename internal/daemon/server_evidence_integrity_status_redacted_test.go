package daemon

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEvidenceIntegrityStatusReturnsRedactedPrefixCounts(t *testing.T) {
	root := t.TempDir()
	copyTestFile(t, repoRoot(t), root, "generated/evidence.generated.json")
	writeDaemonTestFile(t, root, "data/private/learning/observations.jsonl", `{"id":"learn_1","at":"2026-06-18T00:00:00Z","status":"open","summary":"private detail","evidence_refs":["generated/missing.generated.json"],"next_action":"private action"}`+"\n")
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/evidence-integrity/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{`"dangling_evidence_ref_count": 1`, `"prefix": "generated/"`} {
		if !strings.Contains(body, expected) {
			t.Fatalf("missing %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{"generated/missing.generated.json", "private detail", "private action", "evidence_refs"} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("evidence integrity status leaked %s in %s", forbidden, body)
		}
	}
}
