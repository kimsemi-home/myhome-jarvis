package codexcost

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func testPolicy() Policy {
	return Policy{
		Context:                 "CodexCostGovernor",
		Version:                 "v1",
		GeneratedArtifact:       "generated/codex_cost.generated.json",
		PrivateUsageLedger:      "data/private/codex-cost/usage.jsonl",
		AppendOnly:              true,
		PublicStatusRedacted:    true,
		RawUsagePublicAllowed:   false,
		UnitKinds:               requiredUnitKinds,
		LoopScopes:              requiredLoopScopes,
		RecordStatuses:          []string{"recorded", "review_required", "approved", "rejected"},
		WarningUnitThreshold:    100000,
		ReviewUnitThreshold:     500000,
		RequiredFields:          requiredRecordFields,
		AllowedEvidencePrefixes: []string{"generated/", "docs/", ".github/", "data/private/"},
		PublicSummaryFields:     requiredSummaryFields,
		Commands:                []string{"mhj codex-cost status"},
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
