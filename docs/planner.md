# Planner

The planner task graph is owned by Common Lisp SSOT in `lisp/ssot/planner.lisp`.
Codegen emits `generated/planner.generated.json`; Go and Flutter treat that JSON
as a checked generated artifact, not as a hand-edited source.

Current local surfaces:

- `mhj planner status`
- daemon `GET /planner/status`
- Flutter Status `Planner` metric

The generated planner graph keeps Linear writes behind an explicit
`blocked_external_write` task. Local planning, safety inspection, SSOT
verification, quality gates, daemon surfaces, and Flutter surfaces can progress
without mutating Linear.

Planner status validates that checkpoint paths stay repo-relative under
`data/private`, that task ids are unique, and that dependency ids exist.

Validation:

```sh
sbcl --script lisp/scripts/validate-ssot.lisp
go run ./cmd/mhj codegen verify
go test ./internal/planner ./internal/daemon
go run ./cmd/mhj planner status
```
