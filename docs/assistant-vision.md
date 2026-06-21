# Assistant Vision

`myhome-jarvis` is moving toward a local-first household executive assistant:
it observes, plans, acts within bounded authority, verifies outcomes, and learns
from evidence. The SSOT source is `lisp/ssot/assistant-vision.lisp`; the public
generated artifact is `generated/assistant_vision.generated.json`.
`mhj assistant status` and daemon `GET /assistant/status` summarize the vision
with ready/gated/blocked pillar counts and pillar keys, so the closed loop can
see which mission areas are usable, review-gated, or blocked without exposing
private payloads.
`mhj assistant vision-audit` is the public-safe completion check for the
long-running goal. It expands the same vision into per-capability audit rows,
evidence section refs, gate refs, and a `goal_complete` verdict. The verdict
stays false until every capability pillar is ready and no command-center gate
remains open.

## Universal Language

- `intent`: a user goal captured without private payload.
- `capability`: a bounded assistant skill with authority rules.
- `evidence`: repo-relative proof for a claim.
- `decision`: an authority-gated choice with lease and reviewer.
- `work_item`: public-safe closed-loop unit of intent and evidence.
- `cost_unit`: Codex token or paid service usage unit.
- `monetization_loop`: experiment tied to revenue evidence.
- `repo_factory`: repeatable public-safe repo creation flow.
- `merge_evidence`: proof that eligible PR work reached main.
- `evidence_retention`: local log compression archive lifecycle.
- `evidence_noise_budget`: bounded low-signal evidence budget.
- `external_evidence_lake`: public-source signals collected into a private
  local lake for the self-improvement loop.
- `local_runtime_health`: daemon supervisor reachability evidence.
- `household_scope`: user, spouse, household, or shared view.

## Epics

- Universal Language and Governance: keep ontology, SSOT, policies, and Linear
  planning vocabulary aligned.
- Local Media Concierge: make YouTube and OTT usage fast, local, and reliable.
- Household Finance Copilot: summarize spouse/household finances as read-only
  review evidence before any real connector.
- Shorts Factory Repo Control Plane: create and govern new public-safe GitHub
  repositories and Codex project setup for short-form content operations.
- Monetization Console: track revenue hypotheses, experiments, unit economics,
  and review gates.
- Codex Cost Governor: monitor Codex token/coin usage and compare cost to value.
- Self-Improvement Command Center: close PDCA loops across evidence, incidents,
  learning, review, and authority.
- Security Privacy and Authority Hardening: keep public output redacted,
  local-first, bearer-token protected, and self-approval impossible.

## Guardrails

- Keep private data under `data/private` and out of generated public artifacts.
- Keep finance actions read-only and review-only until a new authority gate says
  otherwise.
- Never import cookies, credentials, raw transcripts, local absolute paths, or
  direct personal contact/payment identifiers into public repo surfaces.
- Require public-safety scans before repo creation, generated workflow changes,
  monetization publishing, or external automation.
- Track cost before scaling external or paid loops.
- Treat stale local daemon supervision as runtime health debt before scaling
  external or higher-risk automation.

## Codex Cost Governor

The first implementation is a redacted status surface over a private local
usage ledger. `mhj codex-cost status` and daemon `GET /codex-cost/status`
report budget state and aggregate counts without exposing raw prompts,
transcripts, evidence refs, private notes, tokens, credentials, or local paths.

## Authority Profiles

Assistant authority profiles now map media, finance, repo factory,
monetization, cost, and self-improvement loops to review, public-safety,
public-repo, workflow-change, and verifier-separation gates. Finance and
external publishing paths require review, and every profile keeps self-approval
disabled.

## Local Media Readiness

The Local Media Concierge pillar now has a public-safe readiness benchmark.
`mhj media-readiness status` and daemon `GET /media-readiness/status` report
YouTube/OTT command-planning latency, availability, and playback readiness
without recording browsing history, cookies, credentials, account identifiers,
raw payloads, session data, profile identifiers, or raw URLs.

## Merge Evidence

`mhj merge-evidence status` and daemon `GET /merge-evidence/status` publish the
redacted policy for merging eligible PR work and proving completion. Future
loops should treat merged main evidence, not just an open PR, as the normal done
state when checks and public-safety scans pass. The policy also requires
post-merge evidence for the main quality run, Linear completion, private-data
scan, and merge decision comment.

## Universal Work Item

`mhj work-item status` and daemon `GET /work-item/status` publish the current
closed-loop work item using universal-language fields: intent, capability,
decision, evidence, authority, guardrail, and next safe action. It is a
public-safe planning card only; it does not grant approval, external writes, or
self-approval. The same card carries `capability_readiness` for media, finance
consent, monetization, and Codex cost governance so sibling repositories and
Linear/GitHub updates can interpret ready and gated pillars without reading
private ledgers.

## Local Runtime Health

`mhj assistant status` promotes the compact supervisor summary into
`local_runtime` health evidence. Missing, stale, non-running, or unreachable
daemon supervisor state creates a `local_runtime` gate with
`repair_local_runtime_health` as the safe next action. This gate is
public-safe: it exposes only booleans, a repo-relative evidence ref, debt count,
and summary message. It does not expose process command lines, local absolute
paths, tokens, probe URLs, or request payloads.

## Completion Audit

`mhj assistant vision-audit` reports:

- universal term, Linear epic, and capability requirement counts
- ready/gated/blocked requirement counts
- the active next safe action
- local evidence retention readiness: compress-then-archive mode, gzip
  compression, private log source count, noise budget, dedupe fields, and
  config evidence hash
- local runtime health evidence for the self-improvement loop
- one row per capability pillar with evidence refs and gate refs

The audit uses summary refs such as `media_readiness`, `finance_consent`,
`repo_factory`, `codex_cost`, `storage_archive`, and `authority`. Storage
archive configuration is itself evidence through `config_evidence_sha256`, and
archive promotion is blocked when the evidence noise budget is breached.
Authority review packets also carry a `decision_contract`, which keeps human
review non-delegable, maps Shorts Factory and self-improvement gate scopes to
required evidence keys, and keeps approval, external writes, repo creation,
workflow changes, and self-approval false. The audit does not read or publish
raw private ledgers, local absolute paths, private Linear URLs, finance
payloads, prompts, transcripts, tokens, credentials, or secrets.
