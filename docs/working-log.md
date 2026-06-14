# Working Log

## 2026-06-14 20:57 local

- Linear issue: local continuation, no external Linear writes executed.
- Mode: online-capable, local-only changes in this pass.
- Task: Split GitHub Actions into hash-scoped unit caches and confirm SSOT/generated boundaries.
- Files touched: `.github/workflows/quality.yml`, `.gitignore`, `cmd/mhj/main.go`, `README.md`, `docs/backlog.md`, `docs/ci.md`, `docs/ssot.md`, `docs/working-log.md`.
- Changes: added `mhj codegen verify`; documented current SSOT domain boundaries; split Actions into SSOT, Go, Rust, and Flutter unit jobs with hash-keyed cache markers; kept generated artifact verification in the SSOT job; opted Actions into Node 24 and set Go cache dependency path to `go.mod`.
- Validation after: `go1.26.2 run ./cmd/mhj codegen verify` passed; `MHJ_GO=$HOME/go/bin/go1.26.2 MHJ_GOFMT=$HOME/sdk/go1.26.2/bin/gofmt $HOME/go/bin/go1.26.2 run ./cmd/mhj quality` passed; workflow YAML parsed; public safety scans passed; generated artifacts had no diff after codegen; first split GitHub Actions run succeeded and warmed unit caches.
- External-write note: no Linear mutation, purchase, finance, card, investment, or other external write was executed.
- Next: commit, push, and verify hash-scoped GitHub Actions with `gh`.

## 2026-06-14 20:46 local

- Linear issue: local continuation, no external Linear writes executed.
- Mode: online-capable, local-only changes in this pass.
- Task: Add bounded scheduler heartbeat, backoff, checkpoint, and recovery state.
- Files touched: `internal/scheduler/scheduler.go`, `internal/scheduler/scheduler_test.go`, `cmd/mhj/main.go`, `internal/daemon/server.go`, `internal/daemon/server_test.go`, `lisp/ssot/scheduler.lisp`, `lisp/ssot/codegen.lisp`, `lisp/ssot/package.lisp`, `lisp/ssot/myhome-jarvis.asd`, `lisp/scripts/load-ssot.lisp`, `generated/scheduler.generated.json`, `README.md`, `docs/architecture.md`, `docs/backlog.md`, `docs/closed-loop.md`, `docs/scheduler.md`, `docs/working-log.md`.
- Changes: added Go scheduler policy/state with heartbeat, bounded backoff, rate-limit next-run metadata, private state persistence, and crash recovery; added `mhj loop status`, bounded `mhj loop worker --cycles`, and daemon `GET /loop/status`; added SSOT scheduler policy.
- Validation after: `go1.26.2 test ./internal/scheduler ./internal/daemon` passed; `go1.26.2 run ./cmd/mhj loop status` passed; `go1.26.2 run ./cmd/mhj loop worker --cycles 1` passed with private state/checkpoint persistence; `MHJ_GO=$HOME/go/bin/go1.26.2 MHJ_GOFMT=$HOME/sdk/go1.26.2/bin/gofmt $HOME/go/bin/go1.26.2 run ./cmd/mhj quality` passed.
- External-write note: no Linear mutation, purchase, finance, card, investment, or other external write was executed.
- Next: run public safety scans, then commit and push if clean.

## 2026-06-14 20:37 local

- Linear issue: local continuation, no external Linear writes executed.
- Mode: online-capable, local-only changes in this pass.
- Task: Add fixture-only User, Spouse, and Household view switching.
- Files touched: `crates/mhj-core/src/household.rs`, `crates/mhj-core/src/lib.rs`, `internal/domain/summary.go`, `internal/domain/summary_test.go`, `internal/daemon/server.go`, `internal/daemon/server_test.go`, `apps/flutter/lib/snapshot.dart`, `apps/flutter/lib/daemon_client.dart`, `apps/flutter/lib/main.dart`, `apps/flutter/test/daemon_client_test.dart`, `apps/flutter/test/widget_test.dart`, `lisp/ssot/household.lisp`, `lisp/ssot/codegen.lisp`, `lisp/ssot/package.lisp`, `lisp/ssot/myhome-jarvis.asd`, `lisp/scripts/load-ssot.lisp`, `generated/household.generated.json`, `README.md`, `docs/architecture.md`, `docs/backlog.md`, `docs/flutter.md`, `docs/household.md`, `docs/working-log.md`.
- Changes: added Rust household scope aggregation over finance and commerce fixtures; added Go owner breakdown and household summary projection; exposed daemon `GET /household/summary`; added Flutter Household tab with a segmented User, Spouse, Household switcher.
- Validation after: `cargo test -p mhj-core household` passed; `go1.26.2 test ./internal/domain ./internal/daemon` passed; `flutter test` passed; `MHJ_GO=$HOME/go/bin/go1.26.2 MHJ_GOFMT=$HOME/sdk/go1.26.2/bin/gofmt $HOME/go/bin/go1.26.2 run ./cmd/mhj quality` passed.
- External-write note: no account, finance, purchase, subscription, card, investment, Linear mutation, or other external write was executed.
- Next: run full quality and public safety scans, then commit and push if clean.

