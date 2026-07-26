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
The adapter supports the bank deposit transaction endpoint and card approval
endpoints used by the KFTC MyData testbed. It discovers only consented bank
accounts/cards, then maps domestic and overseas card approvals into the same
transaction IR. A bank-only response intentionally excludes `card_id` because
that field is not supplied by the bank endpoint; when card approvals are
included, the full 14-field contract is evaluated.

Bank transaction direction follows the v2 transaction codes: withdrawal and
withdrawal correction/cancellation codes become debit, deposit and deposit
correction/cancellation codes become credit. The ambiguous `01` (new) code is
rejected instead of being silently counted as household income or spending.

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
raw provider response. Each connection accepts one `op://...` reference to a
versioned credential bundle when `MYHOME_FINANCE_MODE=mydata`; the server
resolves provider identity, auth material, data access, and optional mTLS from
that same bundle. The default remains `fixture`, with external calls disabled.

Live mode is opt-in and requires provider approval, MyData consent/certification,
private review of both spouses' scopes, and a permitted test account. The
public GitHub Actions workflow never calls the provider and never receives a
1Password session. Local server configuration is shaped like this:

```sh
export MYHOME_FINANCE_MODE=mydata
export MYHOME_FINANCE_ALLOW_EXTERNAL=true
export MYHOME_MYDATA_BASE_URL=https://<provider-test-endpoint>
export MYHOME_MYDATA_AUTH_BASE_URL=https://mydata.cert.toss.im
export MYHOME_MYDATA_API_TYPE=2
export MYHOME_MYDATA_INCLUDE_CARDS=true
export MYHOME_FINANCE_CONNECTIONS_JSON='[{"owner":"user","credential_ref":"op://MyHome-Jarvis/Finance-Connector/jooyoon-credential","org_code":"<org>","account_num":"<account>"},{"owner":"spouse","credential_ref":"op://MyHome-Jarvis/Finance-Connector/semmi-credential","org_code":"<org>","account_num":"<account>"}]'
go run ./cmd/mhj finance-consent status
# Feed the provider-issued signed-consent document from a private file/stdin.
go run ./cmd/mhj finance auth user < data/private/finance/user-signed-consent.json
go run ./cmd/mhj finance auth spouse < data/private/finance/spouse-signed-consent.json
go run ./cmd/mhj finance parity
```

Each referenced 1Password field contains one private JSON bundle. The access
token, client credentials, consent/auth token, and optional certificate
reference are rotated together:

```json
{"schema":"myhome.finance.credential/v1","provider":"toss_mydata","client_id":"test_client","client_secret":"<secret>","auth_access_token":"<auth-token>","data_access_token":"<data-token>","mtls_certificate_pem":"<certificate>","mtls_private_key_pem":"<private-key>"}
```

The provider-specific user ceremony still happens in the Toss app and requires
provider-issued test credentials. The server-side verification call is
implemented by `LiveClient.SignVerification`; it accepts the signed consent
proof and uses the same atomic bundle as the data fetch. The provider
authentication boundary and atomic credential bundle are documented in
[`docs/toss-mydata-auth-boundary.md`](toss-mydata-auth-boundary.md). See [Toss MyData integrated authentication](https://toss.im/tosscert/docs/guides/integration/mydata)
and the [KFTC MyData testbed API catalog](https://developers.mydatakorea.org/mdtb/apg/dgi/bas/FSAG0102),
[bank API specification](https://developers.mydatakorea.org/mdtb/apg/mac/bas/FSAG0404?id=1),
and [card API specification](https://developers.mydatakorea.org/mdtb/apg/mac/bas/FSAG0406?id=2)
for the official provider-side prerequisites and request contract. The
implemented replay fixtures cover bank transactions, card list discovery, and
domestic/overseas card approval shapes.
