package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEventsKeepBoundedRecentBuffer(t *testing.T) {
	server, err := New(DefaultConfig(t.TempDir(), "test"))
	if err != nil {
		t.Fatal(err)
	}
	routes := server.Routes()
	for i := 0; i < maxRequestEvents+5; i++ {
		request := httptest.NewRequest(http.MethodGet, "/health", nil)
		request.RemoteAddr = "127.0.0.1:1234"
		recorder := httptest.NewRecorder()
		routes.ServeHTTP(recorder, request)
		if recorder.Code != http.StatusOK {
			t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
		}
	}
	eventsRequest := httptest.NewRequest(http.MethodGet, "/events", nil)
	eventsRequest.RemoteAddr = "127.0.0.1:1234"
	eventsRecorder := httptest.NewRecorder()

	routes.ServeHTTP(eventsRecorder, eventsRequest)

	if eventsRecorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", eventsRecorder.Code, eventsRecorder.Body.String())
	}
	if !bytes.Contains(eventsRecorder.Body.Bytes(), []byte(`"count": 100`)) {
		t.Fatalf("expected bounded event count in %s", eventsRecorder.Body.String())
	}
}
