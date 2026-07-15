# Local finance connection readiness

`mhj local-finance readiness` validates the three real-connection preparation
plans without activating any connector. The default manifest and source plans
live under `fixtures/local_finance_readiness/`.

The verifier fails closed unless all of the following remain true:

- Ledger, Portfolio, and Revenue plans are `plan_only` and hash-sealed.
- Credentials, external networking, external writes, and service installation
  are all disabled.
- Each plan names only a public configuration source and read-only OAuth scope.
- The copied plan artifact SHA-256 and its canonical internal `plan_hash` both
  match.
- The monthly DAG is Ledger on day 2 at 07:00, Portfolio on day 3 at 07:20,
  Revenue on day 5 at 07:40, then Jarvis evidence verification at 08:00 KST.

This proves configuration and ordering only. It does not inspect Keychain,
private connection overrides, databases, account identities, or live service
state, and it never runs `collect monthly` or installs launchd jobs.
