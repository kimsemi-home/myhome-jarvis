# Linear Workflow

Linear is the intended work queue. When credentials are unavailable, the system
must continue local work and record offline events under `data/private`.

Initial commands:

- `mhj linear status`
- `mhj linear sync`
- `mhj linear pull`
- `mhj linear next`
- `mhj linear comment <issue-id> <message>`
- `mhj linear transition <issue-id> <state>`
- `mhj linear create-from-backlog`
- `mhj linear replay-offline`
- `mhj loop once`

The GraphQL client will be implemented directly in Go. TypeScript SDKs are not
allowed.

The client uses `https://api.linear.app/graphql`. Personal API keys are sent as
the raw `Authorization` header value; OAuth access tokens may be supplied with
the `Bearer` prefix already present. Tokens are never printed.

`mhj linear status` and daemon `GET /linear/status` return redacted status
summaries by default. They expose mode, token configured state, sync state,
repo-relative queue path, HTTP status, rate-limit remaining count,
viewer-configured boolean, team count, and message only. `mhj linear sync`,
`mhj linear pull`, `mhj linear next`, and daemon `POST /linear/sync` also return
redacted operation summaries by default: issue identifiers, titles, update
timestamps, and state types may be shown, while raw descriptions, workspace
URLs, team identities, Linear UUIDs, token source, and absolute private paths
are kept out of default CLI and daemon surfaces.

`mhj linear pull` and `mhj linear next` keep only active/open issues before
selection. Set `LINEAR_TEAM_KEY` or `LINEAR_TEAM_ID` in private local config to
scope those commands to one Linear team. The team scope is used for filtering
only; default CLI and daemon summaries still do not print team names, team IDs,
workspace URLs, raw descriptions, or token sources.

Project-owned Linear issues use the SSOT-owned `[myhome-jarvis]` title prefix.
When active team results include both project issues and unrelated active items
such as Linear onboarding tasks, `mhj linear next` prefers the project-prefixed
issue first. When no active project-prefixed issue exists, `mhj linear next`
returns no selected issue instead of selecting unrelated active items. Locally
seeded backlog titles use the same prefix, represent current follow-up project
work, and are deduped by existing Linear issue title before any new issue is
created. If all seed titles already exist, `mhj linear create-from-backlog`
returns a synced zero-created summary instead of recreating duplicates.

Mutation commands use GraphQL variables rather than string interpolation.
When credentials are unavailable or a GraphQL call fails, the command writes a
structured `synced=false` event to `data/private/linear-offline-queue.jsonl`.
`mhj linear replay-offline` reads that private append-only queue and replays
only write-safe comment and transition actions after credentials are available.
Successful replay writes private idempotency evidence to
`data/private/linear-offline-replay.jsonl`, so repeated replay does not repeat
the same comment or transition. When `LINEAR_TEAM_KEY` is configured, replay
only processes queued issue keys with that public prefix, so older entries from
another team key stay skipped instead of being mutated. Failed entries and
entries paused by low rate-limit remaining stay `synced=false` in the original
queue; replay summaries return only counts, repo-relative private paths, coarse
status, HTTP status, rate-limit remaining, and a redacted message. Backlog
issue creation is idempotent through `mhj linear create-from-backlog` but is not
automatically replayed from queued payloads, avoiding stale issue creation.

Approved Linear write commands record private success evidence only after the
Linear API mutation succeeds. The evidence journal is
`data/private/linear-write-evidence.jsonl` and stores redacted events with
action, public issue key when available, and `synced=true` only. Queued offline
actions, failed GraphQL calls, lookup failures, and zero-created backlog syncs
do not increment planner `linear_write_evidence` synced mutation counts.

For implementation work, the default completion path is PR creation, validation,
merge when eligible, and then a Linear completion comment with public-safe
evidence. The comment should reference the PR, feature commit, merge commit,
push/PR/main quality runs, and public-safety scan result. Public repo surfaces
must keep private workspace URLs, local absolute paths, raw review notes, and
credentials out of generated artifacts and default CLI/daemon summaries.
