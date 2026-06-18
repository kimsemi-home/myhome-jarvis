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

## P61

- [x] Include project queue status in loop checkpoints.
  - Linear: KIM-13
  - Acceptance: `mhj loop once` and `mhj loop worker --cycles 1` checkpoint evidence includes redacted `linear_next` project queue observation with selected project issue, issue identifiers, titles, update timestamps, and state types only; raw descriptions, workspace URLs, team identities, UUIDs, tokens, absolute paths, and local roots remain out of checkpoint and CLI output.
  - Validation: `go test ./internal/orchestrator ./cmd/mhj`; `go run ./cmd/mhj loop once`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P62

- [x] Track approved Linear write evidence.
  - Linear: KIM-11
  - Acceptance: successful approved Linear write mutations append private redacted evidence with action, public issue key when available, and `synced=true`; failed mutations, token misses, lookup failures, and queued offline actions do not create synced write evidence; public summaries avoid raw descriptions, workspace URLs, identities, UUIDs, tokens, absolute paths, and local checkout paths.
  - Validation: `go test ./internal/linear`; `go run ./cmd/mhj planner status`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

- [x] Reconcile planner external-write gate.
  - Linear: KIM-12
  - Acceptance: planner status exposes the standing SSOT-owned `external_write_gate` separately from redacted `linear_write_evidence`; `blocked_external_write` remains a boundary, not sync success, and synced mutation counts increase only from successful Linear API mutations recorded in private evidence.
  - Validation: `go test ./internal/planner ./internal/daemon`; `go run ./cmd/mhj planner status`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P63

- [x] Replay Linear offline queue with rate-aware backoff.
  - Linear: KIM-16
  - Acceptance: `mhj linear replay-offline` replays only in-scope write-safe queued comment and transition actions after Linear credentials are available; `LINEAR_TEAM_KEY` scopes replay to matching public issue keys; successful entries are tracked in private replay evidence to prevent duplicate replay; failed, unsupported, already queued, out-of-scope, and low-rate-limit-paused entries remain `synced=false` in the original queue; summaries expose counts, repo-relative private paths, coarse status, HTTP status, and rate-limit remaining only.
  - Validation: `go test ./internal/linear ./cmd/mhj`; SSOT validation and codegen verification; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P64

- [x] Add public-safe connector readiness catalog.
  - Linear: KIM-17
  - Acceptance: Common Lisp SSOT owns planned fixture-only connector categories for MyData, bank, card, securities, commerce, and payment surfaces; generated connector metadata contains only public-safe provider keys, categories, data classes, allowed read-only operations, forbidden operations, and next local preparation steps; `mhj connectors status`, daemon `GET /connectors/status`, and Flutter read-only connector cards expose no credentials, cookies, account identifiers, card numbers, local absolute paths, raw personal data, or external API responses.
  - Validation: `go test ./internal/connectors ./internal/daemon ./cmd/mhj`; `go run ./cmd/mhj connectors status`; `go run ./cmd/mhj codegen verify`; full quality gate; public safety scans; GitHub Actions run.

## P65

- [x] Add public-safe Agent Cluster learning-loop policy.
  - Linear: KIM-18
  - Acceptance: Common Lisp SSOT owns an evidence-first Agent Cluster policy with ontology/context rules, separated agent roles, verification sidecars, incident lifecycle, debt classes, quarantine triggers, failure conditions, and read-only status signals; codegen emits `generated/agent_cluster.generated.json`; `mhj agent-cluster status`, daemon `GET /agent-cluster/status`, and Flutter Cluster cards expose no external agent execution, raw transcript storage, private public evidence, self-approval, self-reported final confidence, credentials, tokens, local absolute paths, or raw private data.
  - Validation: `go test ./internal/agentcluster ./internal/daemon ./cmd/mhj ./internal/knowledge`; `go run ./cmd/mhj agent-cluster status`; `go run ./cmd/mhj codegen verify`; `go run ./cmd/mhj ddd verify`; Flutter focused tests; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P66

