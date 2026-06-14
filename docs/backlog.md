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

## P20

- [x] Add explicit Flutter local-only network mode indicator.
  - Acceptance: Flutter derives a Network status metric from daemon `/health` and `/metrics`; localhost/default mode renders as `Local-only`, LAN-enabled mode renders as `LAN token-gated`, and the offline fallback remains local-only.
  - Validation: `cd apps/flutter && flutter test`; `cd apps/flutter && flutter analyze`; full quality gate.

## P21

- [x] Add read-only daemon LAN auth status surface.
  - Acceptance: daemon `GET /auth/status` exposes local token configured/missing state, repo-relative token path, file mode, and message without returning token contents; Flutter Status renders a `LAN Auth` metric from that surface.
  - Validation: `go test ./internal/daemon`; `cd apps/flutter && flutter test`; `cd apps/flutter && flutter analyze`; full quality gate.

## P22

- [x] Complete fallback Flutter home-control command buttons.
  - Acceptance: static/offline Flutter snapshots expose YouTube, OTT shortcuts, volume up/down/set, and display sleep dry-run commands so the local client keeps the same core home-control surface even before a daemon connection is available.
  - Validation: `cd apps/flutter && flutter test`; `cd apps/flutter && flutter analyze`; full quality gate.

## P23

- [x] Add fallback Flutter URL and search command controls.
  - Acceptance: static/offline Flutter snapshots expose `open-youtube-search`, `open-url`, and generic `open-ott` editable dry-run commands in addition to dedicated shortcuts, so daemon-unavailable clients can still preview YouTube search, safe URL open, and service-selected OTT intents.
  - Validation: `cd apps/flutter && flutter test`; `cd apps/flutter && flutter analyze`; full quality gate.

## P24

- [x] Add command SSOT to Go registry parity checks.
  - Acceptance: Go command tests load `generated/commands.generated.json` and fail on drift in command names, summaries, payload fields, OTT service allowlist, or generated URL targets, keeping Lisp SSOT artifacts and Go execution plans aligned.
  - Validation: `go test ./internal/commands`; full quality gate.

## P25

- [x] Add Git history public-safety gate.
  - Acceptance: `mhj security history` scans reachable Git commits for private identity markers, local absolute paths, forbidden language/dependency files, private/lake data paths except empty keep placeholders, sensitive-looking paths, secret-looking literals, and commit metadata issues without reporting raw matched secrets; CI always runs a full-history public-safety job before the hash-scoped unit summary can pass.
  - Validation: `go test ./internal/security`; `go run ./cmd/mhj security history`; full quality gate; GitHub Actions run.

## P26

- [x] Surface public-safety status in daemon and Flutter.
  - Acceptance: daemon `GET /security/status` exposes only aggregate current-tree and Git-history safety booleans, finding counts, and checked timestamp; Flutter Status renders a `Public Safety` metric from that endpoint and the offline fallback remains clear without exposing raw findings, matched content, or local roots.
  - Validation: `go test ./internal/security ./internal/daemon`; `cd apps/flutter && flutter test`; `cd apps/flutter && flutter analyze`; full quality gate.

## P27

- [x] Redact closed-loop checkpoint safety evidence.
  - Acceptance: `mhj loop once` and `mhj loop worker --cycles 1` use aggregate public-safety status for checkpoint decisions; loop output and private checkpoint evidence include redacted Linear summary and aggregate security status only, with no raw Linear viewer/team identities, raw security findings, local repository root, or absolute private paths.
  - Validation: `go test ./internal/linear ./internal/orchestrator ./internal/scheduler ./internal/security`; `go run ./cmd/mhj loop once`; `go run ./cmd/mhj loop worker --cycles 1`; full quality gate.

## P28

