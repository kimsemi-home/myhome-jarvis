package translation

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func testPolicy() Policy {
	return Policy{
		Context:                 "AgentCluster",
		Version:                 "v1",
		GeneratedArtifact:       "generated/translation.generated.json",
		PrivateLossLedger:       "data/private/translation/losses.jsonl",
		PrivateManifestRoot:     "data/private/translation/manifests",
		ManifestRequired:        true,
		PublicStatusRedacted:    true,
		RawLossPublicAllowed:    false,
		AllowedContexts:         translationTestContexts(),
		RequiredManifestFields:  translationTestManifestFields(),
		LossLevels:              translationTestLossLevels(),
		AllowedLossCategories:   translationTestLossCategories(),
		ForbiddenLossCategories: translationTestForbiddenLossCategories(),
		AllowedEvidencePrefixes: []string{"data/private/", "generated/", "docs/", "cmd/", "internal/", "apps/flutter/", "lisp/", "crates/", "fixtures/", "harness/", ".github/"},
		PublicSummaryFields:     []string{"policy_path", "ledger_path", "manifest_root", "ledger_exists", "manifest_root_exists", "manifest_count", "invalid_manifest_count", "missing_manifest_count", "loss_count", "open_loss_count", "closed_loss_count", "invalid_loss_count", "open_debt_count", "forbidden_loss_count", "by_level", "by_source_context", "by_target_context", "last_observed_at", "checked_at"},
		Commands:                []string{"mhj translation status"},
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
