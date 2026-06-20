# Public-Safe Repo Factory

`PublicSafeRepoFactory` defines the template and gates for creating public
GitHub repositories and Codex projects for Shorts Factory workflows.

The factory is intentionally not a repo creator yet. It is a public-safe SSOT
contract that says which files must exist, which checks must run, and which
review evidence is required before automation can create or mutate a public
repo.

## Template

Each new public repo must start from the generated factory policy:

- generated CI: `.github/workflows/quality.yml`
- security scan policy: `docs/security.md`
- private-data policy: `docs/private-data-policy.md`
- bootstrap checklist: `docs/bootstrap-checklist.md`
- Codex project goal: `.codex/project-goal.md`

## Creation Gates

Repo creation remains blocked until all gates have evidence:

- authority review is approved
- authority review request evidence is recorded privately
- public safety evidence is recorded
- generated CI exists
- private-data policy exists
- bootstrap checklist is complete

The private authority review request ledger proves the review was requested and
queued. It does not count as approval; repo creation stays blocked until a human
review grants the relevant public-repo authority.

## Public Boundary

Generated public files must not include private assets, credentials, local
absolute paths, raw prompts, household finance payloads, or private media
assets. Private working material stays outside the generated public template.

Use `mhj repo-factory status` to inspect the public-safe summary.

## Preflight Decision Packet

Use `mhj repo-factory decision-packet` before any public repository or Codex
project is created. The packet is a public-safe review input, not a creation
approval. It lists template evidence, gate evidence, missing review evidence,
and the next safe action while keeping `creation_allowed` false until human
authority review and public-safety evidence are recorded.

The packet also reads the redacted public-safety status. When `mhj security
check` and `mhj security history` both pass, the `public_safety_evidence` gate
is marked `ready`; raw findings, cache keys, local paths, and credentials are
not printed. Repo creation still stays blocked until human authority review is
approved separately.

The packet also includes `context_pack_evidence`, a public-safe summary from
`mhj context-pack verify`. It reports declaration path, validity, drift/missing
counts, context pack version, ontology version, authority/security contract
versions, verification profile, and exported artifact count. It does not expose
raw private context, local absolute paths, private Linear URLs, credentials, or
review payloads, and a valid context pack never grants repo creation authority.

If a generated policy ever contains a forbidden template value, the packet
marks the affected template as invalid and redacts the value instead of
printing local paths, private owner markers, credentials, or other private
payloads.