## 2026-06-14 20:26 local

- Linear issue: local continuation, no external Linear writes executed.
- Mode: online-capable, local-only changes in this pass.
- Task: Add fixture-only recommendation scoring skeleton and local UI surface.
- Files touched: `crates/mhj-core/src/recommendations.rs`, `crates/mhj-core/src/lib.rs`, `internal/domain/summary.go`, `internal/domain/summary_test.go`, `internal/daemon/server.go`, `internal/daemon/server_test.go`, `apps/flutter/lib/snapshot.dart`, `apps/flutter/lib/daemon_client.dart`, `apps/flutter/lib/main.dart`, `apps/flutter/test/daemon_client_test.dart`, `apps/flutter/test/widget_test.dart`, `lisp/ssot/recommendations.lisp`, `lisp/ssot/codegen.lisp`, `lisp/ssot/package.lisp`, `lisp/ssot/myhome-jarvis.asd`, `generated/recommendations.generated.json`, `README.md`, `docs/architecture.md`, `docs/backlog.md`, `docs/flutter.md`, `docs/recommendations.md`, `docs/working-log.md`.
- Changes: added Rust scoring for cash buffer, subscription review, and recurring purchase review recommendations from local fixtures; added SSOT recommendation policy and generated artifact; extended Go domain summary and daemon `GET /recommendations/summary`; added Flutter Optimize tab fed by daemon snapshot data.
- Validation after: `cargo test -p mhj-core recommendations` passed; `go1.26.2 test ./internal/domain ./internal/daemon` passed; `flutter test` passed; `MHJ_GO=$HOME/go/bin/go1.26.2 MHJ_GOFMT=$HOME/sdk/go1.26.2/bin/gofmt $HOME/go/bin/go1.26.2 run ./cmd/mhj quality` passed.
- External-write note: no purchases, subscription changes, card actions, transfers, investment trades, Linear mutations, or other external writes were executed.
- Next: run full quality and public safety scans, then commit and push if clean.

## 2026-06-14 20:02 local

- Linear issue: local continuation, no external Linear writes executed.
- Mode: online-capable, local-only changes in this pass.
- Task: Add read-only local finance, commerce, and storage summaries to the daemon and surface them in Flutter.
- Files touched: `internal/domain/summary.go`, `internal/domain/summary_test.go`, `internal/daemon/server.go`, `internal/daemon/server_test.go`, `apps/flutter/lib/daemon_client.dart`, `apps/flutter/test/daemon_client_test.dart`, `README.md`, `docs/architecture.md`, `docs/flutter.md`, `docs/storage.md`, `docs/finance-domain.md`, `docs/commerce-domain.md`, `docs/working-log.md`.
- Changes: added `internal/domain` fixture summary builders; added daemon `GET /domain/summary`; added daemon tests for finance net, commerce recurring candidates, and storage policy; extended Flutter daemon snapshot loading to read `/domain/summary`; rendered finance, commerce, and storage status in the Storage tab.
- Validation after: `go1.26.2 test ./internal/domain ./internal/daemon` passed; `flutter test` passed with the updated daemon summary fixture; `dart format lib test` passed.
- Next: add richer Flutter payload editing for `/intent` previews, or run a live daemon/UI smoke after platform runner setup.

## 2026-06-14 13:03 local

- Linear issue: local continuation, no external Linear writes executed.
- Mode: online-capable, local-only changes in this pass.
- Task: Add Flutter command dry-run previews through daemon `POST /intent`.
- Files touched: `apps/flutter/lib/snapshot.dart`, `apps/flutter/lib/daemon_client.dart`, `apps/flutter/lib/main.dart`, `apps/flutter/test/daemon_client_test.dart`, `apps/flutter/test/widget_test.dart`, `docs/flutter.md`, `docs/home-control.md`, `docs/working-log.md`.
- Changes: added `CommandPlan` and `CommandInvocation` models; extended `DaemonSnapshotClient` with `dryRun`; POSTs command, payload, and `execute=false` to `/intent`; connected command rows to a dry-run plan dialog; added daemon client and widget tests for the preview flow.
- Validation after: `dart format lib test` passed; `flutter test` passed with 5 tests; `flutter analyze` passed.
- Next: optionally run a live daemon/UI smoke once a browser/platform runner is added, or add richer plan detail and command payload editing before execution is ever enabled.

