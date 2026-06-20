package daemon

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMergeEvidenceStatusReturnsPublicPolicy(t *testing.T) {
	root := t.TempDir()
	copyTestFile(t, repoRoot(t), root, "generated/merge_evidence.generated.json")
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/merge-evidence/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"context": "MergeEvidencePolicy"`,
		`"merge_preference": "merge_after_checks_pass"`,
		`"private_data_scan_required": true`,
		`"merge_ready": true`,
	} {
		if !strings.Contains(body, expected) {
			t.Fatalf("missing %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{"private_linear_url", "raw_review_notes", "credential"} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("merge evidence status leaked %s in %s", forbidden, body)
		}
	}
}
