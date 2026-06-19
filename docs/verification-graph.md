# Verification Graph

Source: `lisp/ssot/verification-graph.lisp`

Expression: `(verify daemon (public-safety) (ssot) (go) (rust) (flutter) (release-check))`

Generated artifacts:
- `generated/verification_graph.generated.json`
- `generated/github_quality_workflow.generated.yml`
- `.github/workflows/quality.yml`
- `docs/verification-graph.md`

| Unit | Kind | Cache | Evidence |
| --- | --- | --- | --- |
| `public-safety` | release-check | none | GitHub log + command exit code |
| `ssot` | conformance | ssot | GitHub log + command exit code |
| `go` | unit-test | go | GitHub log + command exit code |
| `rust` | integration-test | rust | GitHub log + command exit code |
| `flutter` | lint | flutter | GitHub log + command exit code |