- [x] Align planner progress with completed local rails.
  - Acceptance: SSOT marks completed local planner rails as `completed`; `mhj planner status`, daemon `GET /planner/status`, and Flutter Status expose completed, ready, and external-write-gated counts; `next_task` is omitted once no local ready task remains; planner validation rejects unknown task statuses.
  - Validation: `sbcl --script lisp/scripts/validate-ssot.lisp`; `go run ./cmd/mhj codegen verify`; `go test ./internal/planner ./internal/daemon`; `cd apps/flutter && flutter test`; full quality gate.

## P29

- [x] Make local codegen verification working-tree aware.
  - Acceptance: `mhj codegen verify` snapshots the current `generated` tree, regenerates artifacts from Lisp, and fails only when regeneration changes generated files, so intentional SSOT/generated updates can be verified before commit while stale artifacts are still caught; `mhj quality` uses that verification step.
  - Validation: `go test ./cmd/mhj`; `go run ./cmd/mhj codegen verify`; full quality gate.

## P30

- [x] Surface external-write-gated planner task details.
  - Acceptance: `mhj planner status` and daemon `GET /planner/status` include read-only `blocked_external_write_tasks` metadata when local rails are complete and an external-write-gated step remains; `next_task` stays omitted for completed local work and no Linear mutation is executed.
  - Validation: `go test ./internal/planner ./internal/daemon`; `go run ./cmd/mhj planner status`; full quality gate.

## P31

- [x] Reflect planner gate details in Flutter Status.
  - Acceptance: Flutter parses daemon `blocked_external_write_tasks` and renders a read-only `Planner Gate` status metric with the first gated task id when local planner rails are complete; the UI does not add action buttons or execute Linear mutations.
  - Validation: `cd apps/flutter && flutter test test/daemon_client_test.dart`; `cd apps/flutter && flutter test`; `cd apps/flutter && flutter analyze`; full quality gate.

## P32

- [x] Redact default Linear status surfaces.
  - Acceptance: `mhj linear status`, daemon `GET /linear/status`, and Flutter Linear status rendering expose redacted Linear summary fields only; raw viewer/team identities, token source, and absolute private queue paths are not returned by default while internal sync logic can still use raw GraphQL status.
  - Validation: `go test ./internal/linear ./internal/daemon`; `go run ./cmd/mhj linear status`; `cd apps/flutter && flutter test test/daemon_client_test.dart`; full quality gate.

## P33

- [x] Complete dedicated Rust fixture harness boundary.
  - Acceptance: `crates/mhj-harness` validates the home-control dry-run matrix, service-specific OTT shortcuts, finance fixture totals/owner scopes/review-only candidates, and commerce fixture spend/merchant/recurring-purchase invariants through Rust crate boundaries; no external command, finance, commerce, scraping, credential, or Linear mutation is executed.
  - Validation: `cargo test -p mhj-command -p mhj-harness`; `go run ./cmd/mhj harness home`; `go run ./cmd/mhj harness finance`; `go run ./cmd/mhj harness commerce`; full quality gate.

## P34

- [x] Add bounded daemon HTTP resource defaults.
  - Acceptance: daemon `http.Server` instances use non-zero read-header, read, write, idle, and max-header-size limits by default, including when `New` receives a minimal config; localhost/LAN behavior and auth gates remain unchanged.
  - Validation: `go test ./internal/daemon`; full quality gate.

## P35

- [x] Surface redacted daemon runtime counters.
  - Acceptance: daemon `GET /metrics` exposes aggregate Go runtime counters for goroutine count, heap allocation bytes, heap system bytes, stack in-use bytes, and GC count without exposing local roots, tokens, request payloads, or raw process data; Flutter Status renders runtime and heap metrics when present.
  - Validation: `go test ./internal/daemon`; `cd apps/flutter && flutter test test/daemon_client_test.dart`; full quality gate.

## P36

- [x] Upgrade GitHub Actions maintained refs for Node 24.
  - Acceptance: workflow-owned uses of checkout, setup-go, and cache actions point at Node 24-capable releases; the manual force-to-Node24 environment opt-in is removed; unit hash caches still include the workflow file so action ref changes are verified once and unchanged units can later skip heavy work.
  - Validation: workflow YAML parses; full quality gate; GitHub Actions run.

