package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestMetricsReturnsRuntimeCounters(t *testing.T) {
	server, err := New(DefaultConfig(t.TempDir(), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"requests": 1`,
		`"event_count": 0`,
		`"goroutine_count":`,
		`"heap_alloc_bytes":`,
		`"heap_sys_bytes":`,
		`"stack_inuse_bytes":`,
		`"gc_count":`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	localPathMarker := filepath.Join(string(os.PathSeparator), "Users")
	for _, forbidden := range []string{"root", "token", localPathMarker, "\\"} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("metrics leaked %q in %s", forbidden, body)
		}
	}
}
