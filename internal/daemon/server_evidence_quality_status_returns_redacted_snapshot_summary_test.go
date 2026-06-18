package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEvidenceQualityStatusReturnsRedactedSnapshotSummary(t *testing.T) {
	root := t.TempDir()
	copyTestFile(t, repoRoot(t), root, "generated/evidence_quality.generated.json")
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/evidence-quality/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"policy_path": "generated/evidence_quality.generated.json"`,
		`"ledger_path": "data/private/evidence-quality/snapshots.jsonl"`,
		`"exists": false`,
		`"snapshot_count": 0`,
		`"reassessment_debt_count": 0`,
		`"stale_after_hours": 720`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{
		`"raw_notes":`,
		`"raw_evidence":`,
		`"evidence_ref":`,
		`"evidence_refs":`,
		`"summary":`,
		`"raw_prompt":`,
		`"raw_transcript":`,
		`"token":`,
		`"secret":`,
		`"credential":`,
		`"ledger_path": "/"`,
	} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("evidence quality status leaked %s in %s", forbidden, body)
		}
	}
}