- [x] Add local Learning Ledger for loop gaps.
  - Linear: KIM-19
  - Acceptance: Common Lisp SSOT owns a private append-only learning ledger policy with allowed observation kinds, lifecycle stages, required evidence fields, allowed evidence refs, and redacted public summary fields; codegen emits `generated/learning.generated.json`; `mhj learning record` writes validated observations only to `data/private/learning/observations.jsonl`; `mhj learning status`, daemon `GET /learning/status`, and Flutter Status expose only repo-relative paths, counts, kinds, lifecycle stages, and timestamps without raw summaries, next actions, evidence contents, tokens, credentials, local absolute paths, prompts, transcripts, account IDs, or card numbers.
  - Validation: `go test ./internal/learning ./internal/daemon ./cmd/mhj ./internal/knowledge`; `go run ./cmd/mhj learning status`; `go run ./cmd/mhj codegen verify`; `go run ./cmd/mhj ddd verify`; Flutter focused tests; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P67

- [x] Add local Evidence Graph status.
  - Linear: KIM-20
  - Acceptance: Common Lisp SSOT owns a private Evidence Graph policy with allowed private sources, node kinds, edge kinds, evidence ref prefixes, and public redaction fields; codegen emits `generated/evidence.generated.json`; `mhj evidence status`, daemon `GET /evidence/status`, and Flutter Status expose only source keys, counts, node kinds, edge kinds, dangling-ref counts, and timestamps without raw observation summaries, next actions, evidence ref strings, tokens, credentials, local absolute paths, prompts, transcripts, account IDs, card numbers, or private evidence contents.
  - Validation: `go test ./internal/evidence ./internal/daemon ./cmd/mhj ./internal/knowledge`; `go run ./cmd/mhj evidence status`; `go run ./cmd/mhj codegen verify`; `go run ./cmd/mhj ddd verify`; Flutter focused tests; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P68

- [x] Add external Confidence Assessor status.
  - Linear: KIM-21
  - Acceptance: Common Lisp SSOT owns an external confidence-cap policy that forbids agent self-reporting, treats confidence as a cap, reads Evidence Graph, Learning Ledger, quality gate, and public-safety signals, and lowers or blocks the cap for missing evidence links, dangling evidence refs, open learning debt, missing/failing quality, or public-safety findings; codegen emits `generated/confidence.generated.json`; `mhj confidence status`, daemon `GET /confidence/status`, and Flutter Status expose only redacted counts, booleans, active rule, and confidence cap without raw evidence, summaries, next actions, evidence refs, tokens, credentials, local absolute paths, prompts, transcripts, account IDs, card numbers, or private evidence contents.
  - Validation: `go test ./internal/confidence ./internal/daemon ./cmd/mhj ./internal/knowledge`; `go run ./cmd/mhj confidence status`; `go run ./cmd/mhj codegen verify`; `go run ./cmd/mhj ddd verify`; Flutter focused tests; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P69

- [x] Add Translation Manifest and Loss Ledger status.
  - Linear: KIM-22
  - Acceptance: Common Lisp SSOT owns a Translation Manifest / Loss Ledger policy for context movement, required manifest fields, private manifest and loss-ledger paths, loss levels, forbidden loss categories, and redacted public summary fields; codegen emits `generated/translation.generated.json`; `mhj translation status`, daemon `GET /translation/status`, and Flutter Status expose only counts, context names, levels, booleans, and timestamps. Missing or malformed manifests are counted as open translation debt. Public surfaces must not expose raw semantic notes, raw mappings, known loss details, evidence refs, tokens, credentials, local absolute paths, prompts, transcripts, account IDs, card numbers, Linear private URLs, or private evidence contents.
  - Validation: `go test ./internal/translation ./internal/daemon ./cmd/mhj ./internal/knowledge`; `go run ./cmd/mhj translation status`; `go run ./cmd/mhj codegen verify`; `go run ./cmd/mhj ddd verify`; Flutter focused tests; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P70

