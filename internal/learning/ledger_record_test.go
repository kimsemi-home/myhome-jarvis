package learning

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRecordWritesPrivateObservationAndRedactedStatus(t *testing.T) {
	root := copyPolicyFixture(t)
	payload := []byte(`{"kind":"loop_gap","source":"quality_gate","summary":"Quality failed before a regression was captured.","evidence_refs":["data/private/quality/runs.jsonl","docs/working-log.md"],"owner":"go","next_action":"Add a focused regression test."}`)

	result, err := Record(root, payload)
	if err != nil {
		t.Fatal(err)
	}
	if result.Kind != "loop_gap" || result.Stage != "evidence_recorded" || result.Status != "open" || result.EvidenceRefCount != 2 {
		t.Fatalf("result = %#v", result)
	}
	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if !status.Exists || status.Count != 1 || status.OpenCount != 1 || status.ByKind["loop_gap"] != 1 || status.ByStage["evidence_recorded"] != 1 {
		t.Fatalf("status = %#v", status)
	}
	assertStatusRedacted(t, status)
	assertJournalContainsPrivateObservation(t, root, status.Path)
}

func assertStatusRedacted(t *testing.T, status Status) {
	t.Helper()
	statusPayload, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{"Quality failed", "Add a focused", "evidence_refs"} {
		if strings.Contains(string(statusPayload), forbidden) {
			t.Fatalf("status leaked %q in %s", forbidden, statusPayload)
		}
	}
}

func assertJournalContainsPrivateObservation(t *testing.T, root string, journalPath string) {
	t.Helper()
	journal, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(journalPath)))
	if err != nil {
		t.Fatal(err)
	}
	for _, expected := range []string{`"summary":"Quality failed before a regression was captured."`, `"next_action":"Add a focused regression test."`} {
		if !strings.Contains(string(journal), expected) {
			t.Fatalf("journal missing %s in %s", expected, journal)
		}
	}
	if strings.Contains(string(journal), root) {
		t.Fatalf("journal leaked root in %s", journal)
	}
}
