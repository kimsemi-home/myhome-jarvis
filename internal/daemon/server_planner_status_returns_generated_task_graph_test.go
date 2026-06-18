package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
