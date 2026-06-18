# Evidence Graph

The Evidence Graph connects private observations to the evidence artifacts they
reference.

It is not a public evidence store. Raw observations, summaries, next actions,
and evidence ref strings stay in private journals. Public surfaces expose only
counts, node kinds, edge kinds, source keys, dangling-ref counts, and timestamps.

## SSOT

Common Lisp owns the policy in `lisp/ssot/evidence.lisp` and emits
`generated/evidence.generated.json`.

The policy declares:

- private evidence sources
- node kinds
- edge kinds
- allowed evidence ref prefixes
- public redaction fields
- the `mhj evidence status` command

## Commands

```sh
go run ./cmd/mhj evidence status
```

The command reads private evidence sources such as the Learning Ledger, quality
runs, checkpoints, Control Plane manifests, Incident Lifecycle records, Linear
write evidence, and command intent audit journal. It returns a redacted graph
summary only.

## Daemon Status

`GET /evidence/status` is read-only. It exposes:

- generated policy path
- private root
- total and present source counts
- total node and edge counts
- dangling evidence-ref count
- open learning observation count
- counts by node and edge kind
- source keys, formats, node kinds, presence, and counts
- last observed timestamp

It does not expose raw summaries, next actions, evidence ref strings, tokens,
credentials, local absolute paths, prompts, transcripts, account IDs, card
numbers, or private evidence contents.

## Validation

Use these checks after changing the graph:

```sh
go test ./internal/evidence ./internal/daemon ./cmd/mhj ./internal/knowledge
go run ./cmd/mhj evidence status
go run ./cmd/mhj codegen verify
go run ./cmd/mhj ddd verify
cd apps/flutter && flutter test test/daemon_client_test.dart test/widget_test.dart
```
