package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDomainSummaryReturnsFixtureStatus(t *testing.T) {
	root := repoRoot(t)
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/domain/summary", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"net_minor_units": 4346800`,
		`"recurring_candidate_count": 1`,
		`"recommendations"`,
		`"household"`,
		`"long_term_format": "parquet"`,
		`"private_root": "data/lake"`,
		`"archive_root": "data/private/archive"`,
		`"config_evidence_field": "evidence_noise_budget"`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	if bytes.Contains([]byte(body), []byte(root)) {
		t.Fatalf("domain summary leaked local root in %s", body)
	}
	home, err := os.UserHomeDir()
	if err == nil && home != "" && bytes.Contains([]byte(body), []byte(home)) {
		t.Fatalf("domain summary leaked local home in %s", body)
	}
	for _, forbidden := range []string{`"private_root": "/"`, `"private_root": "\\"`, `"private_root": "../"`} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("domain summary leaked unsafe storage root %s in %s", forbidden, body)
		}
	}
}
