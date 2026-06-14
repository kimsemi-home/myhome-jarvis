# LAN Auth

The daemon is localhost-only by default. LAN binding requires explicit opt-in:

```sh
go run ./cmd/mhj auth status
go run ./cmd/mhj auth token create
MYHOME_EXECUTE=false go run ./cmd/mhj daemon --host 192.168.1.10 --allow-lan
```

The daemon also exposes read-only token state for trusted local clients:

```sh
curl http://127.0.0.1:3888/auth/status
```

This endpoint returns configured/missing state and repo-relative token metadata
only. It never returns the token value.

For non-localhost clients, send:

```text
Authorization: Bearer <local token>
```

Token rules:

- The token file is `data/private/local-token.txt`.
- The file is ignored by Git and written with `0600` permissions.
- `auth status` never prints the token value.
- `auth token create` prints the token once for trusted household clients.
- `auth token rotate` replaces the token and invalidates older clients.

The Flutter daemon client accepts an optional auth token and sends it as a Bearer
header. The default local client does not include a token because localhost does
not require one. The Status tab renders only the configured/missing state from
`/auth/status`.
