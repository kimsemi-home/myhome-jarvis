package main

import (
	"strings"
	"testing"
)

func TestValidateCIWorkflowContractAcceptsRequiredTokens(t *testing.T) {
	root := t.TempDir()
	writeCIWorkflowFixture(t, root, ciWorkflowFixture())

	if err := validateCIWorkflowContract(root); err != nil {
		t.Fatalf("validateCIWorkflowContract() error = %v", err)
	}
}

func TestValidateCIWorkflowContractRejectsMissingCacheInput(t *testing.T) {
	root := t.TempDir()
	workflow := strings.ReplaceAll(ciWorkflowFixture(), "'rust-toolchain.toml', ", "")
	writeCIWorkflowFixture(t, root, workflow)

	err := validateCIWorkflowContract(root)
	if err == nil {
		t.Fatal("expected missing cache input to fail")
	}
	if !strings.Contains(err.Error(), "rust-toolchain.toml") {
		t.Fatalf("expected rust-toolchain.toml error, got %v", err)
	}
}

func TestValidateCIWorkflowContractRejectsPrivilegedTrigger(t *testing.T) {
	root := t.TempDir()
	workflow := strings.Replace(ciWorkflowFixture(), "  pull_request:\n", "  pull_request:\n  pull_request_target:\n", 1)
	writeCIWorkflowFixture(t, root, workflow)

	err := validateCIWorkflowContract(root)
	if err == nil {
		t.Fatal("expected privileged trigger to fail")
	}
	if !strings.Contains(err.Error(), "pull_request_target") {
		t.Fatalf("expected pull_request_target error, got %v", err)
	}
}

func TestValidateCIWorkflowContractRejectsAnyWritePermission(t *testing.T) {
	root := t.TempDir()
	workflow := strings.Replace(ciWorkflowFixture(), "  contents: read\n", "  contents: read\n  id-token: write\n", 1)
	writeCIWorkflowFixture(t, root, workflow)

	err := validateCIWorkflowContract(root)
	if err == nil {
		t.Fatal("expected write permission to fail")
	}
	if !strings.Contains(err.Error(), "id-token: write") {
		t.Fatalf("expected id-token write permission error, got %v", err)
	}
}
