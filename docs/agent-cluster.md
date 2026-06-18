# Agent Cluster Policy

The first Agent Cluster slice is a public-safe governance bootloader, not an
external multi-agent runtime.

Common Lisp SSOT owns the policy in `lisp/ssot/agent-cluster.lisp` and emits
`generated/agent_cluster.generated.json`. Go reads the generated artifact for
`mhj agent-cluster status` and daemon `GET /agent-cluster/status`. Flutter
renders the daemon status as read-only Cluster signals and keeps its static
fallback aligned with the generated artifact.

The policy fixes the current learning-loop order:

```text
reality -> observation -> evidence -> interpretation -> claim -> rulebook
-> design -> code -> verification_evidence -> knowledge_update
```

This keeps changes evidence-first. Code can still change quickly, but a missing
or delayed evidence step must become tracked debt instead of invisible drift.

## Authority

Agent outputs are treated as candidates, not truth. The SSOT separates roles:

- producer agents propose changes and verification plans.
- independent reviewers check meaning, risks, and ontology mapping.
- adversarial reviewers look for missing evidence, authority violations, and
  rollback risks.
- deterministic verifiers run commands and artifact checks.
- governance stewards gate authority, ownership, and revalidation records.

No role may approve its own output or assign final confidence to itself.
Authority gates stay enabled while this repo is public.

## Sidecars

The first sidecar catalog is declarative and read-only:

- verification sidecar for contracts, tests, and generated artifacts.
- confidence assessor for evidence links, coverage, and reliability signals.
- evidence quality assessor for age, schema version, and mapping confidence.
- security audit sidecar for public safety and authority boundaries.
- translation verifier for context movement and loss ledgers.
- control-plane verifier for manifests, policy hashes, and leases.

The sidecars are not yet separate processes. They are executable policy
contracts used by Go, tests, CI, docs, and Flutter status.

## Public Safety

This phase forbids external agent execution, raw transcript storage, private
data in public evidence, self-approval, and self-reported final confidence.
The public artifact contains only roles, stages, sidecar names, failure
conditions, debt classes, and read-only status signals.

Private operational evidence belongs under `data/private` and is exposed only
through redacted counts or repo-relative paths.

The Learning Ledger is the first executable follow-through for this policy. It
records loop gaps and evidence debt under `data/private/learning` and exposes
only redacted counts through `mhj learning status` and daemon
`GET /learning/status`.

The Evidence Graph is the next executable follow-through. It connects private
learning observations to referenced evidence artifacts and exposes only redacted
node, edge, source, dangling-ref, and timestamp counts through
`mhj evidence status` and daemon `GET /evidence/status`.

The Confidence Assessor is the first executable confidence gate. It reads the
redacted Evidence Graph, quality evidence, public-safety status, and open
learning debt, then returns a confidence cap. Agent self-reported confidence is
not accepted.

## Validation

Use these checks after changing the policy:

```sh
go run ./cmd/mhj agent-cluster status
go run ./cmd/mhj evidence status
go run ./cmd/mhj confidence status
go run ./cmd/mhj codegen verify
go run ./cmd/mhj ddd verify
go test ./internal/agentcluster ./internal/daemon ./cmd/mhj
cd apps/flutter && flutter test test/daemon_client_test.dart test/snapshot_test.dart test/widget_test.dart
```
