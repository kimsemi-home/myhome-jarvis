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
The first execution boundary is intentionally narrow: CLI execution requires
`MYHOME_EXECUTE=true`, while daemon execution also requires `--execute` and
request `execute=true`. The executor only permits `open`, `osascript`, and
`pmset` argv plans.

The first Go daemon surface exposes `GET /health`, `GET /version`,
`GET /commands`, `POST /intent`, `POST /harness/run`, `GET /linear/status`,
`POST /linear/sync`, `GET /repo/status`, `GET /loop/status`, `GET /domain/summary`,
`GET /household/summary`, `GET /recommendations/summary`, `GET /metrics`,
`GET /events`, `GET /supervisor/status`, `GET /audit/status`, and
`GET /quality/status`, and `GET /planner/status`.
It binds to `127.0.0.1` by default.
LAN binding requires `--allow-lan` and non-localhost requests must include a
Bearer token stored only in `data/private/local-token.txt`.

The first daemon observability surface is local and bounded. It keeps the
newest 100 request events in memory and records only method, path, status,
duration, timestamp, and coarse error category. It does not record bodies,
headers, bearer tokens, query strings, or local filesystem paths.

The first process supervision surface records daemon runtime state only after
the TCP listener binds successfully. The private state file lives at
`data/private/supervisor/daemon-state.json`, while `mhj daemon status` and
`GET /supervisor/status` expose repo-relative status, pid liveness, and a
token-free health probe.

The first command audit surface records redacted command intent metadata under
`data/private/audit/command-intents.jsonl`. `mhj audit status` and
`GET /audit/status` expose only the repo-relative journal path, count, and last
redacted event. Payloads, argv arrays, URLs, headers, bearer tokens, raw errors,
and local absolute paths are never recorded.

The first quality evidence surface records redacted quality gate summaries
under `data/private/quality/runs.jsonl`. `mhj quality status` and
`GET /quality/status` expose only the repo-relative journal path, count, and
last run summary. Command output, argv, environment variables, raw test output,
and local absolute paths are never recorded.

The first repository safety surface is read-only. `mhj repo status` and daemon
`GET /repo/status` report branch, head SHA, tracked changes, untracked files,
and ignored private data paths using repository-relative paths only.

The first scheduler surface is bounded and local-only. `mhj loop worker
--cycles N` records heartbeat/checkpoint state under `data/private` and uses
rate-limit/backoff metadata instead of spinning forever. `GET /loop/status`
exposes the current closed-loop policy and recovered private state.

The first planner surface is SSOT-backed. Common Lisp owns the task graph,
Linear issue templates, quality requirement, and external-write boundary;
`generated/planner.generated.json`, `mhj planner status`, and
`GET /planner/status` expose only repository-relative planning metadata.

The first Rust domain surface lives in `mhj-core`. It validates finance
transaction fixtures, commerce purchase fixtures, and recommendation scoring
before any real external finance or commerce connectors are introduced. The
dedicated `mhj-finance` crate owns a fixture-only finance IR boundary for
cashflow, owner summaries, and subscription review candidates. The dedicated
`mhj-commerce` crate owns a fixture-only purchase IR boundary for spend
summaries, merchant summaries, and recurring purchase review candidates.
`mhj-core` keeps the existing integrated fixture pipeline while domains are
split out. The dedicated `mhj-storage` crate owns data lake manifests,
repository-relative storage paths, raw JSONL fixture writes, and fixture-only
curated Parquet+Zstd writes for finance and commerce datasets.

The Go daemon exposes the first domain read surface at `GET /domain/summary`.
It reads local fixture JSONL and generated storage policy only; it does not
connect to bank, commerce, or lake services.
`mhj harness finance`, `mhj harness commerce`, and daemon `POST /harness/run`
validate the same read-only fixture summaries without introducing external
credentials or actions.

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
offline fallback. The Status tab also surfaces whether the repository is clean
or dirty, whether the recorded daemon supervisor state is reachable, and how
many command audit and quality gate events are recorded.
Platform runner files are left out until device packaging is required.
