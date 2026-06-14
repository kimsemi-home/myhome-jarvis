# Quality Evidence

`mhj quality` writes a private, redacted run summary after every quality gate
execution.

## Location

```text
data/private/quality/runs.jsonl
```

The file is ignored by Git and stores one JSON object per quality run.

## Recorded Fields

The journal records only:

- timestamp
- overall pass/fail
- duration in milliseconds
- step count
- pass/fail/skip counts
- step names and statuses

It does not record command argv, command output, local absolute paths,
environment variables, tokens, private data, generated artifact contents, or
raw test output.

## Status Surfaces

- `mhj quality status`
- daemon `GET /quality/status`
- Flutter Status `Quality` card

These surfaces return the repo-relative journal path, whether it exists, the
total run count, and the last redacted run summary.

## Validation

```sh
go test ./internal/qualitylog ./internal/daemon
go run ./cmd/mhj quality status
```
