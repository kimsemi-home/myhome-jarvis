# Assistant Authority Profiles

Assistant authority profiles map each long-running capability pillar to the
gates it must pass before the assistant can act. The SSOT source is
`lisp/ssot/authority-profiles.lisp`; the generated policy lives inside
`generated/authority.generated.json`.

Profiles currently cover:

- `local_media_concierge`: local interactive status and deterministic checks.
- `household_finance_copilot`: review-only finance evidence; no external
  writes.
- `shorts_factory_control_plane`: public repo and workflow changes require
  review plus public-safety gates.
- `monetization_console`: revenue experiments require review and public-safety
  gates before publishing or external automation.
- `codex_cost_governor`: local read-only cost metadata.
- `self_improvement_loop`: authority-gated improvement with verifier
  separation.

All profiles keep `self_approval_allowed=false`. Public status may expose only
profile counts and profile keys by gate class. It must not expose raw rationale,
raw evidence, evidence refs, prompts, transcripts, tokens, credentials, local
absolute paths, private Linear URLs, or private evidence contents.
