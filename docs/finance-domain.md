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

- `mhj-core::finance::TransactionIr` models the SSOT transaction fields.
- Fixture JSONL lives at `fixtures/finance_transactions.jsonl`.
- Validation covers owner, ISO-like timestamps, positive money amounts,
  currency codes, account/card presence, debit merchant presence, raw refs, and
  non-empty tags.
- `mhj-core::finance::summarize_cashflow` provides the first deterministic
  fixture-level cashflow calculation.

Go daemon read surface:

- `GET /domain/summary` reports fixture record count, currency, credit, debit,
  net minor units, and categories.
