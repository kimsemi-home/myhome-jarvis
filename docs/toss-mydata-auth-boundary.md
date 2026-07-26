# Toss MyData authentication boundary

Toss authentication and financial data retrieval are two related operations,
but this application binds them to one connection credential bundle. The
bundle is the atomic unit of configuration and rotation: a bank/card request
cannot select a client id, consent token, data token, or certificate from a
different connection.

## Authentication plane

Toss MyData integrated authentication uses the `mydata.cert.toss.im` host and
the `POST /v1/ca/sign_verification` operation. It requires a provider-issued
access token, transaction id, certification transaction id, and consent value.
The approved provider flow returns a user CI and establishes the consent
context. This step is user- and provider-specific, so it is not executed by
the public GitHub Actions workflow or the Flutter client.

See the [Toss MyData integrated authentication guide](https://toss.im/tosscert/docs/guides/integration/mydata)
and [Toss common authentication rules](https://toss.im/tosscert/docs/guides/common).

## Data plane implemented here

After provider consent is established, the Go adapter calls the approved
MyData data endpoint configured by `MYHOME_MYDATA_BASE_URL`:

- `GET /v2/bank/accounts` and `POST /v2/bank/accounts/deposit/transactions`
- `GET /v2/card/cards`
- `GET /v2/card/cards/{card_id}/approval-domestic`
- `GET /v2/card/cards/{card_id}/approval-overseas`

Each connection stores only one `op://...` reference to a
`myhome.finance.credential/v1` JSON bundle in configuration. The bundle keeps
`provider`, `client_id`, `client_secret`, `auth_access_token`,
`data_access_token`, and optional mTLS certificate/private key material together. The server
resolves the bundle through 1Password; the Flutter client and public Actions
jobs never receive it. The adapter filters
accounts/cards by active consent, maps the bank and card response fields into
the canonical transaction IR, and rejects unsupported transaction codes.

The endpoint shapes and required access-token headers are based on the
[KFTC MyData testbed catalog](https://developers.mydatakorea.org/mdtb/apg/dgi/bas/FSAG0102),
[bank specification](https://developers.mydatakorea.org/mdtb/apg/mac/bas/FSAG0404?id=1),
and [card specification](https://developers.mydatakorea.org/mdtb/apg/mac/bas/FSAG0406?id=2).

## Final live verification inputs

The repository can verify the complete adapter and daemon loop with replay
responses. To verify a real provider call, supply outside the public repo:

1. provider-approved test base URL and one credential bundle for each spouse;
2. the corresponding `org_code`, account number, and optional card ids;
3. active approved read-only consent records;
4. an allowed network route to the provider test environment.

Then run `go run ./cmd/mhj finance parity` locally. A passing report at or above
95% is the live evidence; fixture/replay parity must not be presented as a
production provider call.
