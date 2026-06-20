# Assistant Vision

`myhome-jarvis` is moving toward a local-first household executive assistant:
it observes, plans, acts within bounded authority, verifies outcomes, and learns
from evidence. The SSOT source is `lisp/ssot/assistant-vision.lisp`; the public
generated artifact is `generated/assistant_vision.generated.json`.

## Universal Language

- `intent`: a user goal captured without private payload.
- `capability`: a bounded assistant skill with authority rules.
- `evidence`: repo-relative proof for a claim.
- `decision`: an authority-gated choice with lease and reviewer.
- `cost_unit`: Codex token or paid service usage unit.
- `monetization_loop`: experiment tied to revenue evidence.
- `repo_factory`: repeatable public-safe repo creation flow.
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

## Codex Cost Governor

The first implementation is a redacted status surface over a private local
usage ledger. `mhj codex-cost status` and daemon `GET /codex-cost/status`
report budget state and aggregate counts without exposing raw prompts,
transcripts, evidence refs, private notes, tokens, credentials, or local paths.
