package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAssistantStatusReturnsRedactedCommandCenter(t *testing.T) {
	server, err := New(DefaultConfig(repoRoot(t), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/assistant/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"context": "AssistantCommandCenter"`,
		`"policy_path": "generated/assistant_vision.generated.json"`,
		`"ready_pillar_count":`,
		`"gated_pillar_keys":`,
		`"blocked_gate_count":`,
		`"work_item":`,
		`"next_safe_action":`,
		`"compact_state":`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{"raw_prompt", "raw_transcript", "secret", "local_absolute_path"} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("assistant status leaked %s in %s", forbidden, body)
		}
	}
}
