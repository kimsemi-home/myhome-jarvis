# Learning Ledger

The Learning Ledger is the first executable slice of observation-based
self-improvement.

It records loop gaps, evidence debt, review debt, revalidation debt,
translation loss, quality regressions, and SSOT defect candidates as private
append-only observations. Public surfaces show only redacted counts and
classification metadata.

## SSOT

Common Lisp owns the policy in `lisp/ssot/learning.lisp` and emits
`generated/learning.generated.json`.

The policy requires every record to include:

- kind
- source
- summary
- evidence refs
- owner
- next action

The private journal path is `data/private/learning/observations.jsonl`.

## Commands

```sh
go run ./cmd/mhj learning status
go run ./cmd/mhj learning record '{"kind":"loop_gap","source":"quality_gate","summary":"Example gap.","evidence_refs":["data/private/quality/runs.jsonl"],"owner":"go","next_action":"Add a regression test."}'
```

`learning record` writes raw observation details only to the private JSONL
journal. Its CLI result exposes the generated observation ID, repo-relative
journal path, kind, lifecycle stage, status, evidence-ref count, and timestamp.

## Daemon Status

`GET /learning/status` is read-only. It exposes:

- repo-relative journal path
- generated policy path
- total/open/closed counts
- counts by kind
- counts by lifecycle stage
- last kind, stage, status, and observed timestamp

It does not expose summaries, next actions, raw evidence refs, tokens, local
absolute paths, raw prompts, raw transcripts, or private observation contents.

## Validation

Use these checks after changing the ledger:

```sh
go test ./internal/learning ./internal/daemon ./cmd/mhj ./internal/knowledge
go run ./cmd/mhj learning status
go run ./cmd/mhj codegen verify
go run ./cmd/mhj ddd verify
cd apps/flutter && flutter test test/daemon_client_test.dart test/widget_test.dart
```
