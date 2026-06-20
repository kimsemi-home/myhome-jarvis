package daemon

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthorityReviewRequestReturnsRedactedPacket(t *testing.T) {
	server, err := New(DefaultConfig(repoRoot(t), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/authority-review/request", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{`"request_id":`, `"approval_granted": false`} {
		if !strings.Contains(body, expected) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{"raw_evidence", "reviewer_identity", "linear_url", "credential"} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("authority review request leaked %s in %s", forbidden, body)
		}
	}
}
