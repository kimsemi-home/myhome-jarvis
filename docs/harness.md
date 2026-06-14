# Harness

Harnesses are deterministic and must not use Python. They run against local
fixtures and dry-run command plans only.

The dedicated Rust boundary is `crates/mhj-harness`. It validates the
home-control dry-run command matrix through `mhj-command` and validates finance
and commerce fixture invariants through `mhj-finance` and `mhj-commerce`.
The Go CLI and daemon still expose the user-facing harness commands, so local
automation keeps the same JSON report shape while Rust owns the deterministic
fixture harness checks.

The home harness validates:

- YouTube open and search dry-runs.
- OTT known and unknown service behavior.
- Service-specific OTT shortcut targets.
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
- Recurring fixture rows and grouped candidate details.
- Top merchant spend.
- User and household owner summaries.

Fixtures live under `fixtures`.

Validation commands:

```sh
go run ./cmd/mhj harness home
go run ./cmd/mhj harness finance
go run ./cmd/mhj harness commerce
cargo test -p mhj-harness
go test ./internal/commands ./internal/daemon
```
