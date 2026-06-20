package daemon

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWorkItemStatusReturnsRedactedCard(t *testing.T) {
	server, err := New(DefaultConfig(repoRoot(t), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/work-item/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"context": "UniversalWorkItem"`,
		`"work_item_state":`,
		`"approval_state": "not_approved"`,
		`"external_writes_allowed": false`,
	} {
		if !strings.Contains(body, expected) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{"raw_prompt", "linear_url", "credential"} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("work item leaked %s in %s", forbidden, body)
		}
	}
}
