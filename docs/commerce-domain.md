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

- `mhj-core::commerce::PurchaseIr` models purchase fixture rows.
- Fixture JSONL lives at `fixtures/commerce_purchases.jsonl`.
- Validation covers owner, ISO-like timestamps, merchant and item identity,
  quantity, non-negative money amounts, currency consistency, total-price math,
  raw refs, and non-empty tags.
- `mhj-core::commerce::recurring_candidates` provides the first deterministic
  repeated-purchase recommendation skeleton.

Go daemon read surface:

- `GET /domain/summary` reports fixture purchase count, repeated-purchase
  candidate count, candidates, and categories.