## 2026-06-14 13:00 local

- Linear issue: local continuation, no external Linear writes executed.
- Mode: online-capable, local-only changes in this pass.
- Task: Connect the Flutter skeleton to the local daemon snapshot surface.
- Files touched: `apps/flutter/lib/main.dart`, `apps/flutter/lib/snapshot.dart`, `apps/flutter/lib/daemon_client.dart`, `apps/flutter/test/widget_test.dart`, `apps/flutter/test/daemon_client_test.dart`, `README.md`, `docs/backlog.md`, `docs/architecture.md`, `docs/flutter.md`, `docs/working-log.md`.
- Changes: split Flutter snapshot models out of `main.dart`; added `DaemonSnapshotClient` for `/health`, `/commands`, `/linear/status`, and `/metrics`; made the UI load snapshots asynchronously with deterministic fallback; added an HTTP-backed daemon client test; kept widget tests stable with the static snapshot client.
- Validation after: `dart format lib test` passed; `flutter test` passed with widget and daemon client tests; `flutter analyze` passed.
- Next: add daemon intent execution previews in Flutter, then consider explicit user confirmation for Linear backlog seeding.

## 2026-06-14 12:56 local

- Linear issue: local continuation, no external Linear writes executed.
- Mode: online-capable, local-only changes in this pass.
- Task: Add P2 Flutter app skeleton and include Flutter test/analyze in the quality gate.
- Files touched: `apps/flutter/pubspec.yaml`, `apps/flutter/pubspec.lock`, `apps/flutter/lib/main.dart`, `apps/flutter/test/widget_test.dart`, `apps/flutter/README.md`, `cmd/mhj/main.go`, `README.md`, `docs/backlog.md`, `docs/architecture.md`, `docs/flutter.md`, `docs/working-log.md`.
- Changes: installed Flutter 3.44.2 with Dart 3.12.2 through Homebrew; disabled Flutter analytics; added a Dart-only Flutter local client skeleton with status, command, Linear, and storage tabs; added widget tests; updated `mhj quality` so Flutter commands run from `apps/flutter`; marked the P2 Flutter skeleton backlog item complete.
- Validation after: `flutter test` passed; `flutter analyze` passed; `MHJ_GO=$HOME/go/bin/go1.26.2 MHJ_GOFMT=$HOME/sdk/go1.26.2/bin/gofmt $HOME/go/bin/go1.26.2 run ./cmd/mhj quality` passed with Flutter test/analyze included.
- Next: expand the local client to call the daemon endpoints, or seed Linear backlog only after explicit confirmation for external writes.

## 2026-06-14 12:45 local

- Linear issue: local continuation, no external Linear writes executed.
- Mode: online-capable, local-only changes in this pass.
- Task: Add P2 benchmark smoke tests for core Rust fixture pipeline and expose them through the Go CLI quality surface.
- Files touched: `cmd/mhj/main.go`, `crates/mhj-core/src/lib.rs`, `crates/mhj-core/src/benchmark.rs`, `README.md`, `docs/backlog.md`, `docs/performance.md`, `docs/working-log.md`.
- Changes: added `mhj-core::benchmark::run_benchmark_smoke`; added `benchmark_smoke_runs_core_fixture_pipeline` test over finance JSONL parsing, cashflow summary, commerce JSONL parsing, recurring candidate detection, and storage plan generation; added `mhj benchmark smoke`; added an explicit benchmark smoke step to `mhj quality`; marked the P2 benchmark smoke backlog item complete.
- Validation after: `cargo test -p mhj-core benchmark_smoke -- --nocapture` passed; `go1.26.2 run ./cmd/mhj benchmark smoke` passed.
- Next: add Flutter app skeleton after checking toolchain availability and generated-file policy, or seed Linear backlog only after explicit confirmation for external writes.

## 2026-06-14 12:42 local

