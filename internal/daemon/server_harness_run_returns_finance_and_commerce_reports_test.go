package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHarnessRunReturnsFinanceAndCommerceReports(t *testing.T) {
	server, err := New(DefaultConfig(repoRoot(t), "test"))
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name     string
		expected string
	}{
		{name: "finance", expected: `"name": "finance"`},
		{name: "commerce", expected: `"name": "commerce"`},
	} {
		request := httptest.NewRequest(http.MethodPost, "/harness/run", bytes.NewBufferString(`{"name":"`+tc.name+`"}`))
		request.RemoteAddr = "127.0.0.1:1234"
		recorder := httptest.NewRecorder()

		server.Routes().ServeHTTP(recorder, request)

		if recorder.Code != http.StatusOK {
			t.Fatalf("%s status = %d body = %s", tc.name, recorder.Code, recorder.Body.String())
		}
		body := recorder.Body.String()
		for _, expected := range []string{
			tc.expected,
			`"passed": true`,
		} {
			if !bytes.Contains([]byte(body), []byte(expected)) {
				t.Fatalf("expected %s in %s", expected, body)
			}
		}
	}
}
