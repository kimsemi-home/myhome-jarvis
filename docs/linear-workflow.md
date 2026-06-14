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

Mutation commands use GraphQL variables rather than string interpolation.
When credentials are unavailable or a GraphQL call fails, the command writes a
structured `synced=false` event to `data/private/linear-offline-queue.jsonl`.
