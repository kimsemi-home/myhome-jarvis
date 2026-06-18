package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEventsReturnRecentRequestsWithoutQueryData(t *testing.T) {
	server, err := New(DefaultConfig(t.TempDir(), "test"))
	if err != nil {
		t.Fatal(err)
	}
	routes := server.Routes()
	request := httptest.NewRequest(http.MethodGet, "/health?local_token=redacted-value", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	routes.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	eventsRequest := httptest.NewRequest(http.MethodGet, "/events", nil)
	eventsRequest.RemoteAddr = "127.0.0.1:1234"
	eventsRecorder := httptest.NewRecorder()

	routes.ServeHTTP(eventsRecorder, eventsRequest)

	if eventsRecorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", eventsRecorder.Code, eventsRecorder.Body.String())
	}
	body := eventsRecorder.Body.String()
	for _, expected := range []string{
		`"count": 1`,
		`"method": "GET"`,
		`"path": "/health"`,
		`"status": 200`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{"local_token", "redacted-value"} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("event log leaked query data %q in %s", forbidden, body)
		}
	}
}
