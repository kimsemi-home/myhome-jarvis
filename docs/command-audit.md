# Command Audit

Home-control command dispatch writes a private, append-only intent journal.

## Location

```text
data/private/audit/command-intents.jsonl
```

The file is ignored by Git and stores one JSON object per command intent.

## Recorded Fields

The journal records only:

- timestamp
- source (`cli` or `daemon`)
- normalized command name
- dry-run state
- whether execution was requested
- whether execution was allowed
- invocation count
- warning count
- success flag
- coarse error category

It does not record payload JSON, argv arrays, URLs, headers, bearer tokens,
environment variables, local absolute paths, command output, or raw error text.

## Status Surfaces

- `mhj audit status`
- daemon `GET /audit/status`
- Flutter Status `Command Audit` count

These status surfaces return the repo-relative journal path, whether it exists,
the total event count, and the last redacted event.

## Validation

```sh
go test ./internal/audit ./internal/daemon
go run ./cmd/mhj audit status
```
