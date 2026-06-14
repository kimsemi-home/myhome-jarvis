# Performance

Performance claims require measurements.

Initial measurement hooks:

- Command validation latency.
- Harness duration.
- Linear request count and backoff state.
- Future storage rows/sec.

The first implementation favors deterministic structure over premature
optimization.

The first smoke gate is `mhj benchmark smoke`. It runs the Rust core fixture
pipeline repeatedly:

- finance JSONL parse and cashflow summary
- commerce JSONL parse and recurring-purchase candidate detection
- storage plan generation for raw, bronze, silver, and gold lake layers

The test is intentionally small and deterministic. It is meant to catch obvious
regressions in local parsing and planning paths, not to serve as a full
statistical benchmark suite.
