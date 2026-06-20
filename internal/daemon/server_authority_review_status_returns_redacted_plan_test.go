package daemon

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthorityReviewStatusReturnsRedactedPlan(t *testing.T) {
	server, err := New(DefaultConfig(repoRoot(t), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/authority-review/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"public_safe": true`,
		`"redaction": "review-classes-only"`,
		`"high_risk_blocked_decision_count": 6`,
		`"external_writes_allowed_profile_count": 0`,
	} {
		if !strings.Contains(body, expected) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{"raw_rationale", "raw_evidence", "reviewer_identity", "linear_url", "credential"} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("authority review status leaked %s in %s", forbidden, body)
		}
	}
}
