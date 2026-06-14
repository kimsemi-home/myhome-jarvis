# Executable SSOT

The source of truth is Common Lisp code under `lisp/ssot`.

Generated JSON files under `generated` are artifacts, not source of truth.
Codegen must be deterministic: the same SSOT input should produce byte-for-byte
identical output.

Current SSOT boundaries are intentionally separated by domain:

- `project`: repository policy such as allowed languages and Go version.
- `commands`: dry-run home command catalog.
- `finance`, `commerce`, `storage`: local fixture and lakehouse domain policy.
- `household`, `recommendations`, `scheduler`: local household views, optimization hints, and bounded loop policy.
- `security`, `linear`, `planner`: public-safety rules, Linear workflow rules, and planning metadata.

The planner SSOT emits `generated/planner.generated.json`. Go reads that
artifact for `mhj planner status` and daemon `GET /planner/status`; Flutter
only consumes the daemon status. This keeps task graph shape, Linear templates,
quality requirements, and external-write boundaries in one Lisp-owned source.

The command SSOT emits `generated/commands.generated.json`. Go keeps the
runtime command registry and macOS execution planning in `internal/commands`,
but its tests load the generated artifact and fail if command names, summaries,
payload fields, OTT service allowlists, or generated URL targets drift from the
Lisp-owned catalog.

Use `mhj codegen verify` before committing SSOT or generated artifact changes.
It snapshots the current `generated` tree, regenerates artifacts from Lisp, and
fails if regeneration changes any generated file. This verifies intended
working-tree SSOT/generated updates before they are committed.
