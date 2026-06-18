package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReviewStatusReturnsRedactedCapacitySummary(t *testing.T) {
	server, err := New(DefaultConfig(repoRoot(t), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/review/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"policy_path": "generated/review.generated.json"`,
		`"queue_path": "data/private/review/queue.jsonl"`,
		`"capacity_state":`,
		`"review_debt_count":`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{
		`"raw_rationale":`,
		`"raw_review":`,
		`"raw_review_notes":`,
		`"reviewer_identity":`,
		`"reviewer_name":`,
		`"reviewer_email":`,
		`"evidence_ref":`,
		`"evidence_refs":`,
		`"raw_prompt":`,
		`"raw_transcript":`,
		`"token":`,
		`"secret":`,
		`"credential":`,
		`"queue_path": "/"`,
		`"queue_path": "\\"`,
	} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("review status leaked %s in %s", forbidden, body)
		}
	}
}
