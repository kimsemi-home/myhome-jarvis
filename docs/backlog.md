# Backlog

## P0

- [x] Enforce no-Python language policy.
  - Linear: offline
  - Acceptance: `go run ./cmd/mhj security check` reports no Python, Node.js, TypeScript, secret, or private-data tracked-file risks.
  - Validation: Go security check.

- [x] Add Go `mhj` CLI skeleton.
  - Linear: offline
  - Acceptance: `version`, `security check`, `command`, `harness home`, `linear status`, `loop once`, and `quality` commands are present.
  - Validation: `go test ./...`

- [x] Add Common Lisp executable SSOT.
  - Linear: offline
  - Acceptance: `sbcl --script lisp/scripts/validate-ssot.lisp` and `sbcl --script lisp/scripts/codegen.lisp` complete deterministically.
  - Validation: generated artifacts match SSOT.

- [x] Add Rust command validation core.
  - Linear: offline
  - Acceptance: `cargo test -p mhj-command` validates YouTube, OTT, volume, display, and unsafe URL cases.
  - Validation: Rust tests.

- [x] Add home-control harness.
  - Linear: offline
  - Acceptance: `go run ./cmd/mhj harness home` passes deterministic dry-run cases.
  - Validation: Go harness command.

## P1

- [x] Add localhost-only Go daemon.
- [x] Add Linear GraphQL client in Go.
- [x] Add closed-loop orchestrator checkpoint evidence.
- [x] Add finance fixture IR in Rust.
  - Validation: `cargo test -p mhj-core finance`
- [x] Add commerce purchase IR in Rust.
  - Validation: `cargo test -p mhj-core commerce`
- [x] Add Parquet+Zstd-ready storage skeleton.
  - Validation: `cargo test -p mhj-core storage`

## P2

- [x] Add Flutter app skeleton after daemon and harness core are stable.
  - Acceptance: `apps/flutter` shows status, command, Linear, and storage tabs without platform runner files, and can load the same snapshot shape from daemon endpoints.
  - Validation: `cd apps/flutter && flutter test && flutter analyze`
- [x] Add benchmark smoke tests.
  - Acceptance: `go run ./cmd/mhj benchmark smoke` runs the core Rust fixture pipeline benchmark smoke test.
  - Validation: `cargo test -p mhj-core benchmark_smoke -- --nocapture`

## P3

- [x] Add fixture-only recommendation scoring skeleton.
  - Acceptance: Rust ranks cash buffer, subscription review, and recurring purchase review recommendations from local fixtures only; Go exposes the recommendation summary through local daemon surfaces; Flutter shows the ranked items in an Optimize tab.
  - Validation: `cargo test -p mhj-core recommendations`; `go test ./internal/domain ./internal/daemon`; `cd apps/flutter && flutter test && flutter analyze`.

- [x] Add fixture-only household view switching.
  - Acceptance: Rust and Go aggregate local finance and commerce fixture data into User, Spouse, and Household scopes; daemon exposes scope summaries; Flutter shows a segmented scope switcher.
  - Validation: `cargo test -p mhj-core household`; `go test ./internal/domain ./internal/daemon`; `cd apps/flutter && flutter test && flutter analyze`.

- [x] Add bounded scheduler heartbeat and recovery state.
  - Acceptance: `mhj loop worker --cycles 1` records private heartbeat/checkpoint state; `mhj loop status` and daemon `GET /loop/status` expose backoff, rate-limit, heartbeat, and recovery metadata without an unbounded loop.
  - Validation: `go test ./internal/scheduler ./internal/daemon`; full `mhj quality`.

- [x] Split GitHub Actions into hash-scoped unit caches.
  - Acceptance: SSOT, Go, Rust, and Flutter jobs each use a unit hash cache; cache hits skip heavy setup/tests; generated artifacts are verified on SSOT cache misses.
  - Validation: `mhj codegen verify`; full `mhj quality`; GitHub Actions run.

## P4

- [x] Add repository status inspection for closed-loop safety.
  - Acceptance: Go inspects Git branch/head/dirty state with repository-relative paths; daemon exposes `GET /repo/status`; Flutter status shows clean or dirty repository state; private ignored paths remain relative.
  - Validation: `go test ./internal/repo ./internal/daemon`; `cd apps/flutter && flutter test`; full quality gate.

