# ADR 0007: LAN-Only Daemon

## Status

Accepted

## Decision

The daemon binds to `127.0.0.1` by default. LAN access requires explicit config
and a local token for non-localhost clients.

Local token management is handled by `mhj auth token create` and
`mhj auth token rotate`. Token status may be inspected with `mhj auth status`,
which does not print the token value.
