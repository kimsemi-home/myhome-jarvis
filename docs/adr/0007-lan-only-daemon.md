# ADR 0007: LAN-Only Daemon

## Status

Accepted

## Decision

The daemon binds to `127.0.0.1` by default. LAN access requires explicit config
and a local token for non-localhost clients.