- [x] Add explicit home-control command execution boundary.
  - Acceptance: dry-run remains default; CLI execution requires `MYHOME_EXECUTE=true`; daemon execution requires daemon execute mode and request `execute=true`; execution uses argv arrays only and allows only `open`, `osascript`, and `pmset`; non-macOS platforms skip safely.
  - Validation: `go test ./internal/commands ./internal/daemon`; full quality gate.

- [x] Add local LAN bearer-token management.
  - Acceptance: CLI can report token status, create a private local token, and rotate it; daemon non-localhost requests require `Authorization: Bearer`; Flutter daemon client can attach an optional Bearer token; SSOT records the LAN token policy.
  - Validation: `go test ./internal/auth ./internal/daemon`; `cd apps/flutter && flutter test`; `mhj auth status`; full quality gate.

- [x] Add bounded daemon request event log.
  - Acceptance: daemon records only method, path, status, duration, timestamp, and coarse error category in a 100-event in-memory buffer; `GET /events` returns recent events; `GET /metrics` exposes `event_count`; Flutter Status shows the count.
  - Validation: `go test ./internal/daemon`; `cd apps/flutter && flutter test`; full quality gate.

## P5

- [x] Add daemon process supervision state.
  - Acceptance: daemon writes private supervisor state only after successfully binding; `mhj daemon status` and `GET /supervisor/status` expose recorded pid/address/version, repo-relative state path, process liveness, token-free health probe, and stale detection; Flutter Status shows supervisor reachability.
  - Validation: `go test ./internal/supervisor ./internal/daemon`; `go run ./cmd/mhj daemon status`; `cd apps/flutter && flutter test`; full quality gate.

## P6

- [x] Add redacted command intent audit journal.
  - Acceptance: CLI and daemon command intents append private JSONL audit events with command/source/dry-run/execute gate/count/success/error category only; payloads, argv arrays, URLs, headers, tokens, raw errors, and local absolute paths are not recorded; `mhj audit status`, daemon `GET /audit/status`, and Flutter Status expose the count.
  - Validation: `go test ./internal/audit ./internal/daemon`; `go run ./cmd/mhj audit status`; `cd apps/flutter && flutter test`; full quality gate.

## P7

- [x] Add redacted quality gate evidence journal.
  - Acceptance: `mhj quality` appends a private JSONL run summary with overall status, duration, step count, pass/fail/skip counts, and step names/statuses only; command argv, command output, raw test output, environment variables, tokens, and local absolute paths are not recorded; `mhj quality status`, daemon `GET /quality/status`, and Flutter Status expose the last run.
  - Validation: `go test ./internal/qualitylog ./internal/daemon`; `go run ./cmd/mhj quality status`; `cd apps/flutter && flutter test`; full quality gate.

## P8

- [x] Add generated planner task graph.
  - Acceptance: Common Lisp SSOT owns the planner task graph, Linear templates, quality requirement, and external-write boundary; codegen emits `generated/planner.generated.json`; `mhj planner status`, daemon `GET /planner/status`, and Flutter Status expose ready/blocked counts and the next local task without writing to Linear.
  - Validation: `go test ./internal/planner ./internal/daemon`; `go run ./cmd/mhj planner status`; `cd apps/flutter && flutter test`; full quality gate.

## P9

- [x] Add dedicated Rust storage crate boundary.
  - Acceptance: `crates/mhj-storage` is part of the Cargo workspace; it emits deterministic data lake manifests for raw/bronze/silver/gold layers, validates repo-relative lake and partition paths, marks curated layers as Parquet+Zstd plans without claiming Parquet writes, and provides a raw JSONL writer smoke path for local fixtures.
  - Validation: `cargo test -p mhj-storage`; `cargo test --workspace`; full quality gate.

## P10

- [x] Add dedicated Rust finance crate boundary.
  - Acceptance: `crates/mhj-finance` is part of the Cargo workspace; it parses and validates fixture-only transaction IR, computes cashflow and owner summaries, and identifies subscription review candidates without requesting credentials or executing finance/subscription actions.
  - Validation: `cargo test -p mhj-finance`; `cargo test --workspace`; full quality gate.

