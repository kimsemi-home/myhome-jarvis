# Repo Status

The closed-loop agent needs a local view of the Git worktree before it can
decide whether changes are ready to validate, commit, or push.

Surfaces:

- `mhj repo status`
- daemon `GET /repo/status`
- Flutter Status metric showing `Repo: Clean` or `Repo: Dirty`

The status model reports:

- current branch and head SHA
- upstream and origin when configured
- tracked changes
- untracked files
- ignored private paths under `data/private` or `data/lake`

The model intentionally uses repository-relative paths. It does not expose the
local checkout directory.

Validation:

```sh
go test ./internal/repo ./internal/daemon
cd apps/flutter && flutter test
```
