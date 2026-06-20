package codexsustainability

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func testPolicy() Policy {
	return Policy{
		Context:                              "CodexSustainabilityEvidenceLoop",
		Version:                              "v1",
		GeneratedArtifact:                    PolicyRelativePath,
		PrivateEvidenceLedger:                "data/private/codex-sustainability/evidence.jsonl",
		AppendOnly:                           true,
		PublicStatusRedacted:                 true,
		TrendBaselinesVersioned:              true,
		EvidenceMaxAgeHours:                  168,
		TrendBaselineMaxAgeHours:             168,
		CostPerAcceptedChangeReviewThreshold: 500000,
		RecordKinds:                          requiredRecordKinds,
		Metrics:                              requiredMetrics,
		RequiredFields:                       requiredFields,
		ProposalRequiredFields:               requiredProposalFields,
		AllowedEvidencePrefixes:              []string{"generated/", "docs/", ".github/", "data/private/"},
		PublicSummaryFields:                  requiredSummaryFields,
		Commands: []string{
			"mhj codex-sustainability status",
			"mhj codex-sustainability record-quality",
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
