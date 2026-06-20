package mergeevidence

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func testPolicy() Policy {
	return Policy{
		Context:                   "MergeEvidencePolicy",
		Version:                   "v1",
		GeneratedArtifact:         PolicyRelativePath,
		DefaultBehavior:           "merge_when_eligible",
		PublicStatusRedacted:      true,
		MergeWithoutReviewAllowed: false,
		PersistPrivateEvidence:    false,
		Gates: []Gate{
			{Key: "clean_branch", Label: "Clean branch", Evidence: "git status", Required: true, BlocksMerge: true},
			{Key: "required_checks_success", Label: "Checks", Evidence: "GitHub Actions", Required: true, BlocksMerge: true},
			{Key: "public_safety_passed", Label: "Public safety", Evidence: "mhj security", Required: true, BlocksMerge: true},
			{Key: "review_gate_clear", Label: "Review", Evidence: "review state", Required: true, BlocksMerge: true},
			{Key: "generated_drift_clear", Label: "Generated drift", Evidence: "codegen diff", Required: true, BlocksMerge: true},
		},
		RequiredEvidence:    requiredEvidence,
		PublicSummaryFields: requiredSummaryFields,
		Commands:            []string{"mhj merge-evidence status"},
	}
}

func writePolicy(t *testing.T, root string, policy Policy) {
	t.Helper()
	body, err := json.Marshal(policy)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(root, filepath.FromSlash(PolicyRelativePath))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, append(body, '\n'), 0o600); err != nil {
		t.Fatal(err)
	}
}
