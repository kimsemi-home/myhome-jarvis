# Local finance evidence

The local finance evidence manifest proves four indirect capabilities and six
execution rehearsals without account credentials or external writes:

- ledger: deterministic monthly credit summary from a public fixture;
- portfolio: read-only holdings parsing from an official-response-shaped fixture;
- revenue: monthly YouTube revenue and local cost reconciliation from fixtures;
- shorts: a private-upload plan whose production runtime boundary remains
  plan-only, paired with an exact-loopback OAuth, channel-binding, and resumable
  private-upload rehearsal. Its sealed connection-readiness contract covers a
  desktop OAuth client, random callback port, twenty Keychain-backed slots,
  private state roots, and explicit activation while account binding stays off.
- ledger credit collection: the production OAuth token and Gmail attachment
  paths exercised against exact IPv4-loopback emulators, followed by private
  inbox import and monthly SQLite reconciliation. The same proof embeds two
  immutable versions of a synthetic issuer template, cross-version source-ID
  deduplication, and append-only category suggestion history.
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
- shorts activation: a random-port one-shot callback receiver plus readiness-
  bound system-browser and Keychain adapters exercised through fake runners.
  Success, wrong state, user denial, and browser-launch failure are rehearsed;
  runtime entrypoints, real browser/Keychain execution, credentials, external
  network, and writes remain disabled.

Run the bundled proof with:

```sh
go run ./cmd/mhj local-finance evidence
```

Each producer hashes its deterministic artifact and seals a receipt. The
Ledger, Portfolio, Revenue, Finance Operator, Shorts upload, and Shorts
activation rehearsals are copied as
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
purchases, KRW 2,200 in refunds, and KRW 18,700 net card spend. Its nested
template proof binds two source hashes and two distinct profile hashes to one
stable fingerprint set. The second template adds three new classification
suggestions without rewriting the three accepted transactions. Jarvis also
requires immutable template-version registration, missing-source-ID denial,
stable-identity conflict rollback, zero partial writes, and exact nested and
outer report hashes.

The nested onboarding proof adds two independently hash-validated read-only
previews. Jarvis requires each preview to select exactly one versioned profile,
bind its source/profile/fingerprint hashes and purchase/refund totals to the
subsequent import, omit raw rows, and become import-ready only after expected
totals and the credit-liability opening/closing balance equation reconcile.

Jarvis also validates the batch preview independently: two sorted source-name
hashes, two unique source and fingerprint sets, no raw filenames or rows, exact
manifest and preview-set hashes, and a recomputed batch hash. Ambiguous
catalogs, unsupported statements, mismatched totals or balances, duplicate
content or identities, path traversal, symlink root escape, and source-byte
mutations are required attack cases. Unsafe balance or batch claims still fail
after the preview, batch, nested template, and outer hashes are recomputed.

The separate Ledger batch-apply proof remains fixture-only and writes only to
ephemeral SQLite databases. Jarvis independently recomputes the approval
challenge, plan, approval, per-statement apply-set, first-apply, replay, and
outer report hashes. It requires the exact manifest, preview-set, and batch
hashes; five inserted rows on the first apply; five duplicates on replay;
pre-write stale/denied approval rejection; and zero accounts, transactions,
import runs, templates, or category suggestions after an injected later-file
failure. Rehashing the full hierarchy for another batch or claiming a partial
rollback still fails validation.

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
Jarvis also requires the Shorts indirect receipt to contain both the random-port
callback and sealed connection-readiness checks; recomputing a receipt after
removing either check still fails the manifest contract.

The separate Shorts activation proof requires exact IPv4 callback host, path,
random-port, one-shot, duplicate-denial, cancellation, and query-bound checks.
Authorization must be created after binding and handed only to the official
Google origin through an exact fake `/usr/bin/open -u` command. The short-lived
permit binds the readiness plan, authorization URL, state, and redirect; the
four-path rehearsal proves wrong-state and denial produce no token-exchange
plan, successful PKCE remains plan-only, and launch failure yields a manual URL
only in memory. No URL, state, code, or PKCE material enters JSON evidence.
Google device flow remains deferred because its documented YouTube scope list
does not include `youtube.upload`.

The proof also requires a fake Keychain runner, three exact command shapes, no
secret value in argv, default/expired/unlisted/network permit denial, readiness-
plan binding, caller-side material zeroing, and no committed runtime activation
entrypoint. Jarvis rejects unsafe browser or Keychain booleans even if an
attacker recomputes both the report and manifest hashes.

The checked-in manifest contains public synthetic data only. It is deployment
evidence for the indirect implementation, not proof of a live bank, broker, or
Google account connection. No OAuth token, mailbox, brokerage account, or
financial action was accessed or executed.
