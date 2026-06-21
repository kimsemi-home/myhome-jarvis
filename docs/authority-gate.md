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
go run ./cmd/mhj authority-review brief
go run ./cmd/mhj authority-review request
go run ./cmd/mhj authority-review evidence
go run ./cmd/mhj authority-review queue
go run ./cmd/mhj authority-review record '<json-payload>'
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

`mhj authority-review brief` returns the reviewer-facing public handoff for
gated vision work. It combines the current request, evidence ref, queue state,
required review classes, command-center work item, repo-factory gate,
repo-factory preflight summary, local runtime readiness, merge-evidence
posture, Codex sustainability posture, cross-repo context-pack posture,
capability readiness, and next safe action. It is only a handoff artifact:
approval, external writes, repo creation, and self-approval remain explicitly
false.

`mhj authority-review decision-packet` returns the public-safe review packet for
the human decision point. It excludes raw evidence refs and private payloads,
keeps `review_only` as the packet state, and includes explicit non-granting
decision options. The packet also includes the redacted storage archive summary
so reviewers can see that private local logs follow the compress-then-archive
pattern and that the evidence-noise budget configuration is itself hashed as
evidence. It includes the public-safe merge-evidence posture so reviewers can
see the default merge-after-checks behavior, post-merge evidence requirement,
Linear completion requirement, main quality run requirement, and private-data
scan requirement. The repo-factory preflight portion reports creation decision,
creation allowed, blocking gate count, missing evidence keys, and the preflight
next safe action without exposing private template payloads or granting repo
creation. When public-safety checks are green, the preflight summary should
show `public_safety_evidence` as ready so human authority review remains the
visible repo-creation blocker.
The packet also includes the public-safe Codex sustainability posture so
reviewers can see whether heavier Codex use is sustainable, whether trend
evidence is fresh, whether the loop is on trend, and whether cache savings or
review debt support the next automation decision.
It also includes the public-safe context-pack posture so reviewers can see the
pack id, upstream compatibility version, ontology version, authority/security
contract versions, and verification profile that downstream repositories must
declare before acting independently.
The packet also includes capability readiness for media playback, household
finance consent, monetization, and Codex cost governance. It exposes only
states, counts, gate posture, and next-safe-action fields so reviewers can
judge the assistant vision without raw media sessions, finance records,
private ledgers, prompts, local absolute paths, or private Linear URLs.
The packet also includes a public-safe `decision_contract` that names the
human-only reviewer posture, ready capabilities that do not block the review,
each gated capability that needs a separate human decision record, evidence
keys the reviewer must inspect, and grant flags that must remain false. This
makes the review handoff actionable without turning the packet into approval.

Daemon `GET /authority-review/request` returns a public-safe request packet with
a stable request id derived from review classes and counts. It is evidence for a
human review request only: it never grants approval, self-authority, or external
write permission.

Daemon `GET /authority-review/evidence` returns the evidence-ready reference
for that request packet. The reference is stable enough for Linear/GitHub
comments, but remains a request artifact and always reports `not_approved`.
Before a matching request is recorded, its next safe action is to request
authority review; after a matching pending ledger entry exists, it aligns with
brief and queue surfaces by reporting `await_human_authority_review`.

Daemon `GET /authority-review/queue` returns the public-safe queue item state for
the request. A queued item means a human review is pending, not approved, and it
does not enable external writes or self-approval.

`mhj authority-review record <json-payload>` appends the current request to
`data/private/authority-review/requests.jsonl`. The payload must echo the
current request id, evidence ref, queue item ref, queue state, required review
classes, and explicit `false` values for approval, external writes, and
self-approval. The public command output is summary-only: it reports the request
id, queue state, class count, approval state, and private ledger state, but does
not expose evidence refs or queue item refs. Recording a request is not an
approval and does not unlock repo creation, external writes, or self-approval.
After the current request has a matching private ledger entry, public authority
review status reports `review_request_ledger_state` as
`recorded_pending_review` and the command center switches the next safe action
from `request_authority_review` to `await_human_authority_review`.
Recorded review requests also carry a public-safe staleness guard. Status,
brief, and decision-packet surfaces report the request age, stale threshold, and
escalation action so `await_human_authority_review` cannot sit unnoticed. The
guard may request escalation or request refresh of malformed review evidence,
but it never grants approval or exposes reviewer identities, private Linear
URLs, raw ledger rows, prompts, transcripts, secrets, or local absolute paths.

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
- review request age, stale threshold, stale boolean, and escalation action
- local runtime readiness for the authority review brief and decision packet
- merge evidence posture for the authority review brief and decision packet
- Codex sustainability posture for the authority review brief and decision packet
- context pack posture for the authority review brief and decision packet
- capability readiness for media, finance, monetization, and Codex cost in the
  authority review brief and decision packet
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
go run ./cmd/mhj authority-review brief
go run ./cmd/mhj authority-review decision-packet
go run ./cmd/mhj authority-review request
go run ./cmd/mhj authority-review evidence
go run ./cmd/mhj authority-review queue
go run ./cmd/mhj authority-review record '<json-payload>'
go run ./cmd/mhj codegen verify
go run ./cmd/mhj ddd verify
cd apps/flutter && flutter test test/daemon_client_test.dart test/widget_test.dart
```
