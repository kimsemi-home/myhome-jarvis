package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthorityStatusReturnsRedactedGateSummary(t *testing.T) {
	server, err := New(DefaultConfig(repoRoot(t), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/authority/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"policy_path": "generated/authority.generated.json"`,
		`"reasoning_tier_grants_approval": false`,
		`"self_authority_allowed": false`,
		`"public_repo_mode": true`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{
		`"raw_rationale":`,
		`"raw_evidence":`,
		`"evidence_ref":`,
		`"evidence_refs":`,
		`"raw_prompt":`,
		`"raw_transcript":`,
		`"token":`,
		`"secret":`,
		`"credential":`,
		`"policy_path": "/"`,
		`"policy_path": "\\"`,
	} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("authority status leaked %s in %s", forbidden, body)
		}
	}
}