- [x] Add Control Plane Manifest status.
  - Linear: KIM-23
  - Acceptance: Common Lisp SSOT owns a private Control Plane Manifest policy for local orchestration decision receipts, required manifest fields, allowed decision kinds, authority profiles, lease statuses, lease bounds, reviewer/verifier separation, and redacted public summary fields; codegen emits `generated/control_plane.generated.json`; `mhj loop once` and bounded `mhj loop worker --cycles <n>` append private manifests after checkpoint writes; `mhj control-plane status`, daemon `GET /control-plane/status`, and Flutter Status expose only counts, debt totals, booleans, lease bounds, and timestamps. Public surfaces must not expose raw routing rationale, candidate agents, evidence refs, output refs, tokens, credentials, local absolute paths, prompts, transcripts, account IDs, card numbers, Linear private URLs, or private evidence contents.
  - Validation: `go test ./internal/controlplane ./internal/daemon ./cmd/mhj ./internal/knowledge ./internal/evidence`; `go run ./cmd/mhj control-plane status`; `go run ./cmd/mhj loop once`; `go run ./cmd/mhj codegen verify`; `go run ./cmd/mhj ddd verify`; Flutter focused tests; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P71

- [x] Add Incident Lifecycle status.
  - Linear: KIM-24
  - Acceptance: Common Lisp SSOT owns a private Incident Lifecycle policy for classified failures, lifecycle stages, owner roles, quarantine states, required fields, stale quarantine threshold, and redacted public summary fields; codegen emits `generated/incidents.generated.json`; `mhj incidents status`, daemon `GET /incidents/status`, and Flutter Status expose only counts, lifecycle buckets, owner-role buckets, quarantine-state buckets, incident debt, and timestamps. Missing ledger is allowed before the first incident; malformed records, missing owner roles, missing evidence refs, invalid lifecycle stages, and stale quarantine records count as incident debt. Public surfaces must not expose raw incident summaries, root-cause notes, evidence refs, prompts, transcripts, tokens, credentials, local absolute paths, account IDs, card numbers, Linear private URLs, or private evidence contents.
  - Validation: `go test ./internal/incidents ./internal/daemon ./cmd/mhj ./internal/knowledge ./internal/evidence`; `go run ./cmd/mhj incidents status`; `go run ./cmd/mhj codegen verify`; `go run ./cmd/mhj ddd verify`; Flutter focused tests; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P72

- [x] Add Evidence Quality Assessor status.
  - Linear: KIM-25
  - Acceptance: Common Lisp SSOT owns a private Evidence Quality Assessor policy for append-only quality snapshots, quality levels, mapping confidence levels, reassessment reasons, required fields, stale snapshot threshold, and redacted public summary fields; codegen emits `generated/evidence_quality.generated.json`; `mhj evidence-quality status`, daemon `GET /evidence-quality/status`, and Flutter Status expose only counts, quality buckets, mapping buckets, purpose buckets, reassessment-reason buckets, reassessment debt, thresholds, and timestamps. Missing ledger is allowed before the first snapshot; malformed snapshots, missing evidence refs, stale snapshots, low or blocked quality, and low or unknown mapping confidence count as reassessment debt. Public surfaces must not expose raw notes, raw evidence contents, evidence refs, prompts, transcripts, tokens, credentials, local absolute paths, account IDs, card numbers, Linear private URLs, or private evidence contents.
  - Validation: `go test ./internal/evidencequality ./internal/evidence ./internal/daemon ./cmd/mhj ./internal/knowledge`; `go run ./cmd/mhj evidence-quality status`; `go run ./cmd/mhj evidence status`; `go run ./cmd/mhj codegen verify`; `go run ./cmd/mhj ddd verify`; Flutter focused tests; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P73

- [x] Add Authority Gate status.
  - Linear: KIM-26
  - Acceptance: Common Lisp SSOT owns a public-safe Reasoning RBAC and Domain ABAC policy for reasoning tiers, role permissions, domain attributes, high-risk public-repo decision blocks, authority debt classes, and redacted public summary fields; codegen emits `generated/authority.generated.json`; `mhj authority status`, daemon `GET /authority/status`, and Flutter Status expose only outcome, active rule, input and decision counts, blocked/allowed decision keys, risk buckets, and authority debt counts. Reasoning tiers never grant approval, self-authority stays disabled, and public surfaces must not expose raw rationale, raw evidence contents, evidence refs, prompts, transcripts, tokens, credentials, local absolute paths, account IDs, card numbers, Linear private URLs, or private evidence contents.
  - Validation: `go test ./internal/authority ./internal/daemon ./cmd/mhj ./internal/knowledge`; `go run ./cmd/mhj authority status`; `go run ./cmd/mhj codegen verify`; `go run ./cmd/mhj ddd verify`; Flutter focused tests; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P74

