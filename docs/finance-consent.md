# Household Finance Consent

`HouseholdFinanceConsent` is the private consent boundary for real finance
connectors and shared household scopes.

The current finance assistant remains fixture-first, read-only, and
review-only. It can summarize local fixtures and recommend review items, but it
must not transfer money, make payments, trade assets, apply for cards, cancel
cards, or mutate external finance systems.

## Private Ledger

Consent records live only in:

`data/private/finance/consent.jsonl`

Each record must include:

- consent kind
- subject scope
- consent status
- review status
- authority profile
- evidence refs

Private notes, subject details, account ids, card numbers, and raw finance
payloads stay out of public generated files.

## Public Status

Use `mhj finance-consent status` to see only readiness and counts. Use
`mhj finance-consent record <json-payload>` to append approved read-only
consent decisions to the private ledger. The public status and record result do
not expose evidence refs or private consent payload.

Readiness stays `blocked` until active approved consent exists for real finance
connectors, spouse scope, and household scope. Even then the mode remains
`read_only_review_only`.

The record command rejects unknown fields, private notes, absolute or
path-traversal evidence refs, invalid enums, non-RFC3339 timestamps, and
subject scopes that are not public-safe tokens. It never enables transfers,
payments, trades, card actions, or external finance writes.

Finance consent records are part of the private storage archive lane, so
`mhj storage-archive run` compresses and manifests the ledger with the same
noise-budget evidence used for other local operational logs.