## P37

- [x] Make storage fixture temp roots collision-safe.
  - Acceptance: `mhj-storage` curated fixture writer tests use per-process unique temporary roots even when Rust tests run in parallel, so one test cannot remove another test's generated Parquet fixture.
  - Validation: `cargo test -p mhj-storage`; full quality gate.

## P38

- [x] Redact default Linear issue operation surfaces.
  - Acceptance: `mhj linear sync`, `mhj linear pull`, `mhj linear next`, and daemon `POST /linear/sync` return operation summaries without raw issue descriptions, workspace URLs, team identities, Linear UUIDs, token source, or absolute private paths; internal GraphQL operations still use variables and can select next issues without mutating Linear.
  - Validation: `go test ./internal/linear ./internal/daemon`; `go run ./cmd/mhj linear next`; full quality gate.

## P39

- [x] Redact current-tree security report root.
  - Acceptance: `mhj security check` reports the checked root as `.` instead of the local checkout path while preserving repo-relative findings and the existing full-history security gate.
  - Validation: `go test ./internal/security`; `go run ./cmd/mhj security check`; full quality gate.

## P40

- [x] Redact default quality gate CLI output.
  - Acceptance: `mhj quality` prints only overall status plus step names/statuses by default; command argv, raw command output, raw test output, and local absolute paths stay out of stdout while internal pass/fail handling and the private redacted quality journal remain intact.
  - Validation: `go test ./cmd/mhj`; `go run ./cmd/mhj quality`; full quality gate.

## P41

- [x] Cancel superseded quality runs and harden domain summary public surface.
  - Acceptance: GitHub Actions cancels older in-progress quality runs for the same ref when a newer push arrives; daemon `GET /domain/summary` regression coverage proves the generated storage root remains repo-relative and does not leak the local checkout or home path.
  - Validation: workflow YAML parses; `go test ./internal/daemon`; full quality gate; GitHub Actions run.

## P42

- [x] Scan current working-tree contents before public commit.
  - Acceptance: `mhj security check` scans non-private current file contents for private identity markers, local absolute paths, and secret-looking literals before they enter Git history; findings remain redacted to repo-relative path, optional line, code, and coarse message.
  - Validation: `go test ./internal/security`; `go run ./cmd/mhj security check`; full quality gate; GitHub Actions run.

## P43

- [x] Record current content scanning in security SSOT.
  - Acceptance: Common Lisp SSOT emits generated security policy fields for current-content scanning, private-path skipping, private-identity scan, secret-literal scan, and non-reporting of matched secret contents; Go security tests read the generated artifact and fail on drift.
  - Validation: `sbcl --script lisp/scripts/validate-ssot.lisp`; `go run ./cmd/mhj codegen verify`; `go test ./internal/security`; full quality gate; GitHub Actions run.

## P44

- [x] Align Flutter offline command fallback with home-control surface.
  - Acceptance: Flutter static/offline fallback includes `volume-mute` and `mac-sleep` alongside the daemon-sourced command surface; widget tests prove both buttons render without daemon reachability.
  - Validation: `cd apps/flutter && flutter test test/widget_test.dart`; full quality gate; GitHub Actions run.

## P45

- [x] Guard Flutter fallback commands against SSOT drift.
  - Acceptance: Flutter tests read `generated/commands.generated.json` and fail when static/offline fallback command names or payload fields differ from the Lisp-owned command catalog.
  - Validation: `cd apps/flutter && flutter test test/snapshot_test.dart`; full quality gate; GitHub Actions run.

## P46

- [x] Include generated command catalog in Flutter CI cache key.
  - Acceptance: the Flutter hash-scoped GitHub Actions unit reruns when `generated/commands.generated.json` changes, because Flutter fallback tests read that artifact directly; unchanged Flutter/generated command hashes still skip heavy setup/tests.
  - Validation: workflow YAML parses; `cd apps/flutter && flutter test test/snapshot_test.dart`; full quality gate; GitHub Actions run and same-SHA cache-hit rerun.

