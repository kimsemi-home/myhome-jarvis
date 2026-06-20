package main

func requiredVerificationArtifacts() []string {
	return []string{
		"generated/verification_graph.generated.json",
		"generated/github_quality_workflow.generated.yml",
		".github/workflows/quality.yml",
		"generated/gitlab_quality.generated.yml",
		"generated/local_quality.generated.mk",
		"generated/bazel_quality.generated.bzl",
		"generated/control_plane_verification.generated.json",
		"generated/verification_evidence.generated.json",
		"generated/pdca.generated.json",
		"generated/verification_graph.schema.generated.json",
		"generated/verification_conformance.generated.json",
		"generated/verification_tests.generated.json",
		"generated/release_pipeline.generated.json",
		"generated/assistant_vision.generated.json",
		"generated/codex_cost.generated.json",
		"generated/codex_sustainability.generated.json",
		"generated/finance_consent.generated.json",
		"generated/monetization.generated.json",
		"generated/repo_factory.generated.json",
		"docs/verification-graph.md",
	}
}

func requiredVerificationTests() []string {
	return []string{
		"graph-artifacts-exist",
		"backend-artifacts-exist",
		"schema-json-valid",
		"conformance-manifest-linked",
		"release-gates-cover-units",
		"control-plane-verifier-linked",
		"verification-evidence-linked",
		"pdca-linked",
		"local-makefile-ssot-drift-check",
	}
}

func requiredVerificationCommands() []string {
	return []string{
		"go run ./cmd/mhj control-plane verify",
		"go run ./cmd/mhj verification evidence",
		"go run ./cmd/mhj pdca status",
		"test -s generated/control_plane_verification.generated.json",
		"test -s generated/verification_evidence.generated.json",
		"test -s generated/pdca.generated.json",
		"test -s generated/codex_cost.generated.json",
		"test -s generated/codex_sustainability.generated.json",
		"test -s generated/finance_consent.generated.json",
		"test -s generated/monetization.generated.json",
		"test -s generated/repo_factory.generated.json",
	}
}

func requiredEvidenceSources() []string {
	return []string{
		"github-job-logs",
		"unit-cache-keys",
		"generated-backend-specs",
		"verification-manifests",
		"local-quality-run-ledger",
	}
}
