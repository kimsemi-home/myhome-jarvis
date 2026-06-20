# Cross-Repo Context Pack

`CrossRepoContextPack` is the public-safe handoff contract for repositories that
split out from myhome-jarvis.

It answers two questions:

1. When should a responsibility leave the main repository?
2. Which context, ontology, SSOT, authority, security, and verification versions
   must a downstream repository declare before it starts acting independently?

## Split Criteria

A responsibility can split when it crosses at least one boundary:

- responsibility overload
- ownership boundary
- independent release cadence
- private data boundary
- CI cost or cache impact

## Context Pack

The generated pack is `generated/context_pack.generated.json`. It exports only
public-safe metadata:

- assistant mission source
- bounded context and concept registry source
- ontology version
- SSOT artifact versions
- authority and security contract versions
- verification graph profile
- upstream compatibility version

It must not include raw prompts, credentials, local absolute paths, household
finance data, private evidence, browser/session data, or unpublished
monetization details.

## Downstream Declaration

Downstream repositories declare their consumed context in:

```text
.mhj/context-pack.json
```

The declaration must include the consumed context pack version, upstream
compatibility version, ontology version, authority/security contract versions,
verification profile, and SSOT artifact versions.
This repository also carries that declaration so the repo factory can prove
the handoff format with the same verifier it expects downstream repositories
to run.

Verification:

```sh
mhj context-pack verify
mhj context-pack verify path/to/context-pack.json
```

Any stale or incompatible pack, ontology, authority, security, verification, or
artifact version fails verification.
