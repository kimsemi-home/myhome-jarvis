# PDCA Cycle Manifest

The PDCA Cycle Manifest turns the Agent Cluster learning loop into an
executable, public-safe contract.

It does not publish raw observations, prompts, transcripts, evidence refs,
local paths, Linear URLs, command output, credentials, account identifiers, or
private evidence. Those details stay under `data/private`.

## SSOT

Common Lisp owns the policy in `lisp/ssot/pdca.lisp` and emits
`generated/pdca.generated.json`.

The four generated steps are:

- plan: read planner status and decide the bounded work slice.
- do: run the local control-plane loop.
- check: verify the generated Verification EvidenceOps contract.
- act: update the Learning Ledger surface for knowledge updates.

## Command

```sh
go run ./cmd/mhj pdca status
```

The command validates the generated policy, checks that each step artifact is
present, and summarizes any private cycle ledger only as counts and status
buckets.

## Verification

`mhj verification verify` requires the PDCA manifest, conformance link, release
evidence, and generated test manifest entry. The generated GitHub Actions,
GitLab CI, local Makefile, and Bazel backend also run `mhj pdca status` through
the Go verification unit.

## Public Boundary

Public status fields are limited to policy path, repo-relative ledger path,
counts, status buckets, readiness booleans, and timestamps. Raw cycle contents
remain private and append-only.
