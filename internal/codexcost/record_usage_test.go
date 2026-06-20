package codexcost

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestRecordUsageAppendsPrivateRecordAndReturnsRedactedResult(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	result, err := RecordUsage(root, []byte(`{
		"scope":"assistant_loop",
		"unit_kind":"codex_tokens",
		"amount":12345,
		"evidence_refs":["docs/assistant-vision.md"]
	}`))
	if err != nil {
		t.Fatal(err)
	}
	if result.Scope != "assistant_loop" || result.Status != "recorded" {
		t.Fatalf("result = %#v", result)
	}
	if result.EvidenceRefCount != 1 || result.BudgetState != "ok" {
		t.Fatalf("redacted result = %#v", result)
	}
	body := readLedger(t, root)
	if !bytes.Contains(body, []byte(`"semantic_hash":"cost_`)) {
		t.Fatalf("ledger missing semantic hash: %s", body)
	}
	if bytes.Contains(mustJSON(t, result), []byte("evidence_refs")) {
		t.Fatalf("result leaked evidence refs: %#v", result)
	}
}

func TestRecordUsageRejectsInvalidPayloads(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	for _, payload := range []string{
		`{"scope":"assistant_loop","unit_kind":"bad","amount":1,"evidence_refs":["docs/assistant-vision.md"]}`,
		`{"scope":"assistant_loop","unit_kind":"codex_tokens","amount":1,"evidence_refs":["/tmp/evidence.jsonl"]}`,
		`{"scope":"assistant_loop","unit_kind":"codex_tokens","amount":1,"raw_prompt":"private","evidence_refs":["docs/assistant-vision.md"]}`,
	} {
		if _, err := RecordUsage(root, []byte(payload)); err == nil {
			t.Fatalf("expected error for %s", payload)
		}
	}
}

func readLedger(t *testing.T, root string) []byte {
	t.Helper()
	body, err := os.ReadFile(filepath.Join(root, "data/private/codex-cost/usage.jsonl"))
	if err != nil {
		t.Fatal(err)
	}
	return body
}

func mustJSON(t *testing.T, value any) []byte {
	t.Helper()
	body, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	return body
}
