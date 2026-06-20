package codexcost

import (
	"bytes"
	"testing"
)

func TestROISummaryDoesNotExposeRawUsageFields(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeSustainabilityPolicy(t, root)
	writeSustainableLedger(t, root)
	writeStoragePolicy(t, root)
	writeFile(t, root, "data/private/codex-cost/usage.jsonl",
		`{"at":"2026-06-19T00:00:00Z","scope":"repo","unit_kind":"codex_tokens","amount":1,"status":"recorded","evidence_refs":["docs/codex-cost-governor.md"],"raw_prompt":"private prompt","raw_transcript":"private transcript","private_notes":"private note"}`+"\n")
	writeFile(t, root, "data/private/codex-cost/attribution.jsonl",
		`{"at":"2026-06-19T00:00:01Z","scope":"repo","subject_key":"repo:private-subject","unit_kind":"codex_tokens","amount":1,"basis":"manual","evidence_refs":["docs/codex-cost-governor.md"]}`+"\n")

	summary, err := ROISummaryForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	body := mustJSON(t, summary)
	for _, forbidden := range [][]byte{
		[]byte("private prompt"),
		[]byte("private transcript"),
		[]byte("private note"),
		[]byte("private-subject"),
		[]byte("subject_key"),
		[]byte("evidence_refs"),
		[]byte("raw_prompt"),
	} {
		if bytes.Contains(body, forbidden) {
			t.Fatalf("roi summary leaked %s in %s", forbidden, body)
		}
	}
}
