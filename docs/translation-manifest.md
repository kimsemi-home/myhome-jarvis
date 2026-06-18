# Translation Manifest

The Translation Manifest is the first executable slice for context movement and
semantic loss tracking in the Agent Cluster vision.

It does not expose raw translation notes. It summarizes private manifests and
loss records so context moves can be audited without publishing private
evidence, prompts, credentials, account identifiers, or local paths.

## SSOT

Common Lisp owns the policy in `lisp/ssot/translation.lisp` and emits
`generated/translation.generated.json`.

The generated policy defines:

- allowed bounded contexts
- required manifest fields
- private loss ledger and private manifest root
- loss levels from `l0_none` through `l4_forbidden`
- forbidden loss categories for authority, security, consent, deletion,
  audit, legal, and financial meaning
- public redaction fields

## Runtime

```sh
go run ./cmd/mhj translation status
```

The command reads:

- `data/private/translation/losses.jsonl`
- `data/private/translation/manifests`

The public status includes only:

- ledger and manifest root presence
- manifest and loss counts
- missing or malformed manifest counts
- open debt and forbidden loss counts
- source and target context counts
- loss level counts
- timestamps

It does not expose summaries, semantic notes, raw mappings, known loss detail,
evidence refs, prompts, transcripts, tokens, credentials, local absolute paths,
account IDs, card numbers, or private evidence contents.

## Debt

Missing or malformed manifests are not invisible. The status counts them as
open translation debt. `l4_forbidden` or forbidden-category losses are counted
separately so automation can stay conservative when meaning that must never be
lost has been lost or cannot be proven preserved.
