# Codex Cost Governor

The Codex Cost Governor turns assistant and automation usage into local
evidence before the project scales paid or external loops. The SSOT source is
`lisp/ssot/codex-cost.lisp`; the public generated artifact is
`generated/codex_cost.generated.json`.

Usage records belong in the private append-only ledger:

```text
data/private/codex-cost/usage.jsonl
```

Scope attribution records belong in a separate private append-only ledger:

```text
data/private/codex-cost/attribution.jsonl
```

Use `mhj codex-cost record <json-payload>` to append a usage sample locally.
The command accepts loop scope, unit kind, amount, optional status, and
repo-relative evidence refs; it fills the recorded timestamp when omitted and
stores a semantic hash for cache/de-duplication evidence. Public command output
only exposes scope, unit kind, amount, status, evidence ref count, budget state,
and timestamp.

Use `mhj codex-cost attribute <json-payload>` to attach already-recorded cost
to an interpretation scope such as a Linear project, repository, or
monetization experiment without increasing total budget usage. The private
record stores a safe subject key and evidence refs; public output returns only
the scope, amount, basis, subject hash, cost ref, evidence ref count, and
timestamp. When omitted, the cost ref is derived from unit kind, amount, and
evidence refs so the same cost can be viewed through multiple scope lenses
without inflating coverage.

Use `mhj codex-cost guard <json-payload>` before long-running or expensive
assistant loops. The guard reads current cost and sustainability status, then
returns `allow`, `warn`, or `review_required` with public-safe reason codes.
It accepts planned scope, unit kind, estimated units, estimated minutes, and
repo-relative evidence refs. It does not persist raw loop prompts or evidence
contents.

Use `mhj codex-cost roi` to review public-safe cost ROI by loop scope. The
summary keeps one row for every governed scope, including Linear projects,
repositories, and monetization experiments, even when a scope has no usage yet.
It combines private cost ledger totals, Codex sustainability posture, value
proxy units, attribution coverage, cache-savings evidence, accepted merge
evidence, and the storage archive/noise-budget configuration. Accepted changes
come from the Codex sustainability ledger and local first-parent GitHub-style
merge commits; ROI uses the stronger public-safe count and reports only counts,
source labels, and the log limit. The value proxy is explicitly allocated by
cost share until more precise per-scope monetization evidence exists.
ROI reports raw attribution entry units separately from deduplicated coverage
units. Coverage uses cost refs and should not exceed total recorded usage even
when a single cost is attributed to a Linear project, repository, and
monetization experiment at the same time.

Each stored record must include time, loop scope, unit kind, amount, status,
semantic hash, and repo-relative evidence refs. Public status surfaces only
expose counts, thresholds, buckets, total units, budget state, and timestamps.

## Agent Cluster Fit

Cost records are evidence, not raw transcript storage. The policy follows the
same loop as the broader Agent Cluster model:

```text
usage observation
evidence refs
semantic hash inputs
private ledger
redacted public status
review gate when thresholds are crossed
```

Budget state is a governance signal:

- `ok`: continue normal local-first loops.
- `warning`: inspect ROI and evidence before scaling.
- `review_required`: require explicit review before expanding paid or external
  automation.

The storage SSOT includes the private cost ledger in the `compress_then_archive`
source list. `mhj storage-archive run` can therefore compress cost usage JSONL
into private `.jsonl.gz` archive files while recording manifest rows and
enforcing the configured evidence noise budget.
The ROI summary reports the same archive pattern and noise-budget evidence
field so cost decisions can see whether local logs are being collected,
compressed, and governed below the configured noise threshold.
The attribution ledger is in the same private archive lane, so scope coverage
evidence is compacted and governed with the cost and sustainability ledgers.

## Public Boundary

The CLI commands `mhj codex-cost status`, `mhj codex-cost record`,
`mhj codex-cost guard`, `mhj codex-cost roi`, and daemon endpoint
`GET /codex-cost/status` must not expose prompts, transcripts, private notes,
raw evidence refs, credentials, tokens, local absolute paths, account IDs,
card numbers, Linear private URLs, or private evidence contents.
