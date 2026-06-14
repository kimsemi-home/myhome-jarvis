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

Use `mhj codegen verify` before committing SSOT or generated artifact changes.
It regenerates `generated` from Lisp and fails if the checked-in artifacts are
out of date.
