# CI

GitHub Actions is split into hash-scoped units:

- `ssot`: Common Lisp SSOT sources, SSOT loader scripts, and generated JSON artifacts.
- `go`: Go CLI, daemon, internal packages, and generated JSON inputs.
- `rust`: Rust workspace crates and fixture inputs.
- `flutter`: Flutter local client files.

Each unit restores a marker cache keyed by the unit's input hash. If the exact
hash is already known-good, the unit reports a cache hit and skips its heavy
toolchain setup and tests. A lightweight workflow run still exists for each push
so GitHub can report status, but unchanged units avoid repeated work.

Generated artifact verification lives in the `ssot` unit. On a cache miss, CI
runs SSOT validation, regenerates artifacts, and fails if `generated` differs
from the checked-in files.

Local equivalents:

```sh
go run ./cmd/mhj codegen verify
go run ./cmd/mhj quality
go run ./cmd/mhj quality status
```

`mhj quality` also writes a redacted private run summary under
`data/private/quality/runs.jsonl` for closed-loop evidence. The journal stores
step names and statuses only, not command output or absolute local paths.
