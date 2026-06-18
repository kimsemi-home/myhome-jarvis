# myhome-jarvis

`myhome-jarvis` is a local-first home operations system for a household.
It starts with deterministic dry-run home-control commands and grows toward
finance, commerce, storage, and closed-loop Linear-managed development.

## Constraints

- No Python, Node.js, TypeScript, or shell-scripted core logic.
- Allowed implementation languages: Go, Rust, Common Lisp, and Flutter.
- Flutter-required Dart is allowed.
- Go toolchain is pinned to 1.26.2.
- Rust toolchain is pinned by `rust-toolchain.toml`.
- Default behavior is local-only and dry-run.
- Secrets, private data, raw finance data, local tokens, and lake files must
  never be committed.
- External SaaS integrations are optional and must have offline fallback.

## First commands

When the toolchains are installed, the first stable milestone is:

```sh
go run ./cmd/mhj version
go run ./cmd/mhj auth status
go run ./cmd/mhj audit status
go run ./cmd/mhj ci verify
go run ./cmd/mhj security check
go run ./cmd/mhj security history
go run ./cmd/mhj toolchain verify
go run ./cmd/mhj command open-youtube '{}'
go run ./cmd/mhj command open-netflix '{}'
go run ./cmd/mhj command volume-set '{"level":30}'
go run ./cmd/mhj connectors status
go run ./cmd/mhj agent-cluster status
go run ./cmd/mhj learning status
go run ./cmd/mhj evidence status
go run ./cmd/mhj evidence-quality status
go run ./cmd/mhj review status
go run ./cmd/mhj confidence status
go run ./cmd/mhj authority status
go run ./cmd/mhj harness home
go run ./cmd/mhj harness finance
go run ./cmd/mhj harness commerce
go run ./cmd/mhj linear status
go run ./cmd/mhj linear pull
go run ./cmd/mhj linear next
go run ./cmd/mhj repo status
go run ./cmd/mhj planner status
go run ./cmd/mhj loop once
go run ./cmd/mhj loop status
go run ./cmd/mhj loop worker --cycles 1
go run ./cmd/mhj benchmark smoke
go run ./cmd/mhj codegen verify
go run ./cmd/mhj quality status
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
cargo test -p mhj-commerce
cargo test -p mhj-finance
cargo test -p mhj-storage
cd apps/flutter && flutter test && flutter analyze
```

## Repository map

- `cmd/mhj`: Go CLI entrypoint.
- `internal/audit`: private command intent audit journal.
- `internal/auth`: local LAN bearer-token management.
- `internal/commands`: dry-run command planning and validation.
- `internal/connectors`: public-safe fixture-only connector readiness status.
- `internal/agentcluster`: public-safe agent cluster learning-loop status.
- `internal/learning`: private observation ledger and redacted learning status.
- `internal/evidence`: private Evidence Graph summarization and redacted status.
- `internal/evidencequality`: private evidence quality snapshot assessor and redacted reassessment debt status.
- `internal/review`: private human review capacity queue and redacted status.
- `internal/confidence`: external confidence cap status over local evidence.
- `internal/authority`: public-safe Reasoning RBAC and Domain ABAC status gate.
- `internal/translation`: private Translation Manifest and Loss Ledger status.
- `internal/controlplane`: private Control Plane Manifest status for local orchestration decisions.
- `internal/incidents`: private Incident Lifecycle status for classified failures, owner roles, and quarantine debt.
- `internal/daemon`: local API, auth enforcement, metrics, and bounded request events.
- `internal/supervisor`: daemon process state and status checks.
- `internal/security`: forbidden language, secret, and private-data checks.
- `internal/linear`: Linear offline status and local queue.
- `internal/planner`: generated planner task graph status.
- `internal/repo`: Git worktree status inspection for closed-loop safety.
- `internal/domain`: read-only finance, commerce, household, storage, and recommendation summaries.
- `internal/orchestrator`: one-shot checkpoint loop foundation.
- `internal/qualitylog`: private quality gate evidence journal.
- `internal/scheduler`: heartbeat, backoff, rate-limit, and recovery state for bounded loop workers.
- `apps/flutter`: Dart-only Flutter local client with daemon snapshot loading and fixture finance/purchases dashboards.
- `crates/mhj-commerce`: fixture-only purchase IR validation, spend summaries, and recurring purchase candidates.
- `crates/mhj-finance`: fixture-only finance IR validation, cashflow summaries, and subscription review candidates.
- `crates/mhj-core/src/benchmark.rs`: fixture-pipeline benchmark smoke tests.
- `crates/mhj-core/src/household.rs`: fixture-only user/spouse/household scope aggregation.
- `crates/mhj-core/src/recommendations.rs`: fixture-only recommendation scoring skeleton.
- `crates/mhj-storage`: Rust data lake manifest, safe path planning, raw JSONL writes, and fixture-only Parquet+Zstd writer/metadata reader.
- `lisp/ssot`: executable source of truth.
- `generated`: deterministic artifacts emitted from SSOT.
- `.github/workflows`: hash-scoped GitHub Actions quality gates.
- `crates`: Rust command and harness foundations.
- `fixtures`: deterministic JSONL inputs.
- `docs`: working log, architecture notes, ADRs, and backlog.
