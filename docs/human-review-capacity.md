# Human Review Capacity

Human Review Capacity is the executable status surface for the Agent Cluster
principle that people are not an infinite review resource. It tracks whether
the private review queue is available, constrained, or overloaded before
authority decisions are allowed to continue.

## SSOT

Common Lisp owns the policy in `lisp/ssot/review.lisp` and emits
`generated/review.generated.json`.

The generated policy defines the private queue location, thresholds, priority
order, reviewer roles, overload rules, required fields, and redaction fields.

## Runtime

```sh
go run ./cmd/mhj review status
```

The command reads the generated policy and the private queue at
`data/private/review/queue.jsonl`. A missing queue is allowed and reports an
available zero-debt state.

Daemon `GET /review/status` returns the same redacted shape. Flutter Status
renders it as a read-only Review Capacity metric.

## Capacity States

- `available`: no open review pressure is present.
- `constrained`: reviews exist, evidence or reviewer coverage is missing, or
  backup reviewer coverage is below the policy threshold.
- `overloaded`: high-risk review work is open or total open reviews exceed the
  policy threshold.

Authority Gate reads this status. When review capacity is overloaded, high-risk
or review-required decisions stay frozen while low-risk status checks,
deterministic verification, evidence collection, incident response, and
revalidation can continue.

## Public Status

Public status may expose:

- policy and queue paths as repository-relative paths
- total, open, high-risk-open, invalid, missing-evidence, and missing-reviewer
  counts
- backup reviewer count
- review debt count
- capacity state and active rule
- thresholds
- buckets by risk, status, reviewer role, and queue class
- timestamps

It does not expose raw rationale, raw review notes, reviewer identities,
reviewer names or email addresses, evidence refs, prompts, transcripts, tokens,
credentials, cookies, account IDs, card numbers, local absolute paths, private
Linear URLs, or private evidence contents.

## Validation

Use these checks after changing the policy:

```sh
go test ./internal/review ./internal/authority ./internal/daemon ./cmd/mhj ./internal/knowledge
go run ./cmd/mhj review status
go run ./cmd/mhj authority status
go run ./cmd/mhj codegen verify
go run ./cmd/mhj ddd verify
cd apps/flutter && flutter test test/daemon_client_test.dart test/widget_test.dart
```
