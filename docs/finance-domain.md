# Finance Domain

Initial finance entities:

- Household
- Person
- Account
- Card
- Transaction
- Merchant
- Subscription
- Benefit
- Recommendation

The first phase uses mock fixtures and IR only. Real bank, card, securities, or
MyData credentials are not requested.

Rust foundation:

- `mhj-finance::TransactionIr` is the dedicated finance crate boundary for
  fixture-only validation and summaries.
- `mhj-core::finance::TransactionIr` models the SSOT transaction fields.
- Fixture JSONL lives at `fixtures/finance_transactions.jsonl`.
- Validation covers owner, ISO-like timestamps, positive money amounts,
  currency codes, account/card presence, debit merchant presence, raw refs, and
  non-empty tags.
- `mhj-finance::summarize_cashflow` and `mhj-core::finance::summarize_cashflow`
  provide deterministic fixture-level cashflow calculation during the domain
  split.
- `mhj-finance::summarize_by_owner` keeps user, spouse, and household cashflow
  views separate.
- `mhj-finance::subscription_candidates` identifies subscription-like debit
  records as review-only candidates; it does not execute subscription changes.

Go daemon read surface:

- `GET /domain/summary` reports fixture record count, currency, credit, debit,
  net minor units, and categories.

Validation:

```sh
cargo test -p mhj-finance
```
