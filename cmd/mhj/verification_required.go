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
		"generated/verification_graph.schema.generated.json",
		"generated/verification_conformance.generated.json",
		"generated/verification_tests.generated.json",
		"generated/release_pipeline.generated.json",
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
		"local-makefile-ssot-drift-check",
	}
}

func requiredVerificationCommands() []string {
	return []string{
		"test -s generated/control_plane_verification.generated.json",
		"go run ./cmd/mhj control-plane verify",
	}
}
