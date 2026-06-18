package main

import (
	"strings"
	"testing"
)

func TestValidateToolchainPinsAcceptsMatchingPins(t *testing.T) {
	root := writeToolchainFixture(t, "1.26.2", "1.26.2", "1.26.2", "1.96.0", "1.96.0")

	if err := validateToolchainPins(root); err != nil {
		t.Fatalf("validateToolchainPins() error = %v", err)
	}
}

func TestValidateToolchainPinsRejectsDrift(t *testing.T) {
	root := writeToolchainFixture(t, "1.26.2", "1.26.1", "1.26.2", "1.96.0", "1.96.0")

	err := validateToolchainPins(root)
	if err == nil {
		t.Fatal("expected drift to fail")
	}
	if !strings.Contains(err.Error(), "go.mod go directive") {
		t.Fatalf("expected go.mod drift error, got %v", err)
	}
}

func writeToolchainFixture(t *testing.T, goVersion string, goModVersion string, workflowGoVersion string, rustVersion string, workflowRustVersion string) string {
	t.Helper()
	root := t.TempDir()
	writeTestFile(t, root, ".go-version", goVersion+"\n")
	writeTestFile(t, root, "go.mod", "module github.com/kimsemi-home/myhome-jarvis\n\ngo "+goModVersion+"\n")
	writeTestFile(t, root, "rust-toolchain.toml", "[toolchain]\nchannel = \""+rustVersion+"\"\nprofile = \"minimal\"\ncomponents = [\"rustfmt\", \"clippy\"]\n")
	writeTestFile(t, root, ".github/workflows/quality.yml", "env:\n  GO_VERSION: \""+workflowGoVersion+"\"\n  FLUTTER_VERSION: \"3.44.2\"\n  RUST_TOOLCHAIN: \""+workflowRustVersion+"\"\n")
	writeTestFile(t, root, "generated/commands.generated.json", `{"project":{"go_version":"`+goVersion+`"},"commands":[]}`+"\n")
	return root
}
