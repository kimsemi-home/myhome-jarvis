# Finance connector closed loop

The finance path is built as a closed loop before real account access is
enabled:

```text
Toss/MyData-shaped replay or permitted live response
  -> normalized transaction IR
  -> required-field and 95% parity gate
  -> Go domain summary
  -> daemon /domain/summary
  -> Flutter household dashboard
  -> golden test
  -> README image link check
```

The committed normalized fixture is `fixtures/finance_toss_mydata.jsonl` and
the raw-response-shaped replay is `fixtures/finance_toss_mydata_response.json`.
The adapter also supports the bank deposit transaction endpoint used by the
KFTC MyData testbed. The live bank contract intentionally excludes `card_id`
because that field is not supplied by this bank endpoint; all 13 supported
fields must be present on every record and the release threshold remains 95%.

Fixture mode maps 14 canonical fields and reports 100% parity:

```sh
go run ./cmd/mhj finance parity
```

## Live mode and 1Password boundary

Live mode is server-side only:

```text
Flutter -> local Go daemon -> MyData/Toss adapter -> 1Password reference
```

Flutter never receives a token, cookie, account identifier, card number, or
raw provider response. The Go process accepts only an `op://...` reference
when `MYHOME_FINANCE_MODE=mydata`; the deployment environment is responsible
for resolving that reference. The default remains `fixture`, with external
calls disabled.

Live mode is opt-in and requires provider approval, MyData consent/certification,
private review of both spouses' scopes, and a permitted test account. The
public GitHub Actions workflow never calls the provider and never receives a
1Password session. Local server configuration is shaped like this:

```sh
export MYHOME_FINANCE_MODE=mydata
export MYHOME_FINANCE_ALLOW_EXTERNAL=true
export MYHOME_MYDATA_BASE_URL=https://<provider-test-endpoint>
export MYHOME_FINANCE_CONNECTIONS_JSON='[{"owner":"user","onepassword_ref":"op://MyHome-Jarvis/Finance-Connector/jooyoon-token","org_code":"<org>","account_num":"<account>"},{"owner":"spouse","onepassword_ref":"op://MyHome-Jarvis/Finance-Connector/semmi-token","org_code":"<org>","account_num":"<account>"}]'
go run ./cmd/mhj finance parity
```

The provider-specific authentication/certification flow is outside this
public repository. See [Toss MyData integrated authentication](https://toss.im/tosscert/docs/guides/integration/mydata)
and the [KFTC MyData testbed bank transaction API](https://developers.mydatakorea.org/mdtb/apg/mac/bas/FSAG0404?id=1)
for the official provider-side prerequisites and request contract.
