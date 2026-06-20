package codexsustainability

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestStatusDoesNotExposePrivateEvidencePayload(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/codex-sustainability/evidence.jsonl",
		debtLedgerFixture()+`{"at":"2026-06-19T06:00:00Z","record_kind":"usage_sample","metric":"codex_coin","amount":1,"evidence_refs":["docs/assistant-vision.md"],"raw_transcript":"private transcript","private_notes":"private note","local_absolute_path":"/private/path","private_finance_data":"secret","unpublished_revenue_detail":"secret"}`+"\n")

	status, err := statusForRootAt(root, mustTime(t, "2026-06-20T00:00:00Z"))
	if err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range [][]byte{
		[]byte("private prompt"), []byte("private transcript"),
		[]byte("private note"), []byte("/private/path"),
		[]byte("private_finance_data"), []byte("unpublished_revenue_detail"),
		[]byte("evidence_refs"), []byte("raw_prompt"),
	} {
		if bytes.Contains(body, forbidden) {
			t.Fatalf("status leaked %s in %s", forbidden, body)
		}
	}
}
