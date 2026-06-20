package repofactory

import "testing"

func TestStatusSummarizesPublicSafeRepoFactory(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if !status.PublicSafe {
		t.Fatalf("expected public-safe status: %#v", status)
	}
	if !status.GeneratedCIPresent || !status.SecurityScanPresent ||
		!status.PrivateDataPolicyPresent || !status.BootstrapChecklistPresent ||
		!status.CodexProjectTemplatePresent || !status.ContextPackDeclarationPresent {
		t.Fatalf("template coverage = %#v", status)
	}
	if !status.RepoCreationBlockedUntilReview || status.CreationAllowedWithoutReview {
		t.Fatalf("review gate = %#v", status)
	}
}

func TestStatusCountsForbiddenTemplateValues(t *testing.T) {
	root := t.TempDir()
	policy := testPolicy()
	policy.TemplateFiles[0].Path = "/" + "Users" + "/local/repo"
	writePolicy(t, root, policy)

	status := statusFromPolicy(policy)
	if status.PublicSafe || status.ForbiddenTemplateValueCount != 1 {
		t.Fatalf("status = %#v", status)
	}
}
