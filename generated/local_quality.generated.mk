# Generated from lisp/ssot/verification-graph.lisp.
SHELL := /bin/bash
.SHELLFLAGS := -euo pipefail -c
.ONESHELL:
.PHONY: verify verify-public-safety verify-ssot verify-go verify-rust verify-flutter

verify: verify-public-safety verify-ssot verify-go verify-rust verify-flutter

verify-public-safety:
	go run ./cmd/mhj security check
	go run ./cmd/mhj security history

verify-ssot:
	sbcl --script lisp/scripts/validate-ssot.lisp
	sbcl --script lisp/scripts/codegen.lisp
	test -s generated/verification_graph.generated.json
	test -s generated/github_quality_workflow.generated.yml
	test -s generated/gitlab_quality.generated.yml
	test -s generated/local_quality.generated.mk
	test -s generated/bazel_quality.generated.bzl
	test -s generated/control_plane_verification.generated.json
	test -s generated/verification_evidence.generated.json
	test -s generated/pdca.generated.json
	test -s generated/verification_graph.schema.generated.json
	test -s generated/verification_conformance.generated.json
	test -s generated/verification_tests.generated.json
	test -s generated/release_pipeline.generated.json
	test -s generated/finance_consent.generated.json
	test -s generated/codex_cost.generated.json
	test -s generated/codex_sustainability.generated.json
	test -s generated/context_pack.generated.json
	test -s generated/media_readiness.generated.json
	test -s generated/monetization.generated.json
	test -s generated/repo_factory.generated.json
	git diff --exit-code -- generated .github/workflows/quality.yml docs/verification-graph.md

verify-go:
	go run ./cmd/mhj ci verify
	go run ./cmd/mhj verification verify
	go run ./cmd/mhj verification evidence
	go run ./cmd/mhj pdca status
	go run ./cmd/mhj control-plane verify
	go run ./cmd/mhj toolchain verify
	go run ./cmd/mhj code-shape status
	go run ./cmd/mhj harness home
	go run ./cmd/mhj harness finance
	go run ./cmd/mhj harness commerce
	go test ./...
	go vet ./...
	unformatted="$$(gofmt -l cmd internal)"
	if [ -n "$$unformatted" ]; then
	  echo "$$unformatted"
	  exit 1
	fi

verify-rust:
	cargo test --workspace
	cargo test -p mhj-core benchmark_smoke -- --nocapture
	cargo fmt --check
	cargo clippy --workspace -- -D warnings

verify-flutter:
	cd apps/flutter
	flutter test
	flutter analyze
