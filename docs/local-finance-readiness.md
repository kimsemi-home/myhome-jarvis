# Local finance connection readiness

`mhj local-finance readiness` validates the three collector connection plans
and the Finance Operator execution plan without activating any connector. The
default manifest and source plans live under `fixtures/local_finance_readiness/`.

The verifier fails closed unless all of the following remain true:

- Ledger, Portfolio, Revenue, and Finance Operator plans are `plan_only` and
  hash-sealed.
- Credentials, external networking, external writes, and service installation
  are all disabled.
- Each plan names only a public configuration source; collector OAuth scopes,
  when present, are read-only.
- Ledger and Revenue plans additionally prove exact callback, official OAuth
  host, redirect-disabled, response-bounded, and least-scope token boundaries;
  Portfolio proves the exact official KIS origin and token endpoint plus
  order-path, redirect, and oversized-response rejection.
- The copied plan artifact SHA-256 and its canonical internal `plan_hash` both
  match.
- Direct child schedules are disabled; Finance Operator is the sole execution
  owner through one uninstalled daily 08:10 KST launchd template.
- The operator catches up Ledger on day 2, Portfolio on day 3, Revenue and the
  unified snapshot on day 5, then Jarvis evidence verification follows at
  10:00 KST.

This proves configuration and ordering only. It does not inspect Keychain,
private connection overrides, databases, account identities, or live service
state, and it never runs `collect monthly` or installs launchd jobs.
