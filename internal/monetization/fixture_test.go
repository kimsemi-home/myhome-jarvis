package monetization

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func testPolicy() Policy {
	return Policy{
		Context:                  "MonetizationExperimentLedger",
		Version:                  "v1",
		PrivateExperimentLedger:  "data/private/monetization/experiments.jsonl",
		AppendOnly:               true,
		PublicStatusRedacted:     true,
		DecisionEvidenceRequired: true,
		CostEstimateRequired:     true,
		ExperimentStates:         []string{"backlog", "review_required", "running", "closed"},
		DecisionKinds:            []string{"hypothesis_created", "scale_requested", "close_experiment"},
		ReviewStatuses:           requiredReviews,
		ExpectedValueBands:       requiredBands,
		CostUnitKinds:            requiredCostUnits,
		RequiredFields:           requiredFields,
		PublicSummaryFields:      requiredSummaryFields,
		AllowedEvidencePrefixes:  []string{"generated/", "docs/", "data/private/"},
		Commands: []string{
			"mhj monetization status",
			"mhj monetization record <json-payload>",
		},
	}
}

func writePolicy(t *testing.T, root string, policy Policy) {
	t.Helper()
	body, err := json.Marshal(policy)
	if err != nil {
		t.Fatal(err)
	}
	writeFile(t, root, PolicyRelativePath, string(body)+"\n")
}

func writeFile(t *testing.T, root string, rel string, body string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
}
