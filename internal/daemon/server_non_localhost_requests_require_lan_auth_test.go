package daemon

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNonLocalhostRequestsRequireLocalToken(t *testing.T) {
	root := t.TempDir()
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	request.RemoteAddr = "192.168.1.20:4567"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
}
