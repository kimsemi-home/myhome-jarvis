# Supervision

The first process supervision surface records daemon runtime state under
ignored private storage and exposes read-only status checks.

## Commands

- `mhj daemon`
- `mhj daemon status`

## State

When the daemon successfully binds its TCP listener, it writes:

```text
data/private/supervisor/daemon-state.json
```

The state file is private and ignored by Git. It records only:

- daemon name
- process id
- bind host, port, and address
- version
- execute and LAN-bind flags
- started and updated timestamps

It does not record request bodies, response bodies, tokens, environment
variables, local absolute paths, or Linear data.

## Status

`mhj daemon status` and daemon `GET /supervisor/status` return a read-only
snapshot with:

- whether a state file exists
- the repo-relative state path
- recorded process metadata
- whether the recorded pid is still running
- whether `/health` is reachable
- whether the state appears stale

`mhj assistant status` includes a compact `supervisor` summary for the command
center: repo-relative state path, recorded/stale booleans, process/probe
booleans, and message. Detailed pid, address, and probe URL fields stay on
`mhj daemon status`.

The command center also derives `local_runtime` health from the supervisor
summary. When the daemon state is missing, stale, not running, or not reachable,
the assistant raises a public-safe `local_runtime` gate and chooses
`repair_local_runtime_health` as the next safe action. The universal work item
then points at `local_runtime:supervisor` evidence while keeping approval,
external writes, and self-approval disabled.

The health probe does not attach bearer tokens. LAN daemon supervision can
still use the recorded process metadata even if an unauthenticated health probe
is rejected.

## Validation

```sh
go test ./internal/supervisor ./internal/daemon
go run ./cmd/mhj daemon status
```
