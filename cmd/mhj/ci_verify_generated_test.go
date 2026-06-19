package main

import (
	"strings"
	"testing"
)

func writeCIWorkflowFixture(t *testing.T, root string, workflow string) {
	t.Helper()
	writeTestFile(t, root, ".github/workflows/quality.yml", workflow)
	writeTestFile(t, root, "generated/github_quality_workflow.generated.yml", workflow)
}

func TestValidateCIWorkflowContractRejectsGeneratedDrift(t *testing.T) {
	root := t.TempDir()
	writeCIWorkflowFixture(t, root, ciWorkflowFixture())
	writeTestFile(t, root, ".github/workflows/quality.yml",
		strings.Replace(ciWorkflowFixture(), "name: quality", "name: changed", 1))

	err := validateCIWorkflowContract(root)
	if err == nil {
		t.Fatal("expected generated workflow drift to fail")
	}
	if !strings.Contains(err.Error(), "generated/github_quality_workflow.generated.yml") {
		t.Fatalf("expected generated workflow error, got %v", err)
	}
}
