# Codex Cost Governor

The Codex Cost Governor turns assistant and automation usage into local
evidence before the project scales paid or external loops. The SSOT source is
`lisp/ssot/codex-cost.lisp`; the public generated artifact is
`generated/codex_cost.generated.json`.

Usage records belong in the private append-only ledger:

```text
data/private/codex-cost/usage.jsonl
```

Use `mhj codex-cost record <json-payload>` to append a usage sample locally.
The command accepts loop scope, unit kind, amount, optional status, and
repo-relative evidence refs; it fills the recorded timestamp when omitted and
stores a semantic hash for cache/de-duplication evidence. Public command output
only exposes scope, unit kind, amount, status, evidence ref count, budget state,
and timestamp.

Use `mhj codex-cost guard <json-payload>` before long-running or expensive
assistant loops. The guard reads current cost and sustainability status, then
returns `allow`, `warn`, or `review_required` with public-safe reason codes.
It accepts planned scope, unit kind, estimated units, estimated minutes, and
repo-relative evidence refs. It does not persist raw loop prompts or evidence
contents.

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

## Public Boundary

The CLI commands `mhj codex-cost status`, `mhj codex-cost record`,
`mhj codex-cost guard`, and daemon endpoint `GET /codex-cost/status` must not
expose prompts, transcripts, private notes, raw evidence refs, credentials,
tokens, local absolute paths, account IDs, card numbers, Linear private URLs,
or private evidence contents.