- Linear issue: local continuation, no external Linear writes executed.
- Mode: online-capable, local-only changes in this pass.
- Task: Complete P1 Rust finance fixture IR, commerce purchase IR, and Parquet+Zstd-ready storage skeleton.
- Files touched: `crates/mhj-core/Cargo.toml`, `crates/mhj-core/src/lib.rs`, `crates/mhj-core/src/finance.rs`, `crates/mhj-core/src/commerce.rs`, `crates/mhj-core/src/storage.rs`, `fixtures/finance_transactions.jsonl`, `fixtures/commerce_purchases.jsonl`, `lisp/ssot/storage.lisp`, `generated/storage.generated.json`, `docs/finance-domain.md`, `docs/commerce-domain.md`, `docs/storage.md`, `docs/architecture.md`, `docs/backlog.md`, `docs/working-log.md`, `Cargo.lock`.
- Changes: added Rust JSONL parsing and validation primitives; added finance transaction IR fixtures with cashflow summary; added commerce purchase IR fixtures with recurring-purchase candidate detection; added storage dataset planning for raw JSONL and bronze/silver/gold Parquet+Zstd outputs; recorded finance and commerce datasets in SSOT storage policy; marked the three P1 Rust/storage backlog items complete.
- Validation after: `cargo test -p mhj-core finance` passed; `cargo test -p mhj-core commerce` passed; `cargo test -p mhj-core storage` passed; `MHJ_GO=$HOME/go/bin/go1.26.2 MHJ_GOFMT=$HOME/sdk/go1.26.2/bin/gofmt $HOME/go/bin/go1.26.2 run ./cmd/mhj quality` passed; `git diff --check` passed; trailing-whitespace scan found no matches; forbidden Python/Node/TypeScript file pattern scan found no matches.
- Result: P1 local daemon, Linear client, checkpoint evidence, finance IR, commerce IR, and storage skeleton are now all implemented and locally verified.
- Next: decide whether to seed the project backlog into Linear, then move to P2 Flutter skeleton or benchmark smoke tests.

## 2026-06-14 12:36 local

- Linear issue: Chrome-authorized Linear API key for the selected team, stored only in ignored private storage.
- Mode: online
- Task: Connect Linear through Chrome with least-privilege API key and pin the Go project version to 1.26.2.
- Files touched: `go.mod`, `.go-version`, `cmd/mhj/main.go`, `lisp/ssot/project.lisp`, `generated/commands.generated.json`, `README.md`, `docs/working-log.md`.
- Changes: created a Linear API key with selected read/create-comment/create-issue permissions scoped to the selected team; saved it to `data/private/linear-token.txt` with `0600` permissions; updated the Go module directive and local version file to `1.26.2`; added SSOT `go_version`; taught the quality gate to honor `MHJ_GO` and `MHJ_GOFMT` for exact-toolchain validation.
- Validation after: installed and verified `go1.26.2 darwin/arm64`; `MHJ_GO=$HOME/go/bin/go1.26.2 MHJ_GOFMT=$HOME/sdk/go1.26.2/bin/gofmt $HOME/go/bin/go1.26.2 run ./cmd/mhj quality` passed with Flutter skipped because `apps/flutter` is not started; `go1.26.2 run ./cmd/mhj linear status` reported online; `linear pull` read active issues; `linear next` selected the next active issue.
- External-write note: no Linear issue creation, comments, or state transitions were executed in this pass.
- Next: seed project-specific Linear backlog only after explicit confirmation, then add finance/commerce fixture IR in Rust.

## 2026-06-14 12:23 local

- Linear issue: offline continuation, no Linear token configured.
- Mode: offline
- Task: Add Linear pull/next/comment/transition/create-from-backlog command surfaces with structured offline fallback.
- Files touched: `cmd/mhj/main.go`, `internal/daemon/server.go`, `internal/linear/status.go`, `internal/linear/issues.go`, `internal/linear/issues_test.go`, `lisp/ssot/linear.lisp`, `generated/linear.generated.json`, `README.md`, `docs/linear-workflow.md`, `docs/backlog.md`, `docs/working-log.md`.
- Plan: split the Linear package into reusable GraphQL and issue-operation boundaries, keep token handling private, record unsynced actions in `data/private/linear-offline-queue.jsonl`, update docs/SSOT, and run full quality.
- Validation before: `go run ./cmd/mhj quality` passed with Flutter skipped because `apps/flutter` is not started; current Linear status reports offline with `synced=false`.
- Changes: added reusable variable-based Linear GraphQL calls; added `mhj linear pull`, `mhj linear next`, `mhj linear comment <issue-id> <message>`, `mhj linear transition <issue-id> <state>`, and `mhj linear create-from-backlog`; routed `mhj linear sync` and daemon `/linear/sync` through pull behavior; added structured offline queue payloads for comment, transition, and backlog seed actions; updated SSOT generated Linear policy and docs.
- Validation after: `go run ./cmd/mhj linear pull`, `linear next`, `linear comment`, `linear transition`, and `linear create-from-backlog` all reported offline with `synced=false` and wrote ignored private queue events; `go run ./cmd/mhj quality` passed; `git diff --check` passed; deterministic codegen SHA-256 check passed; forbidden Python/Node/TypeScript file scan remained clean.
- Result: Linear GraphQL client P1 item now has status, pull, next, comment, transition, create-from-backlog, sync, and offline fallback surfaces verified locally.
- Next: add finance/commerce fixture IR in Rust, then storage skeleton.