- [x] Add Human Review Capacity status.
  - Linear: KIM-27
  - Acceptance: Common Lisp SSOT owns a private Human Review Capacity policy for review queue priority, reviewer roles, backup reviewer availability, overload rules, capacity thresholds, and redacted public summary fields; codegen emits `generated/review.generated.json`; `mhj review status`, daemon `GET /review/status`, Authority Gate, and Flutter Status expose only capacity state, debt counts, thresholds, buckets, and timestamps. Missing queue is allowed before the first review item; high-risk open reviews, too many open reviews, missing reviewers, missing evidence, and missing backup reviewer coverage become review debt. Public surfaces must not expose raw review notes, reviewer identities, evidence refs, prompts, transcripts, tokens, credentials, local absolute paths, account IDs, card numbers, Linear private URLs, or private evidence contents.
  - Validation: `go test ./internal/review ./internal/authority ./internal/evidence ./internal/daemon ./cmd/mhj ./internal/knowledge`; `go run ./cmd/mhj review status`; `go run ./cmd/mhj authority status`; `go run ./cmd/mhj codegen verify`; `go run ./cmd/mhj ddd verify`; Flutter focused tests; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P75

- [x] Add Code Shape Budget guard.
  - Linear: KIM-28
  - Acceptance: Common Lisp SSOT owns a 75-line source budget with generated legacy-debt baselines; `mhj code-shape status`, daemon `GET /code-shape/status`, GitHub Actions Go unit, and Flutter Status expose only redacted repo-relative budget status. Current oversized files are tracked as legacy debt, while new oversized files or growth beyond baseline fail the budget guard. Public surfaces must not expose source excerpts, local absolute paths, credentials, account IDs, Linear URLs, or private evidence.
  - Validation: `go test ./internal/codeshape ./internal/daemon ./cmd/mhj ./internal/knowledge`; `go run ./cmd/mhj code-shape status`; `go run ./cmd/mhj codegen verify`; `go run ./cmd/mhj ddd verify`; Flutter focused tests; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P76

- [x] Ratchet CLI code shape debt.
  - Linear: KIM-29
  - Acceptance: `cmd/mhj/main.go` sheds codegen, CI workflow contract, and quality orchestration helpers into small package-local files that each stay within 75 lines; generated Code Shape Budget baseline for `cmd/mhj/main.go` drops below the prior 1093-line debt value; `mhj quality`, `mhj ci verify`, and `mhj codegen verify` behavior remains unchanged.
  - Validation: `go test ./cmd/mhj ./internal/codeshape ./internal/knowledge`; `go run ./cmd/mhj ci verify`; `go run ./cmd/mhj codegen verify`; `go run ./cmd/mhj code-shape status`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P77

- [x] Ratchet CLI toolchain verify debt.
  - Linear: KIM-30
  - Acceptance: `cmd/mhj/main.go` sheds toolchain pin parsing and comparison helpers into small package-local files that each stay within 75 lines; generated Code Shape Budget baseline for `cmd/mhj/main.go` drops below the prior 893-line debt value; `mhj toolchain verify`, `mhj ci verify`, `mhj codegen verify`, and `mhj quality` behavior remains unchanged.
  - Validation: `go test ./cmd/mhj ./internal/codeshape ./internal/knowledge`; `go run ./cmd/mhj toolchain verify`; `go run ./cmd/mhj ci verify`; `go run ./cmd/mhj codegen verify`; `go run ./cmd/mhj code-shape status`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P78

