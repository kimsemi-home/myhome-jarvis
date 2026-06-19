# Control Plane Manifest

The Control Plane Manifest is the first executable orchestration-decision
receipt for the local closed loop.

It does not run external agents and it does not publish raw routing rationale.
It records private, redacted manifests for local control-plane decisions so the
system can later answer which policy, ontology version, authority profile,
lease, verifier roles, evidence inputs, and output reference were used.

## SSOT

Common Lisp owns the policy in `lisp/ssot/control-plane.lisp` and emits
`generated/control_plane.generated.json` plus
`generated/control_plane_verification.generated.json`.

The generated policy defines:

- the private append-only manifest ledger
- required manifest fields
- allowed decision kinds
- allowed authority profiles
- lease status and lease bounds
- required reviewer/verifier separation
- public redaction fields

## Runtime

```sh
go run ./cmd/mhj control-plane status
go run ./cmd/mhj control-plane verify
```

`mhj loop once` and bounded `mhj loop worker --cycles <n>` append a private
manifest after writing their private checkpoint evidence.
`mhj control-plane verify` is the deterministic sidecar check for the generated
policy, public redaction, lease bounds, verifier separation, and manifest debt.

The public status includes only:

- manifest ledger presence
- valid and invalid manifest counts
- manifest debt count
- verifier separation requirement and violation count
- lease bounds
- counts by decision kind, authority profile, and lease status
- timestamps

It does not expose raw rationale, candidate agents, evidence refs, output refs,
prompts, transcripts, Linear URLs, tokens, credentials, local absolute paths,
account IDs, card numbers, or private evidence contents.

## Debt

Malformed manifests and verifier-separation violations are counted as manifest
debt. A missing ledger is allowed before the first local orchestration decision;
after loop execution, decisions should leave a private manifest receipt.
Manifest lines containing forbidden raw-rationale or sensitive public markers
are also counted as debt instead of being treated as valid receipts.

## Validation

Use these checks after changing the manifest policy:

```sh
go test ./internal/controlplane ./internal/daemon ./cmd/mhj ./internal/knowledge ./internal/evidence
go run ./cmd/mhj control-plane status
go run ./cmd/mhj control-plane verify
go run ./cmd/mhj loop once
go run ./cmd/mhj codegen verify
go run ./cmd/mhj ddd verify
cd apps/flutter && flutter test test/daemon_client_test.dart test/widget_test.dart
```
