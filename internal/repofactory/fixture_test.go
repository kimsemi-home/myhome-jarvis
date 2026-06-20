package repofactory

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func testPolicy() Policy {
	return Policy{
		Context:                          "PublicSafeRepoFactory",
		Version:                          "v1",
		GeneratedArtifact:                PolicyRelativePath,
		TargetOwner:                      "kimsemi-home",
		PublicRepoDefault:                true,
		CodexProjectRequired:             true,
		RepoCreationAllowedWithoutReview: false,
		PublicSafetyEvidenceRequired:     true,
		AuthorityReviewRequired:          true,
		PrivateAssetsPublicAllowed:       false,
		LocalPathsPublicAllowed:          false,
		TemplateFiles: []TemplateFile{
			{Role: "generated_ci", Path: ".github/workflows/quality.yml", SourceArtifact: "generated/github_quality_workflow.generated.yml", Purpose: "run generated quality checks"},
			{Role: "security_scan", Path: "docs/security.md", SourceArtifact: "generated/security.generated.json", Purpose: "document public safety checks"},
			{Role: "private_data_policy", Path: "docs/private-data-policy.md", SourceArtifact: PolicyRelativePath, Purpose: "keep private assets outside public files"},
			{Role: "bootstrap_checklist", Path: "docs/bootstrap-checklist.md", SourceArtifact: PolicyRelativePath, Purpose: "prove review gate completion"},
			{Role: "codex_project", Path: ".codex/project-goal.md", SourceArtifact: "generated/assistant_vision.generated.json", Purpose: "seed the public-safe Codex project goal"},
			{Role: "context_pack_declaration", Path: ".mhj/context-pack.json", SourceArtifact: "generated/context_pack.generated.json", Purpose: "declare consumed context pack"},
		},
		CreationGates: []CreationGate{
			{Key: "authority_review", Required: true, BlocksRepoCreation: true, Evidence: "approved review record"},
			{Key: "public_safety_evidence", Required: true, BlocksRepoCreation: true, Evidence: "mhj security check and mhj security history"},
			{Key: "generated_ci", Required: true, BlocksRepoCreation: true, Evidence: "generated GitHub Actions workflow"},
			{Key: "private_data_policy", Required: true, BlocksRepoCreation: true, Evidence: "public private-data policy document"},
			{Key: "bootstrap_checklist", Required: true, BlocksRepoCreation: true, Evidence: "completed bootstrap checklist"},
		},
		BootstrapChecklist: requiredChecklistItems,
		AllowedPublicPathPrefixes: []string{
			".github/", ".codex/", ".mhj/", "cmd/", "docs/", "generated/", "internal/", "README.md", "LICENSE",
		},
		ForbiddenPublicFragments: []string{
			"absolute_home_path", "old_private_owner", "private_team_slug", "private_storage_prefix",
		},
		PublicSummaryFields: requiredSummaryFields,
		Commands:            []string{"mhj repo-factory status"},
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
