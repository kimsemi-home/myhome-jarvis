# Monetization Experiments

Monetization Experiment Ledger tracks revenue hypotheses without publishing
private revenue details.

The generated policy is `generated/monetization.generated.json`. Private records
live in `data/private/monetization/experiments.jsonl`.

Each experiment decision must include:

- evidence refs
- a cost estimate
- an expected value band
- a review status

Public status may expose only redacted counts, experiment states, review states,
missing-evidence debt, missing-cost debt, and expected-value band counts.

Commands:

```sh
go run ./cmd/mhj monetization status
go run ./cmd/mhj monetization record '<json-payload>'
```

`mhj monetization record` appends only to the private JSONL ledger. The command
rejects unknown fields, absolute/path-traversal evidence refs, invalid enums,
missing evidence, and missing cost estimates. Public output returns a redacted
decision summary with an evidence ref count instead of evidence ref values,
private counterparties, raw revenue amounts, or private revenue notes.

The private ledger is a storage archive source. `mhj storage-archive run`
compresses it with the same gzip archive lane used for other local operational
logs, and each archive manifest row carries config evidence for the private log
source list, compression/archive settings, and evidence noise budget.

Daemon:

```text
GET /monetization/status
```
