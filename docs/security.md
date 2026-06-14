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
- Reject Python, Node.js, TypeScript, shell-interpolated command execution, and
  tracked private-data artifacts.

The Go security checker is the first enforceable guard. It scans path names and
language/dependency files and exits non-zero when a forbidden file is present.