## P47

- [x] Trust-scope GitHub Actions unit cache saves.
  - Acceptance: SSOT, Go, Rust, and Flutter unit caches can be restored by push and pull-request runs, but new known-good unit cache markers are saved only by push events in the canonical `kimsemi-home/myhome-jarvis` repository; pull requests still run cache-miss validation without publishing cache markers.
  - Validation: workflow YAML parses; full quality gate; GitHub Actions run.

## P48

- [x] Pin Rust toolchain for hash-scoped CI.
  - Acceptance: Rust tests and CI use checked-in `rust-toolchain.toml` with an exact Rust toolchain; the Rust unit cache key includes that file so toolchain changes cannot reuse an old known-good Rust marker.
  - Validation: `rustup toolchain install 1.96.0 --profile minimal --component rustfmt --component clippy`; `cargo test --workspace`; `cargo fmt --check`; `cargo clippy --workspace -- -D warnings`; workflow YAML parses; full quality gate; GitHub Actions run.

## P49

- [x] Add toolchain pin drift check to quality gate.
  - Acceptance: `mhj quality` fails when `.go-version`, `go.mod`, generated project Go version, workflow `GO_VERSION`, `rust-toolchain.toml`, or workflow `RUST_TOOLCHAIN` drift from each other; the default quality output remains redacted to step names and statuses.
  - Validation: `go test ./cmd/mhj`; full quality gate; public safety scans; GitHub Actions run.

## P50

- [x] Run toolchain pin verification in split CI.
  - Acceptance: `mhj toolchain verify` exposes the toolchain pin check as a lightweight CLI; the Go GitHub Actions unit runs it on cache misses; the Go unit cache key includes `.go-version` and `rust-toolchain.toml` so pin-only changes cannot reuse an old Go unit marker.
  - Validation: `go test ./cmd/mhj`; `go run ./cmd/mhj toolchain verify`; workflow YAML parses; full quality gate; GitHub Actions run and same-SHA cache-hit rerun.

## P51

- [x] Guard split CI workflow cache contract.
  - Acceptance: `mhj ci verify` fails when the quality workflow loses public-safety history checks, Go toolchain verification wiring, generated cache inputs, canonical-repo cache save scoping, or generated command catalog coverage; `mhj quality` includes the redacted `ci workflow` step.
  - Validation: `go test ./cmd/mhj`; `go run ./cmd/mhj ci verify`; workflow YAML parses; full quality gate; GitHub Actions run and same-SHA cache-hit rerun.

## P52

- [x] Guard public-repo CI permission boundary.
  - Acceptance: `mhj ci verify` fails when the quality workflow loses top-level read-only contents permission, adds `pull_request_target`, or grants write permissions such as `contents: write` or `write-all`.
  - Validation: `go test ./cmd/mhj`; `go run ./cmd/mhj ci verify`; workflow YAML parses; full quality gate; GitHub Actions run and same-SHA cache-hit rerun.

## P53

- [x] Reject generic CI write permissions.
  - Linear: KIM-5
  - Acceptance: `mhj ci verify` fails on any workflow permission line ending in `write`, such as `id-token: write`, while keeping the public workflow on top-level `contents: read`.
  - Validation: `go test ./cmd/mhj`; `go run ./cmd/mhj ci verify`; workflow YAML parses; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P54

- [x] Scope Linear pull to active team issues.
  - Linear: KIM-6
  - Acceptance: `mhj linear pull` and `mhj linear next` filter out completed/canceled issues, optionally scope to private `LINEAR_TEAM_KEY` or `LINEAR_TEAM_ID`, and keep default summaries free of raw team names, workspace URLs, descriptions, UUIDs, tokens, and absolute paths.
  - Validation: `go test ./internal/linear`; `go run ./cmd/mhj linear next`; SSOT validation and codegen verification; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P55

