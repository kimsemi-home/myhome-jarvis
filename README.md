# myhome-jarvis

`myhome-jarvis` is a local-first home operations system for a household.
It starts with deterministic dry-run home-control commands and grows toward
finance, commerce, storage, and closed-loop Linear-managed development.

## Constraints

- No Python, Node.js, TypeScript, or shell-scripted core logic.
- Allowed implementation languages: Go, Rust, Common Lisp, and Flutter.
- Flutter-required Dart is allowed.
- Go toolchain is pinned to 1.26.2.
- Default behavior is local-only and dry-run.
- Secrets, private data, raw finance data, local tokens, and lake files must
  never be committed.
- External SaaS integrations are optional and must have offline fallback.

## First commands

When the toolchains are installed, the first stable milestone is:

```sh
go run ./cmd/mhj version
go run ./cmd/mhj auth status
go run ./cmd/mhj security check
go run ./cmd/mhj command open-youtube '{}'
go run ./cmd/mhj command volume-set '{"level":30}'
go run ./cmd/mhj harness home
go run ./cmd/mhj linear status
go run ./cmd/mhj linear pull
go run ./cmd/mhj linear next
go run ./cmd/mhj repo status
go run ./cmd/mhj loop once
go run ./cmd/mhj loop status
go run ./cmd/mhj loop worker --cycles 1
go run ./cmd/mhj benchmark smoke
go run ./cmd/mhj codegen verify
go run ./cmd/mhj daemon
go run ./cmd/mhj daemon status
cargo test -p mhj-core recommendations
```

Common Lisp SSOT and Rust command validation are scaffolded for:

```sh
sbcl --script lisp/scripts/validate-ssot.lisp
sbcl --script lisp/scripts/codegen.lisp
cargo test --workspace
cargo test -p mhj-core benchmark_smoke -- --nocapture
cargo test -p mhj-core household
cd apps/flutter && flutter test && flutter analyze
```

## Repository map

- `cmd/mhj`: Go CLI entrypoint.
- `internal/auth`: local LAN bearer-token management.
- `internal/commands`: dry-run command planning and validation.
- `internal/daemon`: local API, auth enforcement, metrics, and bounded request events.
- `internal/supervisor`: daemon process state and status checks.
- `internal/security`: forbidden language, secret, and private-data checks.
- `internal/linear`: Linear offline status and local queue.
- `internal/repo`: Git worktree status inspection for closed-loop safety.
- `internal/domain`: read-only finance, commerce, household, storage, and recommendation summaries.
- `internal/orchestrator`: one-shot checkpoint loop foundation.
- `internal/scheduler`: heartbeat, backoff, rate-limit, and recovery state for bounded loop workers.
- `apps/flutter`: Dart-only Flutter local client with daemon snapshot loading.
- `crates/mhj-core/src/benchmark.rs`: fixture-pipeline benchmark smoke tests.
- `crates/mhj-core/src/household.rs`: fixture-only user/spouse/household scope aggregation.
- `crates/mhj-core/src/recommendations.rs`: fixture-only recommendation scoring skeleton.
- `lisp/ssot`: executable source of truth.
- `generated`: deterministic artifacts emitted from SSOT.
- `.github/workflows`: hash-scoped GitHub Actions quality gates.
- `crates`: Rust command and harness foundations.
- `fixtures`: deterministic JSONL inputs.
- `docs`: working log, architecture notes, ADRs, and backlog.
