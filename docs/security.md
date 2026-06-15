# Security

Security defaults:

- Bind local services to `127.0.0.1` unless LAN bind is explicitly configured.
- Require a local Bearer token for non-localhost daemon requests.
- Keep command execution dry-run unless `MYHOME_EXECUTE=true` or an equivalent
  local private config is set.
- Require daemon `--execute` and per-request `execute=true` before daemon
  command execution.
- Execute only validated argv arrays for `open`, `osascript`, and `pmset`; never
  execute through a shell.
- Store local tokens only under `data/private` with private file permissions.
- Keep raw data and lake data under ignored private paths.
- Keep connector readiness fixture-only until a separate consent and vault
  design is implemented.
- Reject Python, Node.js, TypeScript, shell-interpolated command execution, and
  tracked private-data artifacts.

The Go security checker is the first enforceable guard. `mhj security check`
scans the current working tree's path names, language/dependency files, and
non-private file contents. It exits non-zero when forbidden files, private
identity markers, local absolute paths, or secret-looking literals are present
before commit. Findings report repo-relative path, optional line number, code,
and a coarse message only; matched secret contents are not returned. Its report
uses `root: "."` so default CLI output does not expose the local checkout path.
The executable SSOT records this current-content scan contract in
`generated/security.generated.json`, and Go tests fail if that generated policy
drifts away from the scanner behavior.

`mhj security history` scans reachable Git commits before public pushes. It
checks historical file paths, historical file contents, and commit metadata for
private identity markers, local absolute paths, forbidden Python/Node/TypeScript
artifacts, private/lake data paths except empty keep placeholders,
sensitive-looking file names, and secret-looking literals. Findings report
commit, repo-relative path, line number, code, and a coarse message only;
matched secret contents are not returned. Its report root is also redacted as
`.`.

Daemon `GET /security/status` runs the current-tree and Git-history checks for
local status surfaces, then returns only aggregate booleans, finding counts, and
a checked timestamp. It does not return raw findings, matched content, or the
local repository root.

Closed-loop checkpoints use the same aggregate public-safety status. This keeps
private scheduler evidence useful for recovery without storing raw matched
findings or local roots.

GitHub Actions stays in a read-only public-repository posture. `mhj ci verify`
fails if the quality workflow adds `pull_request_target`, grants any
`*: write` permission, or removes top-level `contents: read`.
