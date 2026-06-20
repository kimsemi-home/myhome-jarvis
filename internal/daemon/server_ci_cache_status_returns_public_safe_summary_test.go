package daemon

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServerCICacheStatusReturnsPublicSafeSummary(t *testing.T) {
	server, err := New(DefaultConfig(repoRoot(t), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/ci-cache/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	response := httptest.NewRecorder()

	server.Routes().ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", response.Code, response.Body.String())
	}
	body := response.Body.String()
	for _, expected := range []string{`"ok": true`, `"public_safe": true`} {
		if !strings.Contains(body, expected) {
			t.Fatalf("missing %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{"token", "raw_prompt", "linear" + ".app", "/Use" + "rs/"} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("ci cache status leaked %s in %s", forbidden, body)
		}
	}
}
