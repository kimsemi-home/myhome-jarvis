# CI

GitHub Actions is split into hash-scoped units:

- `public-safety`: current-tree and full Git-history public-safety checks.
- `ssot`: Common Lisp SSOT sources, SSOT loader scripts, and generated JSON artifacts.
- `go`: Go CLI, daemon, internal packages, and generated JSON inputs.
- `rust`: Rust workspace crates and fixture inputs.
- `flutter`: Flutter local client files and generated command catalog input.

The public-safety job always runs because every new commit can introduce a
public-history risk, even when the Go/Rust/Flutter/SSOT source hashes are
unchanged. It fetches full history, then runs `mhj security check` and
`mhj security history`.

The workflow cancels superseded in-progress runs for the same ref. The latest
push remains authoritative, while older queued or running checks stop instead
of burning runner time after a newer commit replaces them.

The SSOT, Go, Rust, and Flutter units restore marker caches keyed by each unit's
input hash. If the exact hash is already known-good, the unit reports a cache
hit and skips its heavy toolchain setup and tests. A lightweight workflow run
still exists for each push so GitHub can report status, but unchanged units
avoid repeated work.

Unit caches restore for push and pull-request runs, but new unit cache markers
are saved only from push events in the canonical `kimsemi-home/myhome-jarvis`
repository. Pull requests still verify cache misses, but they cannot publish
new known-good markers for later runs.

Workflow-maintained action refs use Node 24-capable releases:
`actions/checkout@v6`, `actions/setup-go@v6`, and `actions/cache@v5`.
Because `.github/workflows/quality.yml` is part of every unit cache key, action
ref changes intentionally invalidate the SSOT, Go, Rust, and Flutter unit
caches once so the new runner surface is verified before future cache hits.
Rust also has a checked-in `rust-toolchain.toml`; the Rust unit cache key
includes that file so compiler or component changes rerun Rust tests before a
new marker can be saved.

Generated artifact verification lives in the `ssot` unit. On a cache miss, CI
runs SSOT validation, regenerates artifacts, and fails if `generated` differs
from the checked-in files.

Locally, `mhj codegen verify` verifies the current working tree rather than
HEAD: it snapshots `generated`, regenerates from Lisp, and fails only if
regeneration changes the generated tree. This lets intentional SSOT/generated
updates pass before commit while still catching stale artifacts.

`mhj quality` also includes a `toolchain pins` step. It fails when the Go
version in `.go-version`, `go.mod`, generated project metadata, or workflow
`GO_VERSION` drift, and when `rust-toolchain.toml` differs from workflow
`RUST_TOOLCHAIN`.

The Go unit runs `home`, `finance`, and `commerce` harness smoke commands
before package tests and vet. Public-safety checks live in their own always-run
job so docs-only or metadata-only risks are not hidden by the Go unit cache.
The Rust unit runs the whole workspace, including `mhj-harness`, so the
dedicated Rust harness boundary is covered whenever command, finance, commerce,
fixtures, or Rust harness inputs change.
The Flutter unit also keys on `generated/commands.generated.json` because its
static/offline fallback tests read the Lisp-owned command catalog directly.

Local equivalents:

```sh
go run ./cmd/mhj codegen verify
go run ./cmd/mhj quality
go run ./cmd/mhj quality status
```

`mhj quality` also writes a redacted private run summary under
`data/private/quality/runs.jsonl` for closed-loop evidence. The journal stores
step names and statuses only, not command output or absolute local paths.
The default `mhj quality` JSON output follows the same redacted shape: overall
status plus step names/statuses, without command argv or raw command output.
