package codexcost

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestStatusDoesNotExposeRawUsageFields(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/codex-cost/usage.jsonl",
		`{"at":"2026-06-19T00:00:00Z","scope":"assistant_loop","unit_kind":"codex_tokens","amount":1,"status":"recorded","evidence_refs":["docs/assistant-vision.md"],"raw_prompt":"private prompt","raw_transcript":"private transcript","private_notes":"private note"}`+"\n")

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range [][]byte{
		[]byte("private prompt"),
		[]byte("private transcript"),
		[]byte("private note"),
		[]byte("evidence_refs"),
		[]byte("raw_prompt"),
	} {
		if bytes.Contains(body, forbidden) {
			t.Fatalf("status leaked %s in %s", forbidden, body)
		}
	}
}
