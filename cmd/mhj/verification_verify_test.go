package main

import (
	"strings"
	"testing"
)

func TestVerifyTestManifestRequiresCoverage(t *testing.T) {
	err := verifyTestManifest(verificationTestsFile{})
	if err == nil {
		t.Fatal("expected missing generated verification test coverage")
	}
	if !strings.Contains(err.Error(), "graph-artifacts-exist") {
		t.Fatalf("expected missing graph artifact test, got %v", err)
	}
}

func TestVerifyReleaseGatesRequireEveryUnit(t *testing.T) {
	graph := verificationGraphFile{Units: []verificationUnit{{ID: "go", Kind: "unit-test"}}}
	release := verificationReleaseFile{Gates: []verificationGate{{ID: "ssot", Kind: "conformance", Required: true}}}
	err := verifyReleaseGates(graph, release)
	if err == nil {
		t.Fatal("expected missing release gate")
	}
	if !strings.Contains(err.Error(), "go") {
		t.Fatalf("expected missing go gate, got %v", err)
	}
}