- [x] Prefer project Linear issues in next selection.
  - Linear: KIM-7
  - Acceptance: SSOT owns the `[myhome-jarvis]` Linear issue title prefix; `mhj linear next` prefers active project-prefixed issues over unrelated active team issues; local backlog seeds use the same prefix; default summaries remain redacted.
  - Validation: `go test ./internal/linear`; `LINEAR_TEAM_KEY=KIM go run ./cmd/mhj linear next`; SSOT validation and codegen verification; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P56

- [x] Require project Linear issue for next selection.
  - Linear: KIM-8
  - Acceptance: `mhj linear next` selects only active `[myhome-jarvis]` issues; when active team issues exist but none are project-prefixed, it returns a redacted synced result without a selected issue; `mhj linear pull` still returns active redacted summaries.
  - Validation: `go test ./internal/linear`; `LINEAR_TEAM_KEY=KIM go run ./cmd/mhj linear next` before and after completing KIM-8; SSOT validation and codegen verification; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P57

- [x] Record planner status in checkpoint evidence.
  - Linear: KIM-9
  - Acceptance: `mhj loop once` and `mhj loop worker --cycles 1` checkpoint JSON includes redacted planner status with counts, repo-relative checkpoint root, quality/offline-fallback flags, and gated task metadata; adjacent checkpoint writes use collision-resistant filenames; checkpoints still omit raw Linear identities, security findings, local roots, absolute paths, tokens, and command output.
  - Validation: `go test ./internal/orchestrator ./internal/scheduler`; `go run ./cmd/mhj loop once`; `go run ./cmd/mhj loop worker --cycles 1`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P58

- [x] Make Linear backlog seeding project-aware and idempotent.
  - Linear: KIM-10
  - Acceptance: backlog seeds represent current project follow-up work; `mhj linear create-from-backlog` queries existing Linear issue titles, creates only missing `[myhome-jarvis]` seeds, and returns a synced zero-created summary when every seed already exists; default summaries still omit raw Linear URLs, workspace identities, UUIDs, tokens, absolute paths, local checkout paths, and raw descriptions.
  - Validation: `go test ./internal/linear`; `sbcl --script lisp/scripts/validate-ssot.lisp`; `go run ./cmd/mhj codegen verify`; `go run ./cmd/mhj linear create-from-backlog`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P59

- [x] Add DDD SSOT and local KnowledgeIndex thin slice.
  - Linear: KIM-14
  - Acceptance: Common Lisp SSOT defines bounded contexts, concept registry, aliases, generated artifact contracts, planning rules, and KnowledgeIndex schema; codegen emits `generated/concepts.generated.json`; `mhj ddd verify` and `mhj knowledge search` work locally; planner status and checkpoints include redacted KnowledgeIndex evidence; docs/logs stay public-safe.
  - Validation: `go test ./internal/knowledge ./internal/planner ./internal/orchestrator ./internal/linear`; `sbcl --script lisp/scripts/validate-ssot.lisp`; `go run ./cmd/mhj codegen verify`; `go run ./cmd/mhj ddd verify`; `go run ./cmd/mhj knowledge search KnowledgeIndex`; `go run ./cmd/mhj loop once`; security/history checks; full quality gate; GitHub Actions run and same-SHA cache-hit rerun.

## P60

- [x] Strengthen DDD SSOT with events and harness contracts.
  - Linear: KIM-15
  - Acceptance: concepts declare valid SSOT-owned `ddd_kind` values and every approved DDD kind is represented; generated concepts include domain events and harness case contracts; `mhj ddd verify` checks DDD kinds, domain events, harness contracts, duplicate concepts, alias drift, generated targets, and KnowledgeIndex policy; `mhj knowledge search DomainEvent` returns event evidence without raw private content.
  - Validation: `go test ./internal/knowledge ./internal/planner ./internal/orchestrator`; `sbcl --script lisp/scripts/validate-ssot.lisp`; `go run ./cmd/mhj codegen verify`; `go run ./cmd/mhj ddd verify`; `go run ./cmd/mhj knowledge search DomainEvent`; `go run ./cmd/mhj loop once`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.
