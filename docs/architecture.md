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
`pmset` argv plans. OTT service shortcuts for Netflix, Disney+, TVING, Wavve,
and Coupang Play are zero-payload dry-run commands over the same safe URL map
as `open_ott`.

The first Go daemon surface exposes `GET /health`, `GET /version`,
`GET /auth/status`, `GET /commands`, `POST /intent`, `POST /harness/run`,
`GET /linear/status`, `POST /linear/sync`, `GET /repo/status`,
`GET /security/status`, `GET /loop/status`, `GET /domain/summary`,
`GET /connectors/status`, `GET /agent-cluster/status`,
`GET /learning/status`, `GET /evidence/status`,
`GET /confidence/status`, `GET /household/summary`,
`GET /recommendations/summary`, `GET /metrics`,
`GET /events`, `GET /supervisor/status`, `GET /audit/status`,
`GET /quality/status`, and `GET /planner/status`.
It binds to `127.0.0.1` by default.
LAN binding requires `--allow-lan` and non-localhost requests must include a
Bearer token stored only in `data/private/local-token.txt`.
The daemon HTTP server applies bounded read-header, read, write, idle, and
header-size limits by default so slow or idle clients cannot hold local server
resources indefinitely.
`GET /auth/status` reports configured/missing state and repo-relative token
path metadata only; it never returns the token value.
`GET /linear/status` reports a redacted Linear summary only: configured/synced
state, repo-relative queue path, HTTP metadata, viewer-configured boolean, and
team count without raw identities or token source.

The first daemon observability surface is local and bounded. It keeps the
newest 100 request events in memory and records only method, path, status,
duration, timestamp, and coarse error category. It does not record bodies,
headers, bearer tokens, query strings, or local filesystem paths. The
`GET /metrics` endpoint also exposes aggregate Go runtime counters for
goroutines, heap, stack, and GC without exposing local roots or raw process
data.

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
and ignored private data paths using repository-relative paths only. Public
release safety also includes `mhj security history`, which scans reachable Git
commits and commit metadata for private identity markers, local absolute paths,
forbidden language artifacts, private/lake data paths except empty keep
placeholders, and secret-looking literals without returning raw matched secret
contents. Daemon `GET /security/status` exposes only aggregate current-tree and
history booleans, finding counts, and a checked timestamp for local UI status.

The first scheduler surface is bounded and local-only. `mhj loop worker
--cycles N` records heartbeat/checkpoint state under `data/private` and uses
rate-limit/backoff metadata instead of spinning forever. `GET /loop/status`
exposes the current closed-loop policy and recovered private state.

The first planner surface is SSOT-backed. Common Lisp owns the task graph,
Linear issue templates, quality requirement, and external-write boundary;
`generated/planner.generated.json`, `mhj planner status`, and
`GET /planner/status` expose only repository-relative planning metadata. The
planner status reports ready, completed, and external-write-gated counts, and
omits `next_task` when the local rails are complete. It still lists
external-write-gated task metadata so the remaining blocked step is visible
without contacting or mutating Linear.

The first Rust domain surface lives in `mhj-core`. It validates finance
transaction fixtures, commerce purchase fixtures, and recommendation scoring
before any real external finance or commerce connectors are introduced. The
dedicated `mhj-harness` crate owns deterministic home, finance, and commerce
fixture harness checks over the Rust command, finance, and commerce crate
boundaries without executing external actions.
The dedicated `mhj-finance` crate owns a fixture-only finance IR boundary for
cashflow, owner summaries, and subscription review candidates. The dedicated
`mhj-commerce` crate owns a fixture-only purchase IR boundary for spend
summaries, merchant summaries, and recurring purchase review candidates.
`mhj-core` keeps the existing integrated fixture pipeline while domains are
split out. The dedicated `mhj-storage` crate owns data lake manifests,
repository-relative storage paths, raw JSONL fixture writes, and fixture-only
curated Parquet+Zstd writes and metadata reads for finance and commerce
datasets.

The Go daemon exposes the first domain read surface at `GET /domain/summary`.
It reads local fixture JSONL and generated storage policy only; it does not
connect to bank, commerce, or lake services.
`mhj harness finance`, `mhj harness commerce`, and daemon `POST /harness/run`
validate the same read-only fixture summaries without introducing external
credentials or actions.

The first recommendation surface is fixture-only. Rust owns the scoring
skeleton, Go projects read-only daemon summaries, and Flutter shows structured
cash buffer, subscription, card-linked spend, and recurring purchase review
items in an Optimize tab with score, rationale, estimated amount, and evidence
count. Recommendations never execute purchases, subscription changes, card
actions, transfers, or investment trades.

