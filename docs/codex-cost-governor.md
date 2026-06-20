# Codex Cost Governor

The Codex Cost Governor turns assistant and automation usage into local
evidence before the project scales paid or external loops. The SSOT source is
`lisp/ssot/codex-cost.lisp`; the public generated artifact is
`generated/codex_cost.generated.json`.

Usage records belong in the private append-only ledger:

```text
data/private/codex-cost/usage.jsonl
```

Each record must include time, loop scope, unit kind, amount, status, and
repo-relative evidence refs. Public status surfaces only expose counts,
thresholds, buckets, total units, budget state, and timestamps.

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

## Public Boundary

The CLI command `mhj codex-cost status` and daemon endpoint
`GET /codex-cost/status` must not expose prompts, transcripts, private notes,
raw evidence refs, credentials, tokens, local absolute paths, account IDs, card
numbers, Linear private URLs, or private evidence contents.
