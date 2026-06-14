# Commerce Domain

Initial commerce entities:

- ProductPurchase
- ProductItem
- Order
- Merchant
- RecurringPurchaseCandidate
- PriceTrend
- PurchaseRecommendation

The first phase uses fixtures and recommendation skeletons only. Purchase
automation and scraping that may violate terms are out of scope.

Rust foundation:

- `mhj-commerce::PurchaseIr` is the dedicated commerce crate boundary for
  fixture-only validation and summaries.
- `mhj-core::commerce::PurchaseIr` models purchase fixture rows.
- Fixture JSONL lives at `fixtures/commerce_purchases.jsonl`.
- Validation covers owner, ISO-like timestamps, merchant and item identity,
  quantity, non-negative money amounts, currency consistency, total-price math,
  raw refs, and non-empty tags.
- `mhj-commerce::summarize_commerce` reports deterministic fixture spend totals
  without executing purchases.
- `mhj-commerce::summarize_by_owner` keeps user, spouse, and household purchase
  views separate.
- `mhj-commerce::summarize_by_merchant` ranks merchants by fixture spend.
- `mhj-commerce::recurring_candidates` and
  `mhj-core::commerce::recurring_candidates` provide repeated-purchase review
  candidates only; they do not automate purchase actions.

Go daemon read surface:

- `GET /domain/summary` reports fixture purchase count, repeated-purchase
  candidate count, candidates, and categories.

Validation:

```sh
cargo test -p mhj-commerce
```
