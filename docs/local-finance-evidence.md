# Local finance evidence

The local finance evidence manifest proves four indirect capabilities and two
execution rehearsals without account credentials or external writes:

- ledger: deterministic monthly credit summary from a public fixture;
- portfolio: read-only holdings parsing from an official-response-shaped fixture;
- revenue: monthly YouTube revenue and local cost reconciliation from fixtures;
- shorts: a private-upload plan whose runtime boundary remains plan-only.
- ledger credit collection: the production Gmail attachment path exercised
  against an exact IPv4-loopback emulator, followed by private inbox import and
  monthly SQLite reconciliation.
- portfolio collection: the production KIS token and read-only balance path
  exercised against an exact IPv4-loopback emulator, followed by temporary
  SQLite persistence and aggregate-only Ledger publication.

Run the bundled proof with:

```sh
go run ./cmd/mhj local-finance evidence
```

Each producer hashes its deterministic artifact and seals a receipt. The Ledger
and Portfolio rehearsals are additionally copied as complete reports with both
file SHA-256 and internal report hash references. Jarvis recomputes every receipt and report hash,
requires the exact component/capability set, and then verifies an aggregate hash
bound to the manifest month. Unknown JSON fields, extra JSON values, missing
components, hash changes, and any enabled external write fail closed.

The rehearsal verifies one bounded retry after an injected 503, allowlisted
sender filtering, append-only attachment receipts, idempotent replay, archive
hash fallback after receipt loss, and a reconciled July result of KRW 20,900 in
purchases, KRW 2,200 in refunds, and KRW 18,700 net card spend.

The Portfolio rehearsal verifies one bounded retry after an injected 503, the
official read-only balance path and transaction ID, one idempotent SQLite
snapshot, one idempotent aggregate-only Ledger event, zero order requests, and
a reconciled KRW 50,000 cash + KRW 150,000 securities = KRW 200,000 total.

The checked-in manifest contains public synthetic data only. It is deployment
evidence for the indirect implementation, not proof of a live bank, broker, or
Google account connection. No OAuth token, mailbox, brokerage account, or
financial action was accessed or executed.
