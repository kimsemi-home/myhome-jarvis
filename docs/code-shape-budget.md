# Code Shape Budget

The Code Shape Budget turns the 75-line preference into an executable guard.
It does not pretend the current repository is already small; instead, it
records existing oversized files as explicit legacy debt and blocks new or
expanded over-budget files.

## SSOT

Common Lisp owns the policy in `lisp/ssot/code-shape.lisp` and emits
`generated/code_shape.generated.json`.

The generated policy defines:

- the 75-line max for ordinary source files
- scanned source roots and extensions
- ignored generated, private, build, and dependency paths
- current legacy-debt file baselines
- public summary fields and forbidden public fields

## Status

```sh
go run ./cmd/mhj code-shape status
```

The command scans Go, Dart, Lisp, and Rust source roots. Files above 75 lines
are allowed only when listed in the generated legacy-debt baseline and still at
or below that baseline. Any new oversized file or growth beyond its baseline is
a budget regression and makes the command fail.

When a legacy file is split or shortened, its generated baseline should be
ratcheted down in the same change so the debt cannot silently grow back.

Daemon `GET /code-shape/status` returns the same redacted status. Flutter shows
that summary as a `Code Shape` metric.

## Public Surface

Status output uses repo-relative paths only. It exposes counts, max observed
file, top legacy-debt entries, regressions, and timestamps. It does not expose
local absolute paths, source excerpts, credentials, tokens, Linear URLs, or
private evidence.

## Validation

```sh
sbcl --script lisp/scripts/validate-ssot.lisp
go run ./cmd/mhj codegen verify
go run ./cmd/mhj code-shape status
go test ./internal/codeshape ./internal/daemon ./cmd/mhj
```

GitHub Actions runs the code-shape guard in the Go unit. Flutter's unit cache
also keys on `generated/code_shape.generated.json`, because the local client
shows that generated status shape.
