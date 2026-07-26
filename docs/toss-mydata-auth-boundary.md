# Toss MyData authentication boundary

Toss authentication and financial data retrieval are two related operations,
but this application binds them to one connection credential bundle. The
bundle is the atomic unit of configuration and rotation: a bank/card request
or authentication request cannot select a client id, consent token, data token,
or certificate from a different connection.

## Authentication plane

Toss MyData integrated authentication uses the `mydata.cert.toss.im` host and
the `POST /v1/ca/sign_verification` operation. The server-side
`LiveClient.SignVerification` method sends the provider-issued access token,
certification transaction id, transaction id, PKCS#7 signed consent, consent
type, and consent digest. A successful response returns a user CI and
establishes the consent context. The CI remains inside the private server-side
flow; it is not returned by daemon status or the Flutter client.

The OAuth client-credentials exchange that produces the Toss access token is
also server-side. Store the resulting access token in the same bundle as the
data access token; do not put it in a separate environment variable. The
provider may require mTLS in the test environment, so the optional certificate
and private key in that same bundle are installed into the per-connection HTTP
client. Partial or invalid mTLS material is rejected.

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
jobs never receive it. `SignVerification` and the data adapter use the same
resolved bundle. The adapter filters
accounts/cards by active consent, maps the bank and card response fields into
the canonical transaction IR, and rejects unsupported transaction codes.

The endpoint shapes and required access-token headers are based on the
[KFTC MyData testbed catalog](https://developers.mydatakorea.org/mdtb/apg/dgi/bas/FSAG0102),
[bank specification](https://developers.mydatakorea.org/mdtb/apg/mac/bas/FSAG0404?id=1),
and [card specification](https://developers.mydatakorea.org/mdtb/apg/mac/bas/FSAG0406?id=2).

## Final live verification inputs

The repository can verify the complete adapter and daemon loop with replay
responses. To verify a real provider call, supply outside the public repo:

1. provider-approved test data and auth base URLs plus one credential bundle for
   each spouse;
2. the corresponding `org_code`, account number, and optional card ids;
3. active approved read-only consent records and provider-issued auth proof;
4. an allowed network route to the provider test environment, including mTLS
   when required.

Then run `go run ./cmd/mhj finance parity` locally. A passing report at or above
95% is the live evidence; fixture/replay parity must not be presented as a
production provider call. The local authentication command reads one private
signed-consent JSON document from stdin, resolves the selected owner's bundle,
and returns only `{ "provider": "toss_mydata", "owner": "...", "verified": true }`:

```sh
go run ./cmd/mhj finance auth user < data/private/finance/user-signed-consent.json
go run ./cmd/mhj finance auth spouse < data/private/finance/spouse-signed-consent.json
go run ./cmd/mhj finance parity
```

The signed-consent documents are never committed, printed, or sent to the
Flutter client.
