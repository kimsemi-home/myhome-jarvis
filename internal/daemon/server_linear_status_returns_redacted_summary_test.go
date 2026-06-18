package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
