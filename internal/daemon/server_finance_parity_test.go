package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFinanceParityReturnsFixtureContract(t *testing.T) {
	server, err := New(DefaultConfig(repoRoot(t), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/finance/parity", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	for _, expected := range []string{`"provider": "toss_mydata"`, `"parity_percent": 100`, `"pass": true`} {
		if !bytes.Contains(recorder.Body.Bytes(), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, recorder.Body.String())
		}
	}
}
