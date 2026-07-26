# Low-cost deployment profile

The cost-conscious deployment shape is one small Go daemon instance plus the
Flutter client. Keep the daemon on the home network or a private network by
default; do not expose the finance endpoint directly to the public internet.

1. Run the generated local loop before packaging:

   ```sh
   make -f generated/local_quality.generated.mk verify
   ```

2. Start the daemon with `MYHOME_BIND_HOST=127.0.0.1` for local use. If a
   remote device is needed, put it behind an authenticated private-network
   path and keep `MYHOME_EXECUTE=false`.

3. Keep `MYHOME_FINANCE_MODE=fixture` until consent and provider approval are
   complete. For a later live deployment, record active read-only household
   consent first (`go run ./cmd/mhj finance-consent status`), then inject only
   server-side `op://...` references through the runtime secret store. The
   adapter discovers consented bank accounts and cards by default; set
   `MYHOME_MYDATA_INCLUDE_CARDS=false` only when card approval data is out of
   scope. Do not commit a resolved token or account number.

4. Back up only the private consent and review ledgers through an encrypted
   channel. Public repository artifacts contain fixture summaries and the
   golden image, never raw household data.

The deployment gate is therefore reproducible on a laptop first and portable
to a small VM/container later, without making hosting a prerequisite for UI
or data-contract development.

The live path is intentionally a private daemon path, not a public finance
API. It requires active consent, a provider-issued access token in 1Password,
and the provider's approved MyData base URL. Without those runtime inputs the
fixture path remains the safe, deterministic fallback.

## GitHub Actions publication

After a successful `quality` run on `main`, the public repository's
`publish` workflow packages the verified commit, Go daemon binary, Flutter
client, fixtures, generated contracts, and docs as a short-retention Actions
artifact. It uses the free public-repository runner path and does not start a
paid server or expose household data. A later release/tag policy can promote
the same verified archive without changing the finance contract.
