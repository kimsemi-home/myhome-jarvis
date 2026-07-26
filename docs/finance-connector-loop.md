# Finance connector closed loop

The finance path is built as a closed loop before real account access is
enabled:

```text
Toss/MyData-shaped fixture
  -> normalized transaction IR
  -> required-field and 95% parity gate
  -> Go domain summary
  -> daemon /domain/summary
  -> Flutter household dashboard
  -> golden test
  -> README image link check
```

The committed fixture is `fixtures/finance_toss_mydata.jsonl`. It is a
provider-shaped test input, not an external response dump. The current local
adapter maps 14 canonical fields and reports 100% parity against the 95%
release threshold:

```sh
go run ./cmd/mhj finance parity
```

## 1Password boundary

The future live mode is server-side only:

```text
Flutter -> local Go daemon -> MyData/Toss adapter -> 1Password reference
```

Flutter never receives a token, cookie, account identifier, card number, or
raw provider response. The Go process accepts only an `op://...` reference
when `MYHOME_FINANCE_MODE=mydata`; the deployment environment is responsible
for resolving that reference. The default remains `fixture`, with external
calls disabled.

The real provider adapter is intentionally not enabled by this loop. It
requires provider approval, MyData consent/certification, private review of
both spouses' scopes, redacted logging, and a replay fixture captured from a
permitted test account.
