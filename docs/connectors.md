# Connector Readiness

The next product phase starts with a public-safe connector readiness catalog,
not real account integrations.

The catalog is owned by Common Lisp SSOT in `lisp/ssot/connectors.lisp` and
emitted as `generated/connectors.generated.json`. Go exposes it through
`mhj connectors status` and daemon `GET /connectors/status`; Flutter renders
the same data as read-only connector cards.

Current connector entries are fixture-only plans for:

- MyData aggregation
- bank accounts
- card spending
- securities accounts
- commerce purchases
- payment wallets

Allowed operations are limited to local fixture reads, summaries, and review
recommendations. The catalog explicitly blocks credential requests, external
API calls, cookie import, scraping, transfers, trades, purchases, payment
execution, card applications, and card cancellations.

Public-safety rules:

- Do not store tokens, cookies, account identifiers, card numbers, personal
  contact fields, local absolute paths, raw finance data, raw commerce data, or
  external API responses in generated artifacts.
- Do not add connector action buttons to Flutter until a separate explicit
  execution boundary exists.
- Keep real connector work behind new Linear issues with acceptance criteria
  for consent, local vault storage, redacted status, fixture replay, and
  public-safety checks.

Validation:

```sh
go run ./cmd/mhj connectors status
go test ./internal/connectors ./internal/daemon
go run ./cmd/mhj codegen verify
go run ./cmd/mhj quality
```
