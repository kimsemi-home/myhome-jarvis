# Recommendations

The first recommendation engine is fixture-only and read-only.

Scope:

- Score local finance and commerce fixture evidence.
- Rank household cash buffer, subscription review, card-linked spend review,
  and recurring purchase review items.
- Expose recommendations through the local daemon summary surfaces.
- Show recommendations in the Flutter Optimize tab.

Non-goals:

- No bank, card, brokerage, commerce, or payment credentials.
- No purchases, subscription changes, card signups, card cancellations, account
  closures, transfers, or investment trades.
- No scraping or cookie-based commerce integration.

Validation:

```sh
cargo test -p mhj-core recommendations
go test ./internal/domain ./internal/daemon
cd apps/flutter && flutter test && flutter analyze
```
