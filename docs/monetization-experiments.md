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
```

Daemon:

```text
GET /monetization/status
```
