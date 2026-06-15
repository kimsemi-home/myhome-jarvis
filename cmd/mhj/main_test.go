package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestChangedGeneratedFilesReportsAddedModifiedAndDeleted(t *testing.T) {
	before := map[string][]byte{
		"generated/a.json": []byte("same"),
		"generated/b.json": []byte("old"),
		"generated/c.json": []byte("removed"),
	}
	after := map[string][]byte{
		"generated/a.json": []byte("same"),
		"generated/b.json": []byte("new"),
		"generated/d.json": []byte("added"),
	}

	changed := changedGeneratedFiles(before, after)
	expected := []string{
		"generated/b.json",
		"generated/c.json",
		"generated/d.json",
	}
	if !reflect.DeepEqual(changed, expected) {
		t.Fatalf("changed files = %#v, expected %#v", changed, expected)
	}
}

func TestQualityReportJSONRedactsCommandAndOutput(t *testing.T) {
	report := qualityReport{
		OK: true,
		Steps: []qualityStep{{
			Name:    "flutter test",
			Status:  "pass",
			Command: []string{"/private/toolchains/flutter", "test"},
			Output:  "loading /private/checkout/apps/flutter/test/widget_test.dart",
		}},
	}

	payload, err := json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}
	body := string(payload)
	for _, expected := range []string{`"ok":true`, `"name":"flutter test"`, `"status":"pass"`} {
		if !strings.Contains(body, expected) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{`command`, `output`, `/private/toolchains`, `/private/checkout`} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("quality report leaked %s in %s", forbidden, body)
		}
	}
}

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

func TestValidateCIWorkflowContractAcceptsRequiredTokens(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, ".github/workflows/quality.yml", ciWorkflowFixture())

	if err := validateCIWorkflowContract(root); err != nil {
		t.Fatalf("validateCIWorkflowContract() error = %v", err)
	}
}

func TestValidateCIWorkflowContractRejectsMissingCacheInput(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, ".github/workflows/quality.yml", strings.ReplaceAll(ciWorkflowFixture(), "'rust-toolchain.toml', ", ""))

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
	writeTestFile(t, root, ".github/workflows/quality.yml", strings.Replace(ciWorkflowFixture(), "  pull_request:\n", "  pull_request:\n  pull_request_target:\n", 1))

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
	writeTestFile(t, root, ".github/workflows/quality.yml", workflow)

	err := validateCIWorkflowContract(root)
	if err == nil {
		t.Fatal("expected write permission to fail")
	}
	if !strings.Contains(err.Error(), "id-token: write") {
		t.Fatalf("expected id-token write permission error, got %v", err)
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

func ciWorkflowFixture() string {
	return `name: quality
on:
  push:
  pull_request:
concurrency:
  cancel-in-progress: true
permissions:
  contents: read
jobs:
  public-safety:
    steps:
      - uses: actions/checkout@v6
        with:
          fetch-depth: 0
      - run: |
          go run ./cmd/mhj security check
          go run ./cmd/mhj security history
  go:
    steps:
      - uses: actions/cache/restore@v5
        with:
          key: go-${{ hashFiles('.github/workflows/quality.yml', '.go-version', 'rust-toolchain.toml', 'generated/*.json') }}
      - run: go run ./cmd/mhj ci verify
      - run: go run ./cmd/mhj toolchain verify
      - uses: actions/cache/save@v5
        if: steps.unit-cache.outputs.cache-hit != 'true' && github.event_name == 'push' && github.repository == 'kimsemi-home/myhome-jarvis'
  flutter:
    steps:
      - uses: actions/cache/restore@v5
        with:
          key: flutter-${{ hashFiles('generated/commands.generated.json', 'generated/connectors.generated.json') }}
`
}

func writeTestFile(t *testing.T, root string, rel string, body string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
}
