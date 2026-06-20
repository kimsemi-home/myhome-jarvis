package monetization

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestRecordExperimentAppendsPrivateRecordAndReturnsRedactedResult(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	result, err := RecordExperiment(root, []byte(`{
		"experiment_id":"shorts-factory-loop-001",
		"hypothesis_key":"shorts_factory_repo_bootstrap",
		"state":"backlog",
		"decision_kind":"hypothesis_created",
		"review_status":"not_required",
		"expected_value_band":"medium",
		"cost_estimate_units":1200,
		"cost_unit_kind":"codex_tokens",
		"evidence_refs":["generated/assistant_vision.generated.json","docs/monetization-experiments.md"]
	}`))
	if err != nil {
		t.Fatal(err)
	}
	if result.ExperimentID != "shorts-factory-loop-001" ||
		result.EvidenceRefCount != 2 {
		t.Fatalf("result = %#v", result)
	}
	if result.MonetizationDebtCount != 0 || result.RecordedAt == "" {
		t.Fatalf("redacted result = %#v", result)
	}
	body := readMonetizationLedger(t, root)
	if !bytes.Contains(body, []byte(`"evidence_refs"`)) {
		t.Fatalf("ledger missing private evidence refs: %s", body)
	}
	if bytes.Contains(mustMonetizationJSON(t, result), []byte("evidence_refs")) {
		t.Fatalf("result leaked evidence refs: %#v", result)
	}
}

func TestRecordExperimentRejectsInvalidPayloads(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	for _, payload := range []string{
		`{"experiment_id":"x","hypothesis_key":"y","state":"bad","decision_kind":"hypothesis_created","review_status":"not_required","expected_value_band":"medium","cost_estimate_units":1,"cost_unit_kind":"codex_tokens","evidence_refs":["docs/monetization-experiments.md"]}`,
		`{"experiment_id":"x","hypothesis_key":"y","state":"backlog","decision_kind":"hypothesis_created","review_status":"not_required","expected_value_band":"medium","cost_estimate_units":0,"cost_unit_kind":"codex_tokens","evidence_refs":["docs/monetization-experiments.md"]}`,
		`{"experiment_id":"x","hypothesis_key":"y","state":"backlog","decision_kind":"hypothesis_created","review_status":"not_required","expected_value_band":"medium","cost_estimate_units":1,"cost_unit_kind":"codex_tokens","evidence_refs":["/tmp/evidence.jsonl"]}`,
		`{"experiment_id":"x","hypothesis_key":"y","state":"backlog","decision_kind":"hypothesis_created","review_status":"not_required","expected_value_band":"medium","cost_estimate_units":1,"cost_unit_kind":"codex_tokens","private_revenue_notes":"private","evidence_refs":["docs/monetization-experiments.md"]}`,
	} {
		if _, err := RecordExperiment(root, []byte(payload)); err == nil {
			t.Fatalf("expected error for %s", payload)
		}
	}
}

func readMonetizationLedger(t *testing.T, root string) []byte {
	t.Helper()
	body, err := os.ReadFile(filepath.Join(root, "data/private/monetization/experiments.jsonl"))
	if err != nil {
		t.Fatal(err)
	}
	return body
}

func mustMonetizationJSON(t *testing.T, value any) []byte {
	t.Helper()
	body, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	return body
}
