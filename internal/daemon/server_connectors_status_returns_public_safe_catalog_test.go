package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConnectorsStatusReturnsPublicSafeCatalog(t *testing.T) {
	server, err := New(DefaultConfig(repoRoot(t), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/connectors/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"fixture_only": true`,
		`"real_credentials_allowed": false`,
		`"external_api_calls_allowed": false`,
		`"generated_path": "generated/connectors.generated.json"`,
		`"key": "mydata"`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{
		`"token":`,
		`"secret":`,
		`"cookie":`,
		`"account_id":`,
		`"card_number":`,
		`"local_path":`,
		`"generated_path": "/"`,
		`"generated_path": "\\"`,
	} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("connector status leaked %s in %s", forbidden, body)
		}
	}
}
