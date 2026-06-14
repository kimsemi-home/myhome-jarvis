# Harness

Harnesses are deterministic and must not use Python. They run against local
fixtures and dry-run command plans only.

The home harness validates:

- YouTube open and search dry-runs.
- OTT known and unknown service behavior.
- Volume set, up, and down boundaries.
- Display sleep dry-run argv.
- Safe and unsafe URL handling.
- Movie and sleep mode dry-runs.

The finance harness validates:

- Fixture transaction record count and KRW currency.
- Credit, debit, and net totals.
- Subscription review candidate count and total.
- User and household owner summaries.

The commerce harness validates:

- Fixture purchase record count and KRW currency.
- Total purchase spend.
- Recurring purchase candidate count and candidate details.
- User and household owner summaries.

Fixtures live under `fixtures`.

Validation commands:

```sh
go run ./cmd/mhj harness home
go run ./cmd/mhj harness finance
go run ./cmd/mhj harness commerce
go test ./internal/commands ./internal/daemon
```
