# Closed Loop

Each autonomous cycle should:

1. Determine Linear status or offline fallback.
2. Inspect repository state.
3. Pick one small task.
4. Record a working-log start entry.
5. Modify one file or a tightly connected set of files.
6. Run the relevant quality gate.
7. Record results and checkpoint evidence.

The initial `loop once` command records a local checkpoint and never loops
forever.
