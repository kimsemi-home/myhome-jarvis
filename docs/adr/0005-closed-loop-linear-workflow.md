# ADR 0005: Closed-Loop Linear Workflow

## Status

Accepted

## Decision

Linear is the intended task queue. If Linear is unavailable, local work
continues with offline queue entries and explicit `synced=false` state.
