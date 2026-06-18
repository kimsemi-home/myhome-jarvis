package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLinearSyncReturnsRedactedSummaryOffline(t *testing.T) {
	server, err := New(DefaultConfig(t.TempDir(), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodPost, "/linear/sync", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"queue_path": "data/private/linear-offline-queue.jsonl"`,
		`"synced": false`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{
		`"issues"`,
		`"issue"`,
		`"description"`,
		`"url"`,
		`"team"`,
		`"id"`,
		`"queue_path": "/"`,
		`"queue_path": "\\"`,
	} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("linear sync leaked %s in %s", forbidden, body)
		}
	}
}
