# Verification Graph

Source: `lisp/ssot/verification-graph.lisp`

Expression: `(verify daemon (public-safety) (ssot) (go) (rust) (flutter) (release-check))`

Generated artifacts:
- `generated/verification_graph.generated.json`
- `generated/github_quality_workflow.generated.yml`
- `.github/workflows/quality.yml`
- `generated/gitlab_quality.generated.yml`
- `generated/local_quality.generated.mk`
- `generated/bazel_quality.generated.bzl`
- `generated/control_plane_verification.generated.json`
- `generated/verification_evidence.generated.json`
- `generated/pdca.generated.json`
- `generated/verification_graph.schema.generated.json`
- `generated/verification_conformance.generated.json`
- `generated/verification_tests.generated.json`
- `generated/release_pipeline.generated.json`
- `generated/assistant_vision.generated.json`
- `generated/codex_cost.generated.json`
- `generated/finance_consent.generated.json`
- `generated/monetization.generated.json`
- `generated/repo_factory.generated.json`
- `docs/verification-graph.md`

Backends:
- `github-actions` -> `generated/github_quality_workflow.generated.yml`
- `gitlab-ci` -> `generated/gitlab_quality.generated.yml`
- `local-makefile` -> `generated/local_quality.generated.mk`
- `bazel` -> `generated/bazel_quality.generated.bzl`

Conformance:
- schema: `generated/verification_graph.schema.generated.json`
- manifest: `generated/verification_conformance.generated.json`
- tests: `generated/verification_tests.generated.json`
- release: `generated/release_pipeline.generated.json`

| Unit | Kind | Cache | Evidence |
| --- | --- | --- | --- |
| `public-safety` | release-check | none | GitHub log + command exit code |
| `ssot` | conformance | ssot | GitHub log + command exit code |
| `go` | unit-test | go | GitHub log + command exit code |
| `rust` | integration-test | rust | GitHub log + command exit code |
| `flutter` | lint | flutter | GitHub log + command exit code |
