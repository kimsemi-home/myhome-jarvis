package daemon

import (
	"bytes"
	"github.com/kimsemi-home/myhome-jarvis/internal/auth"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthStatusReturnsConfiguredStateWithoutToken(t *testing.T) {
	root := t.TempDir()
	token, err := auth.Create(root, false)
	if err != nil {
		t.Fatal(err)
	}
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/auth/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"configured": true`,
		`"path": "data/private/local-token.txt"`,
		`"mode": "-rw-------"`,
		`"message": "local LAN token is configured"`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	if bytes.Contains([]byte(body), []byte(token.Token)) {
		t.Fatalf("auth status leaked token in %s", body)
	}
}
