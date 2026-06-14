# Scheduler

The first scheduler surface is bounded and local-only.

Commands:

- `mhj loop status`
- `mhj loop worker --cycles 1`

Safety properties:

- Heartbeats are recorded every cycle.
- Next-run timestamps encode rate limiting.
- Consecutive failures use bounded exponential backoff.
- State is recovered from `data/private/scheduler`.
- Checkpoint references stay under ignored private storage.
- The worker never runs forever unless a future explicit worker supervisor
  wraps it. The current process supervisor records daemon state only.

Validation:

```sh
go test ./internal/scheduler ./internal/daemon
go run ./cmd/mhj loop worker --cycles 1
go run ./cmd/mhj loop status
```
