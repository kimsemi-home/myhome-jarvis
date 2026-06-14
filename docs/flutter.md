# Flutter

The first Flutter client lives in `apps/flutter`.

Current scope:

- Dart-only Flutter skeleton.
- Status, command, finance, purchases, Linear, storage, household, and
  optimization tabs.
- Dry-run command rows with editable payload fields for the initial
  home-control surface.
- Explicit zero-payload OTT shortcut buttons for Netflix, Disney+, TVING,
  Wavve, and Coupang Play, plus the generic `open_ott` command.
- Static/offline fallback command buttons for volume up/down/set and display
  sleep, matching the core home-control surface even without daemon reachability.
- Static/offline fallback payload editors for YouTube search, safe URL open,
  and generic OTT service selection.
- Daemon snapshot client for `/health`, `/commands`, `/linear/status`, and
  `/metrics`.
- Domain summary rendering from `/domain/summary` for finance, commerce, and
  storage fixture status.
- Dedicated fixture-only Finance tab for cashflow totals, subscription spend,
  card-linked debit review totals, categories, and owner breakdowns.
- Dedicated fixture-only Purchases tab for commerce spend, recurring purchase
  candidates, categories, and owner spend breakdowns.
- Repository status rendering from `/repo/status` as a clean/dirty status
  metric.
- Local-only network mode rendering from `/health` and `/metrics`, with
  LAN-enabled daemon mode shown as token-gated.
- LAN auth status rendering from `/auth/status` without displaying token
  contents.
- Supervisor status rendering from `/supervisor/status` as a reachable/stale
  daemon process metric.
- Command audit rendering from `/audit/status` as a redacted journal count.
- Quality evidence rendering from `/quality/status` as the latest quality gate
  result.
- Public-safety rendering from `/security/status` as aggregate current-tree and
  Git-history status without raw findings or local roots.
- Planner status rendering from `/planner/status` as completed/ready/gated task
  graph progress plus the first external-write-gated task id.
- Structured recommendation rendering from fixture-only local summaries,
  including score, rationale, estimated amount, evidence count, and
  card-linked spend review items that never execute card actions.
- User, Spouse, and Household fixture scope switching.
- Dry-run preview client for `/intent`; command buttons always send
  `execute=false`, even though the daemon has a separately gated execution
  boundary.
- Optional Bearer token support for LAN daemon clients.
- Widget and client tests for the first local operations screens.

Platform runner files are intentionally deferred until packaging or device
integration is needed.

Validation:

```sh
cd apps/flutter
flutter test
flutter analyze
```
