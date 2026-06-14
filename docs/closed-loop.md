# Closed Loop

Each autonomous cycle should:

1. Determine Linear status or offline fallback.
2. Inspect repository state.
3. Query the Local KnowledgeIndex for relevant concepts and must-read files.
4. Pick one small task.
5. Record a working-log start entry.
6. Modify one file or a tightly connected set of files.
7. Run the relevant quality gate.
8. Record results and checkpoint evidence.

The initial `loop once` command records a local checkpoint and never loops
forever. Checkpoint filenames include sub-second precision so adjacent loop
cycles do not overwrite each other.

The bounded worker surface is `mhj loop worker --cycles <n>`. Each cycle writes
private scheduler state with heartbeat, next-run, backoff, and checkpoint
metadata. `mhj loop status` and daemon `GET /loop/status` recover that state
without claiming external sync success.

Closed-loop checkpoints store redacted status summaries. Linear status evidence
keeps mode, token-configured state, sync state, repo-relative queue path, HTTP
status, rate-limit remaining count, viewer-configured boolean, and team count
only. Linear next evidence stores the redacted next project issue summary,
issue identifiers, titles, update timestamps, and state types only.
Planner evidence stores SSOT-backed counts, quality/offline-fallback flags,
repo-relative checkpoint root, gated task metadata, the standing external-write
gate, redacted synced Linear write evidence counts, and a redacted
KnowledgeIndex evidence summary only.
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
