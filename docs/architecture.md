# Architecture

`myhome-jarvis` is organized as a local-first system with strict language
boundaries.

- Go owns long-running processes, local API surfaces, Linear sync, quality
  gates, orchestration, security checks, and process supervision.
- Rust owns hot-path validation, deterministic harness logic, finance and
  commerce normalization, storage readers and writers, and benchmarks.
- Common Lisp owns executable SSOT, DSL-like definitions, deterministic
  codegen, and closed-loop planning policies.
- Flutter owns the eventual local client UI after the daemon and command core
  are stable.

The system defaults to dry-run. Any real macOS command execution must be
explicitly enabled and must use argv arrays, never shell interpolation.

The first Go daemon surface exposes `GET /health`, `GET /version`,
`GET /commands`, `POST /intent`, `POST /harness/run`, `GET /linear/status`,
`POST /linear/sync`, `GET /domain/summary`, `GET /household/summary`,
`GET /recommendations/summary`, and `GET /metrics`. It binds to `127.0.0.1`
by default.

The first Rust domain surface lives in `mhj-core`. It validates finance
transaction fixtures, commerce purchase fixtures, and Parquet+Zstd-ready lake
dataset plans before any real external finance, commerce, or storage connectors
are introduced.

The Go daemon exposes the first domain read surface at `GET /domain/summary`.
It reads local fixture JSONL and generated storage policy only; it does not
connect to bank, commerce, or lake services.

The first recommendation surface is fixture-only. Rust owns the scoring
skeleton, Go projects read-only daemon summaries, and Flutter shows the ranked
items in an Optimize tab. Recommendations never execute purchases, subscription
changes, card actions, transfers, or investment trades.

The first household surface is also fixture-only. Finance and commerce owner
fields are projected into User, Spouse, and Household scopes so the UI can
switch views without introducing real account credentials.

The first Flutter surface lives in `apps/flutter`. It is a Dart-only local
client with status, command, Linear, storage, household, and optimization tabs.
It can load snapshots from the localhost daemon while keeping a deterministic
offline fallback. Platform runner files are left out until device packaging is
required.
