package daemon

import (
	"bytes"
	"github.com/kimsemi-home/myhome-jarvis/internal/qualitylog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestQualityStatusReturnsLastPrivateRun(t *testing.T) {
	root := t.TempDir()
	if err := qualitylog.AppendRun(root, qualitylog.NewRun(time.Now(), true, []qualitylog.Step{
		{Name: "go test", Status: "pass"},
		{Name: "flutter analyze", Status: "pass"},
	})); err != nil {
		t.Fatal(err)
	}
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/quality/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"exists": true`,
		`"count": 1`,
		`"ok": true`,
		`"pass_count": 2`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
}
