# Evidence Quality Assessor

Evidence Quality Assessor is the executable status surface for the Agent
Cluster rule that evidence quality changes over time. A useful artifact today
can become stale, lose mapping confidence after ontology/schema changes, or
become blocked by security, quarantine, translation, or counter-evidence.

It is not a public evidence browser. Raw notes, raw evidence contents, evidence
refs, prompts, transcripts, Linear URLs, local absolute paths, tokens,
credentials, account IDs, card numbers, and private evidence contents stay out
of public status surfaces.

## SSOT

Common Lisp owns the policy in `lisp/ssot/evidence-quality.lisp` and emits
`generated/evidence_quality.generated.json`.

The generated policy defines:

- the private append-only quality snapshot ledger
- quality levels
- mapping confidence levels
- allowed assessment purposes
- reassessment reasons
- stale snapshot threshold
- required fields
- public redaction fields

## Runtime

```sh
go run ./cmd/mhj evidence-quality status
```

The command reads `data/private/evidence-quality/snapshots.jsonl` when present
and returns a redacted quality summary only. A missing ledger is allowed before
the first snapshot. Malformed snapshots, missing evidence refs, stale
snapshots, low or blocked quality, and low or unknown mapping confidence count
as reassessment debt.

Daemon `GET /evidence-quality/status` returns the same redacted shape. Flutter
Status renders the snapshot count or reassessment debt count as a read-only
local metric.

## Public Status

Public status may expose:

- ledger presence
- snapshot count
- invalid snapshot count
- missing evidence-ref count
- reassessment debt count
- stale snapshot count
- low and blocked quality counts
- mapping drift count
- stale threshold
- counts by quality level, mapping confidence, purpose, and reassessment reason
- timestamps

It does not expose raw notes, raw evidence contents, evidence refs, prompts,
transcripts, tokens, credentials, local absolute paths, account IDs, card
numbers, private Linear URLs, or private evidence contents.

## Validation

Use these checks after changing the assessor policy:

```sh
go test ./internal/evidencequality ./internal/evidence ./internal/daemon ./cmd/mhj ./internal/knowledge
go run ./cmd/mhj evidence-quality status
go run ./cmd/mhj evidence status
go run ./cmd/mhj codegen verify
go run ./cmd/mhj ddd verify
cd apps/flutter && flutter test test/daemon_client_test.dart test/widget_test.dart
```
