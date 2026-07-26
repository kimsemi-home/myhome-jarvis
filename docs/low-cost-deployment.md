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
   complete. For a later live deployment, inject only the server-side
   `MYHOME_FINANCE_1PASSWORD_REF=op://...` reference through the runtime
   secret store. Do not commit the resolved value.

4. Back up only the private consent and review ledgers through an encrypted
   channel. Public repository artifacts contain fixture summaries and the
   golden image, never raw household data.

The deployment gate is therefore reproducible on a laptop first and portable
to a small VM/container later, without making hosting a prerequisite for UI
or data-contract development.

## GitHub Actions publication

After a successful `quality` run on `main`, the public repository's
`publish` workflow packages the verified commit, Go daemon binary, Flutter
client, fixtures, generated contracts, and docs as a short-retention Actions
artifact. It uses the free public-repository runner path and does not start a
paid server or expose household data. A later release/tag policy can promote
the same verified archive without changing the finance contract.