- [x] Ratchet CLI loop orchestration debt.
  - Linear: KIM-31
  - Acceptance: `cmd/mhj/main.go` sheds bounded loop orchestration, checkpoint, scheduler worker, and control-plane manifest append helpers into small package-local files that each stay within 75 lines; generated Code Shape Budget baseline for `cmd/mhj/main.go` drops below the prior 791-line debt value; `mhj loop once`, `mhj loop status`, `mhj loop worker --cycles 1`, and `mhj quality` behavior remains unchanged.
  - Validation: `go test ./cmd/mhj ./internal/codeshape ./internal/orchestrator ./internal/scheduler`; `go run ./cmd/mhj loop once`; `go run ./cmd/mhj loop status`; `go run ./cmd/mhj loop worker --cycles 1`; `go run ./cmd/mhj code-shape status`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P79

- [x] Ratchet CLI status helper debt.
  - Linear: KIM-32
  - Acceptance: `cmd/mhj/main.go` sheds read-only status helpers into small package-local files that each stay within 75 lines; generated Code Shape Budget baseline for `cmd/mhj/main.go` drops below the prior 647-line debt value; status command routing, JSON output shape, and code-shape failure behavior remain unchanged.
  - Validation: `go test ./cmd/mhj ./internal/codeshape ./internal/agentcluster ./internal/authority`; representative status commands; `go run ./cmd/mhj code-shape status`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P80

- [x] Ratchet CLI quality helper debt.
  - Linear: KIM-33
  - Acceptance: `cmd/mhj/main.go` sheds quality report, command-runner, benchmark-smoke, gofmt, and Go file collection helpers into small package-local files that each stay within 75 lines; generated Code Shape Budget baseline for `cmd/mhj/main.go` drops below the prior 508-line debt value; `mhj quality`, `mhj benchmark smoke`, redacted quality JSON, and code-shape behavior remain unchanged.
  - Validation: `go test ./cmd/mhj ./internal/codeshape ./internal/qualitylog`; `go run ./cmd/mhj benchmark smoke`; `go run ./cmd/mhj code-shape status`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P81

- [x] Ratchet CLI command handler debt.
  - Linear: KIM-34
  - Acceptance: `cmd/mhj/main.go` sheds auth, daemon, harness, knowledge, learning, JSON, and env helpers into small package-local files that each stay within 75 lines; generated Code Shape Budget baseline for `cmd/mhj/main.go` drops below the prior 375-line debt value; CLI routing, daemon flags, harness failure behavior, knowledge/learning output, and JSON indentation remain unchanged.
  - Validation: `go test ./cmd/mhj ./internal/codeshape ./internal/auth ./internal/supervisor ./internal/knowledge ./internal/learning`; representative handler commands; `go run ./cmd/mhj code-shape status`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P82

- [x] Ratchet CLI integration handler debt.
  - Linear: KIM-35
  - Acceptance: `cmd/mhj/main.go` sheds security, command execution/audit, and Linear CLI routing into small package-local files that each stay within 75 lines; generated Code Shape Budget baseline for `cmd/mhj/main.go` drops below the prior 241-line debt value; `mhj security check/history`, `mhj command`, and Linear CLI behavior remain unchanged.
  - Validation: `go test ./cmd/mhj ./internal/codeshape ./internal/security ./internal/linear ./internal/commands`; representative security, command, and Linear commands; `go run ./cmd/mhj code-shape status`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P83

- [x] Ratchet CLI router debt under 75 lines.
  - Linear: KIM-36
  - Acceptance: `cmd/mhj/main.go` keeps only the process entrypoint and drops under the normal 75-line budget; remaining CLI dispatch moves into small basic/status/operation route files that each stay within 75 lines; version, commands, status, quality, codegen, loop, daemon, Linear, security, command, and knowledge routing behavior remain unchanged.
  - Validation: `go test ./cmd/mhj ./internal/codeshape ./internal/commands ./internal/linear`; representative basic, status, operation, loop, codegen, and quality commands; `go run ./cmd/mhj code-shape status`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P84

