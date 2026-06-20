package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRepoFactoryStatusReturnsPublicSafeSummary(t *testing.T) {
	root := t.TempDir()
	copyTestFile(t, repoRoot(t), root, "generated/repo_factory.generated.json")
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/repo-factory/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"policy_path": "generated/repo_factory.generated.json"`,
		`"generated_ci_present": true`,
		`"security_scan_present": true`,
		`"private_data_policy_present": true`,
		`"repo_creation_blocked_until_review": true`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{"absolute_home_path", "old_private_owner", "client_secret", "access_token"} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("repo factory status leaked %s in %s", forbidden, body)
		}
	}
}
