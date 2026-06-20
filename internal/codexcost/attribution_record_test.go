package codexcost

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestAttributeCostWritesPrivateRedactedAttribution(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())

	result, err := AttributeCost(root, []byte(`{
		"scope":"repo",
		"subject_key":"repo:kimsemi-home/myhome-jarvis",
		"unit_kind":"codex_tokens",
		"amount":100,
		"basis":"merged_pr",
		"evidence_refs":["docs/codex-cost-governor.md"]
	}`))
	if err != nil {
		t.Fatal(err)
	}
	if result.Scope != "repo" || result.SubjectHash == "" {
		t.Fatalf("result = %#v", result)
	}
	body := mustJSON(t, result)
	if bytes.Contains(body, []byte("subject_key")) ||
		bytes.Contains(body, []byte("evidence_refs")) {
		t.Fatalf("result leaked private attribution fields: %s", body)
	}
	ledger := readAttributionLedger(t, root)
	if !bytes.Contains(ledger, []byte(`"semantic_hash":"cost_attr_`)) {
		t.Fatalf("ledger missing attribution hash: %s", ledger)
	}
}

func readAttributionLedger(t *testing.T, root string) []byte {
	t.Helper()
	body, err := os.ReadFile(filepath.Join(root, "data/private/codex-cost/attribution.jsonl"))
	if err != nil {
		t.Fatal(err)
	}
	return body
}