## P11

- [x] Add dedicated Rust commerce crate boundary.
  - Acceptance: `crates/mhj-commerce` is part of the Cargo workspace; it parses and validates fixture-only purchase IR, computes spend, owner, and merchant summaries, and identifies recurring purchase review candidates without scraping, credentials, or purchase automation.
  - Validation: `cargo test -p mhj-commerce`; `cargo test --workspace`; full quality gate.

## P12

- [x] Add finance and commerce fixture harnesses.
  - Acceptance: `mhj harness finance` and `mhj harness commerce` validate deterministic fixture summaries from local data only; daemon `/harness/run` accepts `finance` and `commerce`; full quality gate includes both harnesses.
  - Validation: `go test ./internal/commands ./internal/daemon`; `go run ./cmd/mhj harness finance`; `go run ./cmd/mhj harness commerce`; full quality gate.

## P13

- [x] Add fixture-only Parquet+Zstd curated writer.
  - Acceptance: `crates/mhj-storage` materializes finance and commerce JSONL fixtures into bronze/silver/gold Parquet files with Zstd compression; raw layer curated writes are rejected; written files use repo-relative lake paths and metadata proves row count and compression.
  - Validation: `cargo test -p mhj-storage`; `cargo test --workspace`; full quality gate.

## P14

- [x] Add fixture-only Parquet metadata reader.
  - Acceptance: `crates/mhj-storage` can inspect finance and commerce curated Parquet fixture files through repo-relative lake paths; reader reports row count, row group count, column count, schema version, and Zstd compression without exposing row contents; raw layer curated reads are rejected.
  - Validation: `cargo test -p mhj-storage`; `cargo test --workspace`; full quality gate.

## P15

- [x] Add fixture-only card usage review recommendations.
  - Acceptance: Rust, Go, daemon summaries, and Flutter Optimize surface include a card-linked spend review recommendation derived only from local fixture card-linked debit records; recommendation is review-only and does not expose card IDs or execute card actions.
  - Validation: `cargo test -p mhj-core recommendations`; `cargo test -p mhj-finance`; `go test ./internal/domain ./internal/commands ./internal/daemon`; `cd apps/flutter && flutter test`; full quality gate.

## P16

- [x] Add explicit OTT shortcut command buttons.
  - Acceptance: SSOT, Go command registry, home harness, daemon command specs, and Flutter command UI expose dry-run shortcuts for Netflix, Disney+, TVING, Wavve, and Coupang Play while retaining the generic `open_ott` command and execution boundary.
  - Validation: `sbcl --script lisp/scripts/validate-ssot.lisp`; `go test ./internal/commands ./internal/daemon`; `go run ./cmd/mhj harness home`; `cd apps/flutter && flutter test`; full quality gate.

## P17

- [x] Add fixture-only Flutter finance dashboard.
  - Acceptance: Flutter parses finance totals, subscription spend, card-linked debit totals, categories, and owner breakdown from daemon `/domain/summary`; the app exposes a dedicated Finance tab without requesting credentials or executing finance actions.
  - Validation: `cd apps/flutter && flutter test`; `cd apps/flutter && flutter analyze`; full quality gate.

## P18

- [x] Add fixture-only Flutter purchases dashboard.
  - Acceptance: Flutter parses commerce spend, recurring purchase candidates, categories, and owner spend breakdown from daemon `/domain/summary`; the app exposes a dedicated Purchases tab without scraping, credentials, or purchase automation.
  - Validation: `cd apps/flutter && flutter test`; `cd apps/flutter && flutter analyze`; full quality gate.

## P19

- [x] Add structured fixture-only recommendation UI.
  - Acceptance: Flutter parses recommendation kind, rationale, score, estimated monthly amount, and evidence count from daemon `/domain/summary`; the Optimize tab renders purchase, subscription, card, and cash-buffer recommendations as read-only review signals without action buttons or external writes.
  - Validation: `cd apps/flutter && flutter test`; `cd apps/flutter && flutter analyze`; full quality gate.
