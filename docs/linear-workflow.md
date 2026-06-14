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

Mutation commands use GraphQL variables rather than string interpolation.
When credentials are unavailable or a GraphQL call fails, the command writes a
structured `synced=false` event to `data/private/linear-offline-queue.jsonl`.
