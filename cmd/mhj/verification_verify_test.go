package main

import (
	"strings"
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/qualitylog"
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

func TestVerifyGraphCommandsRequireControlPlaneVerifier(t *testing.T) {
	graph := verificationGraphFile{Units: []verificationUnit{{ID: "go", Commands: []string{
		"go run ./cmd/mhj ci-cache status",
		"go run ./cmd/mhj verification verify",
		"test -s generated/control_plane_verification.generated.json",
	}}}}
	err := verifyGraphCommands(graph)
	if err == nil || !strings.Contains(err.Error(), "control-plane verify") {
		t.Fatalf("expected missing control-plane verifier command, got %v", err)
	}
}

func TestVerifyEvidenceSourcesRequireQualityLedger(t *testing.T) {
	manifest := verificationEvidenceFile{Sources: []verificationEvidenceSource{
		{ID: "github-job-logs", Kind: "remote", Evidence: "logs"},
		{ID: "unit-cache-keys", Kind: "cache", Evidence: "cache"},
		{ID: "generated-backend-specs", Kind: "artifact", Evidence: "backend"},
		{ID: "verification-manifests", Kind: "artifact", Evidence: "manifest"},
	}}
	graph := verificationGraphFile{Evidence: []string{"redacted local quality run ledger"}}
	err := verifyEvidenceSources(manifest, graph, emptyQualityStatus())
	if err == nil || !strings.Contains(err.Error(), "local-quality-run-ledger") {
		t.Fatalf("expected missing quality ledger evidence, got %v", err)
	}
}

func emptyQualityStatus() qualitylog.Status {
	return qualitylog.Status{Path: qualitylog.RelativePath}
}
