# Local finance evidence

The local finance evidence manifest proves four indirect capabilities and five
execution rehearsals without account credentials or external writes:

- ledger: deterministic monthly credit summary from a public fixture;
- portfolio: read-only holdings parsing from an official-response-shaped fixture;
- revenue: monthly YouTube revenue and local cost reconciliation from fixtures;
- shorts: a private-upload plan whose production runtime boundary remains
  plan-only, paired with an exact-loopback OAuth, channel-binding, and resumable
  private-upload rehearsal.
- ledger credit collection: the production OAuth token and Gmail attachment
  paths exercised against exact IPv4-loopback emulators, followed by private
  inbox import and monthly SQLite reconciliation.
- portfolio collection: the production KIS token and read-only balance path
  exercised against an exact IPv4-loopback emulator, followed by temporary
  SQLite persistence and aggregate-only Ledger publication.
- revenue collection: the production Google OAuth token, YouTube channel
  lookup, and monetary Analytics day/video query contracts exercised against
  exact IPv4 loopback, followed by atomic SQLite replacement, cost replay, and
  aggregate-only Ledger replay.
- finance operator: the production subprocess, retry, checkpoint, next-day
  resume, completed-stage skip, aggregate-only snapshot, and replay paths
  exercised with three local child emulators.
- shorts upload: authorization-code/PKCE and refresh-token forms, authenticated
  channel lookup, canonical private metadata, bounded resumable recovery, and
  state-store replay exercised against one exact-loopback emulator.

Run the bundled proof with:

```sh
go run ./cmd/mhj local-finance evidence
```

Each producer hashes its deterministic artifact and seals a receipt. The
Ledger, Portfolio, Revenue, Finance Operator, and Shorts rehearsals are copied as
complete reports with both file SHA-256 and internal report hash references.
Jarvis recomputes every receipt and report hash, requires the exact
component/capability set, and then
verifies an aggregate hash bound to the manifest month. Unknown JSON fields,
extra JSON values, missing components, hash changes, and any enabled external
write fail closed.

The Ledger rehearsal verifies authorization-code plus PKCE and refresh-token
exchanges, official token-origin pinning, redirect rejection, a 1 MiB response
bound, one bounded Gmail retry after an injected 503, allowlisted sender
filtering, append-only attachment receipts, idempotent replay, archive hash
fallback after receipt loss, and a reconciled July result of KRW 20,900 in
purchases, KRW 2,200 in refunds, and KRW 18,700 net card spend.

The Portfolio rehearsal verifies the client-credentials token contract, exact
official KIS origin and token endpoint, order-path rejection, redirect
rejection, bounded token responses, one bounded retry after an injected 503,
the official read-only balance path and transaction ID, one idempotent SQLite
snapshot, one idempotent aggregate-only Ledger event, zero order requests, and
a reconciled KRW 50,000 cash + KRW 150,000 securities = KRW 200,000 total.

The Revenue rehearsal verifies PKCE authorization-code and refresh-token
contracts for the exact two read-only scopes, official token-origin pinning,
redirect rejection, a 1 MiB OAuth response bound, the bound-channel and
monetary read-only query contracts, one bounded retry after an injected 503,
two stable daily and video rows after full-month replacement, idempotent cost
import, and one idempotent aggregate-only Ledger event. Its synthetic estimated
result is KRW 8,300 gross minus KRW 2,000 cost, for KRW 6,300 net.

The Finance Operator rehearsal verifies the real child exit-code convention,
three-attempt bound, day-2/day-3/day-5 catch-up order, next-day failure resume,
completed-stage skip, one aggregate-only snapshot, and idempotent replay. The
snapshot reconciles KRW 18,700 net card spend, KRW 6,300 net tracked income,
KRW -12,400 tracked surplus, and KRW 200,000 total assets without persisting
raw child output.

The Shorts rehearsal verifies the exact youtube.upload scope, official OAuth
and channel origins, one authenticated channel-binding hash, redirect and
oversized-response denial, private canonical metadata with subscriber
notifications disabled, one injected resumable-upload interruption, and a new
runner replay that creates no additional session, probe, chunk, or video.

The checked-in manifest contains public synthetic data only. It is deployment
evidence for the indirect implementation, not proof of a live bank, broker, or
Google account connection. No OAuth token, mailbox, brokerage account, or
financial action was accessed or executed.
