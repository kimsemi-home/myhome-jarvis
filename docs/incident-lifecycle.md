# Incident Lifecycle

Incident Lifecycle is the executable follow-through for the Agent Cluster rule
that observed failures must be classified, assigned, verified, and turned into
knowledge instead of disappearing as untracked rough edges.

It is not a public incident report. Raw summaries, root-cause notes, private
evidence refs, prompts, transcripts, Linear URLs, local absolute paths, tokens,
credentials, account IDs, card numbers, and private evidence contents stay out
of public status surfaces.

## SSOT

Common Lisp owns the policy in `lisp/ssot/incidents.lisp` and emits
`generated/incidents.generated.json`.

The generated policy defines:

- the private append-only incident ledger
- allowed incident kinds
- lifecycle stages
- owner roles
- allowed status and quarantine states
- stale quarantine threshold
- required fields
- public redaction fields

## Runtime

```sh
go run ./cmd/mhj incidents status
```

The command reads `data/private/incidents/incidents.jsonl` when present and
returns a redacted lifecycle summary only. A missing ledger is allowed before
the first incident. Malformed records, missing owner roles, missing evidence
refs, invalid stages/statuses, and stale quarantine records count as incident
debt.

Daemon `GET /incidents/status` returns the same redacted shape. Flutter Status
renders the incident count or incident debt count as a read-only local metric.

## Public Status

Public status may expose:

- ledger presence
- valid incident count
- open and closed counts
- incident debt count
- missing owner count
- missing evidence-ref count
- stale quarantine count
- counts by kind, stage, status, owner role, and quarantine state
- timestamps

It does not expose raw summaries, root-cause notes, private evidence refs,
private incident contents, prompts, transcripts, tokens, credentials, local
absolute paths, account IDs, card numbers, or private Linear URLs.

## Validation

Use these checks after changing the lifecycle policy:

```sh
go test ./internal/incidents ./internal/daemon ./cmd/mhj ./internal/knowledge ./internal/evidence
go run ./cmd/mhj incidents status
go run ./cmd/mhj codegen verify
go run ./cmd/mhj ddd verify
cd apps/flutter && flutter test test/daemon_client_test.dart test/widget_test.dart
```
