# Executable SSOT

The source of truth is Common Lisp code under `lisp/ssot`.

Generated JSON files under `generated` are artifacts, not source of truth.
Codegen must be deterministic: the same SSOT input should produce byte-for-byte
identical output.

Current SSOT boundaries are intentionally separated by domain:

- `project`: repository policy such as allowed languages and Go version.
- `ddd`: bounded contexts, canonical concepts, aliases, generated artifact
  contracts, planning rules, and KnowledgeIndex schema.
- `commands`: dry-run home command catalog.
- `finance`, `commerce`, `storage`: local fixture and lakehouse domain policy.
- `household`, `recommendations`, `scheduler`: local household views, optimization hints, and bounded loop policy.
- `security`, `connectors`, `agent-cluster`, `linear`, `planner`:
  public-safety rules, fixture-only connector readiness, evidence-first agent
  cluster policy, Linear workflow rules, and planning metadata.
- `learning`: private observation ledger policy for loop gaps and evidence
  debt.

The planner SSOT emits `generated/planner.generated.json`. Go reads that
artifact for `mhj planner status` and daemon `GET /planner/status`; Flutter
only consumes the daemon status. This keeps task graph shape, Linear templates,
quality requirements, and external-write boundaries in one Lisp-owned source.

The DDD SSOT emits `generated/concepts.generated.json`. Go verifies that
bounded contexts, DDD kinds, concept aliases, domain events, harness case
contracts, generated targets, and local KnowledgeIndex policy stay coherent
with `mhj ddd verify`.

The command SSOT emits `generated/commands.generated.json`. Go keeps the
runtime command registry and macOS execution planning in `internal/commands`,
but its tests load the generated artifact and fail if command names, summaries,
payload fields, OTT service allowlists, or generated URL targets drift from the
Lisp-owned catalog. Flutter static/offline fallback tests also read the same
artifact and fail if fallback command names or payload fields drift from the
catalog.

The security SSOT emits `generated/security.generated.json`. Go owns the
current-tree and Git-history scanners, while the generated policy records that
current non-private file contents are scanned for private identity markers and
secret-looking literals, private paths are skipped, and matched secret contents
must not be reported.

The connector SSOT emits `generated/connectors.generated.json`. Go reads that
artifact for `mhj connectors status` and daemon `GET /connectors/status`;
Flutter consumes the daemon status and keeps a static fixture-only fallback.
The artifact is limited to public-safe planned connector metadata and forbids
real credentials, external API calls, scraping, payments, transfers, trades,
purchases, and card actions in this phase.

The Agent Cluster SSOT emits `generated/agent_cluster.generated.json`. Go reads
that artifact for `mhj agent-cluster status` and daemon
`GET /agent-cluster/status`; Flutter consumes the daemon status and keeps a
static fallback aligned with generated signal keys. The artifact records
evidence-first ordering, role separation, sidecars, incident lifecycle, debt
classes, quarantine triggers, failure conditions, and public-safe status
signals. It forbids external agent execution, raw transcript storage, private
data in public evidence, self-approval, and self-reported final confidence in
this phase.

The Learning Ledger SSOT emits `generated/learning.generated.json`. Go reads
that artifact for `mhj learning status`, `mhj learning record`, and daemon
`GET /learning/status`; Flutter consumes the daemon status as a read-only
Learning metric. The generated policy keeps the journal under `data/private`,
requires evidence refs, owner, and next action for every observation, and keeps
public status redacted to counts, kinds, lifecycle stages, and timestamps.

Use `mhj codegen verify` before committing SSOT or generated artifact changes.
It snapshots the current `generated` tree, regenerates artifacts from Lisp, and
fails if regeneration changes any generated file. This verifies intended
working-tree SSOT/generated updates before they are committed.