- [x] Ratchet CLI test debt under 75 lines.
  - Linear: KIM-37
  - Acceptance: `cmd/mhj/main_test.go` is split into focused generated-diff, quality redaction, toolchain pin, CI workflow, fixture, and helper test files that each stay within 75 lines; no `cmd/mhj/*.go` file exceeds the normal 75-line budget; existing test assertions and behavior remain unchanged.
  - Validation: `go test ./cmd/mhj ./internal/codeshape`; `go run ./cmd/mhj code-shape status`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P85

- [x] Ratchet Linear status test debt under 75 lines.
  - Linear: KIM-38
  - Acceptance: `internal/linear/status_test.go` is split into focused client and summary test files that each stay within 75 lines; fake-token GraphQL endpoint checks and public-safe summary redaction assertions remain unchanged; `internal/linear/status_test.go` is removed from Code Shape legacy debt.
  - Validation: `go test ./internal/linear ./internal/codeshape`; `go run ./cmd/mhj code-shape status`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P86

- [x] Ratchet connector status test debt under 75 lines.
  - Linear: KIM-39
  - Acceptance: `internal/connectors/status_test.go` keeps connector catalog and unsafe-operation assertions under 75 lines by moving shared repo-root fixture discovery into a focused helper file; connector public-safety assertions remain unchanged; `internal/connectors/status_test.go` is removed from Code Shape legacy debt.
  - Validation: `go test ./internal/connectors ./internal/codeshape`; `go run ./cmd/mhj code-shape status`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P87

- [x] Ratchet command executor test debt under 75 lines.
  - Linear: KIM-40
  - Acceptance: `internal/commands/executor_test.go` keeps allowed execution, dry-run skip, and platform skip assertions under 75 lines by moving unsafe executable rejection into a focused safety test file; command execution behavior remains unchanged; `internal/commands/executor_test.go` is removed from Code Shape legacy debt.
  - Validation: `go test ./internal/commands ./internal/codeshape`; `go run ./cmd/mhj code-shape status`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P88

- [x] Ratchet scheduler test debt under 75 lines.
  - Linear: KIM-41
  - Acceptance: `internal/scheduler/scheduler_test.go` keeps policy bounds and heartbeat/private state assertions under 75 lines by moving failure backoff and recovery coverage into a focused recovery test file; scheduler behavior remains unchanged; `internal/scheduler/scheduler_test.go` is removed from Code Shape legacy debt.
  - Validation: `go test ./internal/scheduler ./internal/codeshape`; `go run ./cmd/mhj code-shape status`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P89

- [x] Ratchet Agent Cluster status test debt under 75 lines.
  - Linear: KIM-42
  - Acceptance: `internal/agentcluster/status_test.go` keeps public-safe learning-loop policy assertions under 75 lines by moving self-approval rejection and repo-root fixture discovery into focused test/helper files; Agent Cluster status behavior remains unchanged; `internal/agentcluster/status_test.go` is removed from Code Shape legacy debt.
  - Validation: `go test ./internal/agentcluster ./internal/codeshape`; `go run ./cmd/mhj code-shape status`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P90

- [x] Ratchet auth local token debt under 75 lines.
  - Linear: KIM-43
  - Acceptance: `internal/auth/local.go` keeps only the shared local-token model and relative path helper under 75 lines by moving status, create/rotate, read, and token generation behavior into focused files; local LAN token behavior remains unchanged; `internal/auth/local.go` is removed from Code Shape legacy debt.
  - Validation: `go test ./internal/auth ./internal/codeshape`; `go run ./cmd/mhj code-shape status`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.

## P91

- [x] Ratchet domain summary test debt under 75 lines.
  - Linear: KIM-44
  - Acceptance: `internal/domain/summary_test.go` keeps the end-to-end fixture summary flow under 75 lines by moving finance, commerce, storage, recommendation, household, and repo-root assertions/helpers into focused test files; fixture summary behavior remains unchanged; `internal/domain/summary_test.go` is removed from Code Shape legacy debt.
  - Validation: `go test ./internal/domain ./internal/codeshape`; `go run ./cmd/mhj code-shape status`; full quality gate; public safety scans; GitHub Actions run and same-SHA cache-hit rerun.
