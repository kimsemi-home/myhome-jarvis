# Flutter

The first Flutter client lives in `apps/flutter`.

Current scope:

- Dart-only Flutter skeleton.
- `shadcn_ui` is installed as the Flutter shadcn component baseline, with the
  app wrapped in `ShadApp.custom` so existing Material navigation can migrate
  incrementally instead of being rewritten at once.
- Reference links for follow-up agent work: package
  [`shadcn_ui`](https://pub.dev/packages/shadcn_ui), official docs
  [Flutter Shadcn UI](https://mariuti.com/flutter-shadcn-ui/), API reference
  [`shadcn_ui` library](https://pub.dev/documentation/shadcn_ui/latest/shadcn_ui/),
  and upstream repo
  [`nank1ro/flutter-shadcn-ui`](https://github.com/nank1ro/flutter-shadcn-ui).
- Shared shadcn-style theme, surface, and badge wrappers live in the Flutter UI
  layer so AI agents can inspect tokens and migrated components directly.
- Screen contracts for agent-readable migration work live in
  [`docs/flutter-screen-contracts.md`](flutter-screen-contracts.md).
- The low-risk shadcn pilot decision lives in
  [`docs/flutter-shadcn-low-risk-pilot.md`](flutter-shadcn-low-risk-pilot.md);
  Connectors is the selected read-only pilot tab.
- The navigation/chrome decision lives in
  [`docs/flutter-navigation-chrome-decision.md`](flutter-navigation-chrome-decision.md);
  Material tabs and scaffold remain under `ShadApp.custom` for this phase.
- Status, command, finance, purchases, Linear, storage, household, and
  optimization tabs.
- Dry-run command rows with editable payload fields for the initial
  home-control surface.
- Explicit zero-payload OTT shortcut buttons for Netflix, Disney+, TVING,
  Wavve, and Coupang Play, plus the generic `open_ott` command.
- Static/offline fallback command buttons for volume up/down/set/mute, display
  sleep, and Mac sleep, matching the core home-control surface even without
  daemon reachability.
- Static/offline fallback payload editors for YouTube search, safe URL open,
  and generic OTT service selection.
- Daemon snapshot client for `/health`, `/commands`, `/linear/status`, and
  `/metrics`.
- Redacted Linear status rendering from `/linear/status` using sync state,
  viewer-configured boolean, team count, and repo-relative queue path only.
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
- Public-safe external evidence lake readiness appears in Connectors from the
  fixture-only metadata contract; it shows context-pack/status metadata while
  keeping raw payload import, credentials, private archives, and collector
  writes blocked.
- Runtime status rendering from `/metrics`, showing goroutine count and
  formatted heap allocation when the daemon provides those counters.
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
- Migrated shadcn-style surfaces include Status metrics, Linear/Storage list
  rows, command rows, payload text inputs, payload service select, dry-run
  preview dialog, dry-run actions, Connector
  readiness cards, Agent Cluster cards, Finance/Purchases read-only dashboard
  cards, Purchase recurring-candidate cards, Household result cards, Optimize
  recommendation cards, and dry-run preview actions; all keep the existing
  daemon/offline snapshot contracts.
- User, Spouse, and Household fixture scope switching.
- Dry-run preview client for `/intent`; command buttons always send
  `execute=false`, even though the daemon has a separately gated execution
  boundary.
- Optional Bearer token support for LAN daemon clients.
- Widget and client tests for the first local operations screens.
- Static/offline fallback command tests read `generated/commands.generated.json`
  and fail if command names or payload fields drift from the SSOT catalog.

Platform runner files are intentionally deferred until packaging or device
integration is needed.

Validation:

```sh
cd apps/flutter
flutter test
flutter analyze
```

Next UI migration order:

1. Keep `ShadApp.custom` while Material tabs, scaffold, and tests remain stable;
   see `docs/flutter-navigation-chrome-decision.md`.
2. Move future read-only cards to shared shadcn wrappers, then revisit
   scaffold/tab chrome only after the documented navigation replacement gate is
   fully met. The dashboard-test part of that gate is pinned by
   `apps/flutter/test/widget_finance_purchases_test.dart`.
3. Add focused widget tests before introducing additional interactive shadcn
   controls such as menus, sheets, or table-like views.