The first finance dashboard is fixture-only. Flutter reads finance totals,
subscription spend, card-linked debit review totals, categories, and owner
breakdowns from daemon `/domain/summary`; it does not request credentials,
connect to banks or cards, or execute transfers, subscription changes, card
actions, or investment trades.

The first purchases dashboard is fixture-only. Flutter reads commerce spend,
recurring purchase candidates, categories, and owner spend breakdowns from
daemon `/domain/summary`; it does not request commerce credentials, scrape
stores or payment services, or automate purchases.

The first household surface is also fixture-only. Finance and commerce owner
fields are projected into User, Spouse, and Household scopes so the UI can
switch views without introducing real account credentials.

The first connector readiness surface is public-safe and fixture-only. Common
Lisp SSOT owns `generated/connectors.generated.json`, Go exposes
`mhj connectors status` and daemon `GET /connectors/status`, and Flutter renders
read-only connector cards. The catalog lists planned connector keys, categories,
fixture mode, data classes, allowed read-only operations, forbidden operations,
and next local preparation steps. It never stores or returns credentials,
cookies, account identifiers, card numbers, local paths, raw private data, or
external API responses.

The first Agent Cluster surface is also public-safe and read-only. Common Lisp
SSOT owns `generated/agent_cluster.generated.json`, Go exposes
`mhj agent-cluster status` and daemon `GET /agent-cluster/status`, and Flutter
renders the Cluster tab from daemon status signals. The policy records the
evidence-first learning loop, separated producer/reviewer/verifier/steward
roles, verification sidecars, incident lifecycle, debt classes, quarantine
triggers, and failure conditions. It does not execute external agents, store raw
transcripts, allow private data in public evidence, allow self-approval, or
accept self-reported final confidence.

The first Learning Ledger surface turns observation-based self-improvement into
private evidence. Common Lisp SSOT owns `generated/learning.generated.json`, Go
records validated observations with `mhj learning record`, and daemon
`GET /learning/status` exposes only repo-relative path, total/open/closed
counts, kind counts, lifecycle-stage counts, and last coarse classification.
Raw observation summaries, next actions, evidence refs, prompts, transcripts,
tokens, local absolute paths, and private journal contents never leave
`data/private/learning/observations.jsonl`.

The first Evidence Graph surface makes private evidence traceable without
turning it public. Common Lisp SSOT owns `generated/evidence.generated.json`,
Go exposes `mhj evidence status` and daemon `GET /evidence/status`, and Flutter
shows a redacted Evidence Graph metric. The graph connects learning
observations to evidence artifacts through counted edge kinds and reports
dangling refs as debt. Source keys, counts, node kinds, edge kinds, and
timestamps can leave the private boundary; raw evidence refs, summaries, next
actions, tokens, credentials, local absolute paths, and private evidence
contents cannot.

The first Confidence Assessor surface makes confidence an external cap instead
of an agent self-report. Common Lisp SSOT owns
`generated/confidence.generated.json`, Go exposes `mhj confidence status` and
daemon `GET /confidence/status`, and Flutter shows a redacted Confidence
metric. Public-safety findings or failing quality evidence block confidence;
missing evidence links and dangling refs cap confidence at low; open learning
debt or missing quality evidence cap confidence at medium. The assessor exposes
only counts, booleans, active rule, and the cap.

The first Translation Manifest surface makes context movement explicit. Common
Lisp SSOT owns `generated/translation.generated.json`, Go exposes
`mhj translation status` and daemon `GET /translation/status`, and Flutter
shows a redacted Translation metric. Private manifests and semantic loss
records stay under `data/private/translation`; public status exposes only
counts, context names, levels, booleans, and timestamps.

The first Flutter surface lives in `apps/flutter`. It is a Dart-only local
client with status, command, finance, purchases, Linear, storage, connector
readiness, Agent Cluster, household, and optimization tabs.
It can load snapshots from the localhost daemon while keeping a deterministic
offline fallback. The command tab includes explicit OTT shortcuts plus editable
payload commands for search, URL, and volume operations. The Finance tab shows
fixture cashflow totals and review-only subscription/card-linked spend signals.
The Purchases tab shows fixture commerce spend and recurring purchase review
signals without purchase automation.
The Optimize tab renders structured review-only recommendations from daemon
summaries, including purchase, subscription, card-linked spend, and cash-buffer
signals.
The Status tab also surfaces local-only or token-gated LAN network mode, LAN
auth configured/missing state without token contents, whether the repository is
clean or dirty, aggregate public-safety status without raw findings, whether
the recorded daemon supervisor state is reachable, and how many command audit
quality gate, open learning observations, and Evidence Graph links are
recorded. It also shows the external Confidence cap instead of any agent
self-reported certainty, plus the Translation manifest debt count for context
movement.
Platform runner files are left out until device packaging is required.
