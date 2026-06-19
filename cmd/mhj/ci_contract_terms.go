package main

func requiredCIWorkflowTokens() []string {
	return append(baseCIWorkflowTokens(), generatedCIWorkflowTokens()...)
}

func baseCIWorkflowTokens() []string {
	return []string{
		"cancel-in-progress: true", "permissions:", "contents: read",
		"fetch-depth: 0", "go run ./cmd/mhj security check",
		"go run ./cmd/mhj security history", "go run ./cmd/mhj ci verify",
		"go run ./cmd/mhj verification verify",
		"go run ./cmd/mhj code-shape status", "go run ./cmd/mhj toolchain verify",
		"'.go-version'", "'rust-toolchain.toml'", "'generated/*.json'",
		"'generated/commands.generated.json'", "'generated/connectors.generated.json'",
		"'generated/agent_cluster.generated.json'", "'generated/learning.generated.json'",
		"'generated/evidence.generated.json'", "'generated/confidence.generated.json'",
		"'generated/translation.generated.json'", "'generated/control_plane.generated.json'",
		"'generated/incidents.generated.json'", "'generated/evidence_quality.generated.json'",
		"'generated/review.generated.json'", "'generated/code_shape.generated.json'",
		"'generated/authority.generated.json'", "LISP: \"sbcl-bin\"",
		"40ants/setup-lisp@v4", "ros -Q run -- --script lisp/scripts/validate-ssot.lisp",
		"ros -Q run -- --script lisp/scripts/codegen.lisp",
		"github.event_name == 'push' && github.repository == 'kimsemi-home/myhome-jarvis'",
	}
}

func generatedCIWorkflowTokens() []string {
	return []string{
		"generated/verification_graph.generated.json",
		"generated/github_quality_workflow.generated.yml",
		"generated/gitlab_quality.generated.yml",
		"generated/local_quality.generated.mk",
		"generated/bazel_quality.generated.bzl",
		"generated/verification_graph.schema.generated.json",
		"generated/verification_conformance.generated.json",
		"generated/verification_tests.generated.json",
		"generated/release_pipeline.generated.json",
		"docs/verification-graph.md",
		"git diff --exit-code -- generated .github/workflows/quality.yml docs/verification-graph.md",
	}
}
