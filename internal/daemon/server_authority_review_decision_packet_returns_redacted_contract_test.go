package daemon

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthorityReviewDecisionPacketReturnsRedactedContract(t *testing.T) {
	server, err := New(DefaultConfig(repoRoot(t), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/authority-review/decision-packet", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range authorityReviewDecisionPacketExpectedFields() {
		if !strings.Contains(body, expected) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range authorityReviewDecisionPacketForbiddenFields() {
		if strings.Contains(body, forbidden) {
			t.Fatalf("authority review decision packet leaked %s in %s", forbidden, body)
		}
	}
}

func authorityReviewDecisionPacketExpectedFields() []string {
	return []string{
		`"context": "AuthorityReviewDecisionPacket"`,
		`"decision_contract": {`,
		`"reviewer_posture": "human_only_non_delegable"`,
		`"review_only": true`,
		`"can_apply_decision": false`,
		`"approval_granted": false`,
		`"external_writes_allowed": false`,
		`"repo_creation_allowed": false`,
		`"workflow_changes_allowed": false`,
		`"self_approval_allowed": false`,
	}
}

func authorityReviewDecisionPacketForbiddenFields() []string {
	return []string{
		"raw_evidence", "reviewer_identity", "linear_url",
		"workspace_url", "credential", "private_notes",
	}
}
