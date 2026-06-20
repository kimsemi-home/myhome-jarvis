package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestContextPackStatusReturnsPublicSafeHandoff(t *testing.T) {
	root := t.TempDir()
	copyTestFile(t, repoRoot(t), root, "generated/context_pack.generated.json")
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/context-pack/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"policy_path": "generated/context_pack.generated.json"`,
		`"pack_id": "myhome-jarvis/context-pack"`,
		`"public_safe": true`,
		`"verification_profile": "quality"`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{
		`"raw_prompt":`, `"credential":`, `"private_evidence":`, `"local_absolute_path":`,
	} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("context pack status leaked %s in %s", forbidden, body)
		}
	}
}
