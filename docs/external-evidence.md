# External Evidence Lake

`external-evidence` is the first local intake rail for outside signals used by
the self-improvement loop: news, economic indicators, trend sources, GitHub
signals, and communities.

The public repository owns only the contract:

- source descriptors and allowed source classes
- public-safe status fields
- preprocessing and archive policy
- repo split assessment

Collected payloads stay under `data/private/external-evidence`. The collector
writes hash-addressed raw files plus JSONL bronze, silver, gold, and manifest
records. Public status never returns raw payload bodies, credentials, cookies,
local absolute paths, accounts, or private notes.

Commands:

```sh
mhj external-evidence status
mhj external-evidence repo-split-decision
mhj external-evidence repo-bootstrap
mhj external-evidence collect --max-sources 2
mhj storage-archive run
```

## Repo Split

The current recommendation is
`keep_contract_in_myhome_jarvis_defer_repo_creation`.

That means `myhome-jarvis` should keep the public-safe SSOT, validation,
redacted status, archive contract, and split decision record. A future
`kimsemi-home/myhome-external-evidence-lake` repo becomes useful when adapter
count, release cadence, fixture volume, or lake implementation complexity
exceeds this repository's boundary.

Creating that repo remains authority-gated. A public repo must never contain raw
payloads, credentials, cookies, local absolute paths, or private household data.

Use `mhj external-evidence repo-split-decision` before creating any external
evidence repository. The packet compares keeping the public contract in
`myhome-jarvis` with splitting collector and lake experiments into a future
repo. It includes privacy risk, maintenance burden, SSOT handoff, context-pack
handoff, GitHub Actions cost, archive/cache behavior, and ontology/version
discovery evidence.

The packet is review-only. It keeps repo creation, external writes, workflow
changes, and self-approval false until an authority decision record explicitly
grants the creation path. Raw, bronze, silver, and gold lake payloads remain
private regardless of the option selected.

Use `mhj external-evidence repo-bootstrap` after the split decision packet. It
joins the split packet, repo-factory preflight, context-pack handoff, and
authority approval status into a single public-safe bootstrap packet. Without an
active `repo_creation` approval lease for the exact candidate repo, the packet
keeps `creation_allowed=false` and reports
`blocked_missing_repo_creation_approval`.

When approval is present, the allowed surface is still only the minimal public
skeleton: README/license policy, `.codex` goal, `.mhj` context-pack declaration,
security/private-data docs, and generated GitHub Actions quality workflow. The
packet exposes hash-cache inputs for generated artifacts, source descriptors,
workflow dependencies, context-pack version, and ontology version so unchanged
bootstrap units can skip unnecessary GitHub Actions work.
