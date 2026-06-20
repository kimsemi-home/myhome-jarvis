package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFinanceConsentStatusReturnsCountsOnly(t *testing.T) {
	root := t.TempDir()
	copyTestFile(t, repoRoot(t), root, "generated/finance_consent.generated.json")
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/finance-consent/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"readiness_state": "blocked"`,
		`"finance_mode": "read_only_review_only"`,
		`"record_count": 0`,
		`"missing_required_consent_count": 3`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{"private_notes", "evidence_refs", "account_id", "card_number"} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("finance consent status leaked %s in %s", forbidden, body)
		}
	}
}
