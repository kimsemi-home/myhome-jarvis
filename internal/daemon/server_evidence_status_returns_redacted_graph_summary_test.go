package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEvidenceStatusReturnsRedactedGraphSummary(t *testing.T) {
	root := t.TempDir()
	copyTestFile(t, repoRoot(t), root, "generated/evidence.generated.json")
	writeDaemonTestFile(t, root, "data/private/learning/observations.jsonl", `{"id":"learn_1","at":"2026-06-18T00:00:00Z","kind":"loop_gap","source":"quality_gate","stage":"evidence_recorded","status":"open","summary":"private observation","evidence_refs":["data/private/quality/runs.jsonl"],"owner":"go","next_action":"private action"}`+"\n")
	writeDaemonTestFile(t, root, "data/private/quality/runs.jsonl", `{"at":"2026-06-18T00:01:00Z","ok":true}`+"\n")
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/evidence/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"policy_path": "generated/evidence.generated.json"`,
		`"private_root": "data/private"`,
		`"node_count": 3`,
		`"edge_count": 1`,
		`"open_learning_count": 1`,
		`"key": "learning"`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{
		`"summary":`,
		`"next_action":`,
		`"evidence_refs":`,
		`"private observation"`,
		`"private action"`,
		`"data/private/quality/runs.jsonl"`,
		`"token":`,
		`"secret":`,
		`"private_root": "/"`,
		`"private_root": "\\"`,
	} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("evidence status leaked %s in %s", forbidden, body)
		}
	}
}
