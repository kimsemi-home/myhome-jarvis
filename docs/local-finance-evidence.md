# Local finance evidence

The local finance evidence manifest proves that four indirect capabilities run
without account credentials or external writes:

- ledger: deterministic monthly credit summary from a public fixture;
- portfolio: read-only holdings parsing from an official-response-shaped fixture;
- revenue: monthly YouTube revenue and local cost reconciliation from fixtures;
- shorts: a private-upload plan whose runtime boundary remains plan-only.

Run the bundled proof with:

```sh
go run ./cmd/mhj local-finance evidence
```

Each producer hashes its deterministic artifact and seals a receipt. Jarvis
recomputes every receipt hash, requires the exact component/capability set, and
then verifies an aggregate hash bound to the manifest month. Unknown JSON fields,
extra JSON values, missing components, hash changes, and any enabled external
write fail closed.

The checked-in manifest contains public synthetic data only. It is deployment
evidence for the indirect implementation, not proof of a live bank, broker, or
Google account connection.
