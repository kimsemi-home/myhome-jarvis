# Closed Loop

Each autonomous cycle should:

1. Determine Linear status or offline fallback.
2. Inspect repository state.
3. Pick one small task.
4. Record a working-log start entry.
5. Modify one file or a tightly connected set of files.
6. Run the relevant quality gate.
7. Record results and checkpoint evidence.

The initial `loop once` command records a local checkpoint and never loops
forever. Checkpoint filenames include sub-second precision so adjacent loop
cycles do not overwrite each other.

The bounded worker surface is `mhj loop worker --cycles <n>`. Each cycle writes
private scheduler state with heartbeat, next-run, backoff, and checkpoint
metadata. `mhj loop status` and daemon `GET /loop/status` recover that state
without claiming external sync success.

Closed-loop checkpoints store redacted status summaries. Linear evidence keeps
mode, token-configured state, sync state, repo-relative queue path, HTTP status,
rate-limit remaining count, viewer-configured boolean, and team count only.
Planner evidence stores SSOT-backed counts, quality/offline-fallback flags,
repo-relative checkpoint root, and gated task metadata only.
Public-safety evidence stores aggregate current-tree and Git-history booleans,
finding counts, and checked timestamp only. Checkpoints do not store raw Linear
viewer/team identities, raw security findings, matched content, local
repository roots, or absolute private paths.

Before committing or pushing, closed-loop work can inspect repository state with
`mhj repo status` or daemon `GET /repo/status`. The response uses
repository-relative paths only, including ignored private data paths.

Quality evidence can be inspected with `mhj quality status` or daemon
`GET /quality/status`. The private journal records step names and statuses only,
not command output, argv arrays, raw test output, tokens, or local absolute
paths.
