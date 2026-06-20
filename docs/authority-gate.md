# Authority Gate

Authority Gate is the executable status surface for Reasoning RBAC and Domain
ABAC inside the Agent Cluster policy. It does not grant authority to an agent.
It limits what can proceed based on public safety, external confidence,
evidence quality, incident debt, control-plane debt, translation debt, and the
public-repository boundary. Human review capacity is also an input, so a
reviewer overload can require review before broader changes proceed.

Reasoning tiers may produce candidates, reviews, and verification plans, but a
tier alone never grants approval. Self-authority remains disabled. High-risk
decisions stay blocked in public repo mode even when low-risk read and
deterministic verification work is allowed.

## SSOT

Common Lisp owns the policy in `lisp/ssot/authority.lisp` and emits
`generated/authority.generated.json`.

The generated policy defines:

- required input status surfaces
- reasoning tiers
- role permissions
- domain attributes used for ABAC decisions
- decision keys and risk levels
- assistant authority profiles for media, finance, repo factory,
  monetization, cost, and self-improvement loops
- outcome names
- authority debt classes
- public redaction fields

## Runtime

```sh
go run ./cmd/mhj authority status
go run ./cmd/mhj authority-review status
```

The command reads the generated Authority Gate policy plus redacted status from
Confidence Assessor, Evidence Quality Assessor, Incident Lifecycle, Control
Plane Manifest, Translation Manifest, Human Review Capacity, and public-safety
checks.

Daemon `GET /authority/status` returns the same redacted shape. Flutter Status
renders the current outcome as a read-only Authority Gate metric.

Daemon `GET /authority-review/status` returns a read-only review plan for the
same gate. The plan can show when authority review is requestable, which public
review classes apply, and profile-level gate counts. It does not approve
high-risk work, grant self-authority, change external-write permissions, expose
reviewer identities, or publish raw rationale and evidence.

## Outcomes

- `blocked`: public safety failed, confidence is blocked or low/unknown, or a
  forbidden translation loss exists.
- `review_required`: authority debt exists in evidence quality, incidents,
  control-plane manifests, translation debt, or human review capacity.
- `limited`: low and medium public-safe work may proceed, while high-risk
  decisions remain blocked in public repo mode.

## Public Status

Public status may expose:

- policy path
- outcome and active rule
- input and decision counts
- allowed and blocked decision counts
- total authority debt count
- public repo mode boolean
- reasoning-tier and self-authority booleans
- public-safety boolean
- confidence cap
- debt counts by input surface
- allowed and blocked decision keys
- counts by risk
- assistant profile counts and gated profile keys
- timestamp

It does not expose raw rationale, raw evidence contents, evidence refs,
prompts, transcripts, tokens, credentials, cookies, account IDs, card numbers,
local absolute paths, private Linear URLs, or private evidence contents.

## Validation

Use these checks after changing the gate policy:

```sh
go test ./internal/authority ./internal/daemon ./cmd/mhj ./internal/knowledge
go run ./cmd/mhj review status
go run ./cmd/mhj authority status
go run ./cmd/mhj authority-review status
go run ./cmd/mhj codegen verify
go run ./cmd/mhj ddd verify
cd apps/flutter && flutter test test/daemon_client_test.dart test/widget_test.dart
```