## 2026-06-14 12:04 local

- Linear issue: offline continuation, no Linear token configured.
- Mode: offline
- Task: Add P1 localhost daemon endpoints and direct Go Linear GraphQL status/offline foundation.
- Files touched: `cmd/mhj/main.go`, `internal/daemon/server.go`, `internal/daemon/server_test.go`, `internal/linear/status.go`, `internal/linear/status_test.go`, `internal/commands/registry.go`, `lisp/ssot/linear.lisp`, `generated/linear.generated.json`, `README.md`, `docs/architecture.md`, `docs/linear-workflow.md`, `docs/backlog.md`, `docs/working-log.md`.
- Plan: add standard-library Go daemon routes for health/version/commands/intent/harness/metrics and expand Linear status from placeholder to token-aware GraphQL status with safe offline queue fallback.
- Validation before: worktree contains bootstrap files only; `go`, `gofmt`, `cargo`, `rustc`, `sbcl`, and `flutter` are still missing on PATH; Homebrew is available but no toolchain install is assumed in this task.
- Changes: added localhost-only Go daemon routes for health/version/commands/intent/harness/linear/metrics; added CLI `daemon` and `linear sync`; replaced placeholder Linear status with direct GraphQL viewer/team status client using safe token loading; added daemon and Linear tests; installed Go, SBCL, and Rust via rustup so the core quality gate can run locally. A Homebrew `rust` attempt was stopped because it would install Python as a formula dependency; `python@3.14` was already required by existing `pipx`, so it was not force-removed.
- Validation after: `go test ./...` pass; `go vet ./...` pass; `gofmt -l` clean through `go run ./cmd/mhj quality`; `cargo test --workspace` pass; `cargo fmt --check` pass; `cargo clippy --workspace -- -D warnings` pass; `sbcl --script lisp/scripts/validate-ssot.lisp` pass; `sbcl --script lisp/scripts/codegen.lisp` pass; generated JSON SHA-256 values remained unchanged across consecutive codegen runs; `go run ./cmd/mhj security check` pass; `go run ./cmd/mhj harness home` pass; required command dry-runs produced deterministic argv; invalid volume, URL, and OTT commands failed safely; `go run ./cmd/mhj linear status` reported offline with `synced=false`; `go run ./cmd/mhj loop once` wrote an ignored private checkpoint; daemon smoke test on `127.0.0.1:3899` returned healthy `/health` and dry-run `/intent`.
- Result: P0 stable milestone and P1 daemon/checkpoint foundation are verified; Linear issue mutation/pull workflow, finance/commerce IR, storage, and Flutter remain future work.
- Next: implement Linear issue/comment mutations and next-issue pull with offline replay, then add finance/commerce fixture IR.

## 2026-06-14 11:52 local

- Linear issue: offline bootstrap, no Linear token configured.
- Mode: offline
- Task: Phase 0 audit and P0 bootstrap for language policy, CLI skeleton, executable SSOT, and command harness foundations.
- Files touched: repository metadata, docs, Go CLI/internal packages, Lisp SSOT, Rust crates, generated artifacts, fixtures.
- Plan: create the smallest runnable closed-loop foundation without Python, Node.js, or TypeScript.
- Validation before: root directory was empty; `git status` failed because `.git` did not exist; `go`, `cargo`, `rustc`, `sbcl`, and `flutter` were not present on PATH.
- Changes: initialized local Git metadata in the existing root and began bootstrapping policy, CLI, SSOT, harness, and offline Linear queue.
- Validation after: static forbidden-file scan found no Python, Node.js, or TypeScript source/dependency files; trailing-whitespace scan found no files; sensitive-path scan found only allowed `.env.example`; `go test ./...`, `cargo test --workspace`, `sbcl --script lisp/scripts/validate-ssot.lisp`, and `flutter test` could not run because their executables are missing on PATH.
- Result: P0 bootstrap files created; runtime validation blocked by missing Go, Rust/Cargo, SBCL, and Flutter toolchains in this shell.
- Next: install or expose the required toolchains on PATH, then run `go run ./cmd/mhj security check`, `go test ./...`, `cargo test --workspace`, `sbcl --script lisp/scripts/validate-ssot.lisp`, and `go run ./cmd/mhj harness home`.
