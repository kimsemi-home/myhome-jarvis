package daemon

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMediaReadinessStatusReturnsRedactedBenchmark(t *testing.T) {
	root := t.TempDir()
	copyTestFile(t, repoRoot(t), root, "generated/media_readiness.generated.json")
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/media-readiness/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{`"context": "MediaReadinessBenchmark"`, `"case_count": 3`} {
		if !strings.Contains(body, expected) {
			t.Fatalf("missing %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{`"payload"`, `"query"`, "cookie", "account_id"} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("media readiness status leaked %s in %s", forbidden, body)
		}
	}
}
