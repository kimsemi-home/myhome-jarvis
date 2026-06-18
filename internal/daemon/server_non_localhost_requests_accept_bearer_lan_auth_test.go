package daemon

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/auth"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNonLocalhostRequestsAcceptBearerLocalToken(t *testing.T) {
	root := t.TempDir()
	token, err := auth.Create(root, false)
	if err != nil {
		t.Fatal(err)
	}
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	request.RemoteAddr = "192.168.1.20:4567"
	request.Header.Set("Authorization", "Bearer "+token.Token)
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
}
