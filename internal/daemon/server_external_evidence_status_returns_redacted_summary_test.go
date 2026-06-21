package daemon

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestExternalEvidenceStatusReturnsRedactedSummary(t *testing.T) {
	root := t.TempDir()
	copyTestFile(t, repoRoot(t), root, "generated/external_evidence.generated.json")
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/external-evidence/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"policy_path": "generated/external_evidence.generated.json"`,
		`"repo_creation_gate": "authority_review_required"`,
		`"raw_payload_public_allowed": false`,
	} {
		if !strings.Contains(body, expected) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{"raw_body", "cookie_value", "credential_value", "file://"} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("external evidence status leaked %s in %s", forbidden, body)
		}
	}
}
