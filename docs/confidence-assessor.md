# Confidence Assessor

The Confidence Assessor is the first executable confidence cap for the Agent
Cluster.

It does not accept an agent's self-reported certainty. It reads redacted local
status from the Evidence Graph, Learning Ledger, quality evidence, and public
safety checks, then returns the maximum confidence level currently allowed.

## SSOT

Common Lisp owns the policy in `lisp/ssot/confidence.lisp` and emits
`generated/confidence.generated.json`.

The generated policy requires:

- confidence is a cap, not an average score
- self-reported confidence is forbidden
- public status is redacted
- raw evidence is not exposed
- evidence graph, learning ledger, quality gate, and public safety inputs
- deterministic cap rules

## Commands

```sh
go run ./cmd/mhj confidence status
```

The command exposes:

- policy path
- assessor key
- confidence cap
- blocked boolean
- self-report allowance
- evidence link count
- dangling evidence ref count
- open learning count
- quality recorded/ok booleans
- public safety boolean
- active rule
- checked timestamp

It does not expose summaries, next actions, raw evidence refs, raw prompts,
raw transcripts, tokens, credentials, local absolute paths, account IDs, card
numbers, or private evidence contents.

## Cap Rules

- Public-safety findings block confidence.
- Latest failing quality gate blocks confidence.
- Missing evidence links cap confidence at low.
- Dangling evidence refs cap confidence at low.
- Open learning debt caps confidence at medium.
- Missing quality evidence caps confidence at medium.
- Clear evidence links, quality, public safety, and learning debt allow high.

## Daemon Status

`GET /confidence/status` is read-only and returns the same redacted summary as
the CLI. Flutter renders the result as a Status metric.

## Validation

Use these checks after changing the assessor:

```sh
go test ./internal/confidence ./internal/daemon ./cmd/mhj ./internal/knowledge
go run ./cmd/mhj confidence status
go run ./cmd/mhj codegen verify
go run ./cmd/mhj ddd verify
cd apps/flutter && flutter test test/daemon_client_test.dart test/widget_test.dart
```
