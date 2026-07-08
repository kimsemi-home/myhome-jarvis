# Flutter Screen Contracts

This document is the agent-readable boundary for the Flutter app. UI changes
should start here before moving more surfaces to `shadcn_ui`.

The app root is `apps/flutter/lib/main.dart`. `JarvisApp` wraps the existing
Material shell in `ShadApp.custom`, and `JarvisHome` loads a `JarvisSnapshot`.
If the daemon load fails, the UI must render `JarvisSnapshot.offlineFallback()`
instead of exposing errors or local paths.

Shared visual tokens start in `apps/flutter/lib/ui/astryx_theme.dart`. The
file mirrors the public Meta Astryx neutral theme into Flutter constants, and
`apps/flutter/lib/ui/shadcn_theme.dart` maps those constants into
`ShadThemeData` plus Material interop. Astryx remains a design-system/token
reference for Flutter, not a runtime package dependency.

## Shared Snapshot

Authoritative UI input:

- `JarvisSnapshot.metrics`
- `JarvisSnapshot.commands`
- `JarvisSnapshot.linearItems`
- `JarvisSnapshot.storageItems`
- `JarvisSnapshot.financeDashboard`
- `JarvisSnapshot.purchaseDashboard`
- `JarvisSnapshot.connectors`
- `JarvisSnapshot.agentClusterSignals`
- `JarvisSnapshot.householdScopes`
- `JarvisSnapshot.recommendations`

Daemon-backed snapshots are built in `apps/flutter/lib/daemon_client/build_snapshot.dart`
from local status endpoints. Static and offline fallbacks live under
`apps/flutter/lib/snapshot/`.

## Global States

Every tab should keep these states explicit and user-visible when relevant:

- `loading`: `JarvisScaffold.loading` shows the top progress indicator.
- `offline`: `JarvisHome` renders `JarvisSnapshot.offlineFallback()` when daemon
  loading fails.
- `local-only`: network and daemon status must describe local or token-gated LAN
  mode without displaying secrets.
- `stale`: stale daemon or generated artifact state may be shown as text/badges,
  but must not trigger writes.
- `blocked`: external write, private data, and authority gates should remain
  visible instead of silently disappearing.
- `verified`: public-safe, quality, generated, or fixture status may be shown
  only as metadata.
- `redacted`: local paths, tokens, raw private findings, and payload contents
  outside the snapshot contract must stay hidden.

## Tab Contracts

| Tab | Input | User-visible states | Forbidden actions |
| --- | --- | --- | --- |
| Status | `metrics` from health, auth, repo, security, quality, planner, runtime, audit, and governance status maps. | loading, offline, local-only, stale, dirty, warning, blocked, verified. | Do not show bearer tokens, local absolute paths, raw security findings, raw audit rows, or private evidence payloads. |
| Commands | `commands` plus dry-run previews from the command client. | loading, offline fallback commands, dry-run, blocked execution, warning. | Do not execute commands from the UI; previews must keep `execute=false`. Do not add external writes without a separate gated contract. |
| Finance | `financeDashboard` fixture/domain summary. | empty, fixture-only, owner scoped, card-linked review, verified metadata. | Do not import bank payloads, expose account identifiers, trigger card actions, or write to finance ledgers. |
| Purchases | `purchaseDashboard` fixture/domain summary. | empty, fixture-only, recurring candidate, owner scoped, verified metadata. | Do not import receipts, expose merchant-private raw payloads beyond fixture summaries, or execute purchase actions. |
| Linear | `linearItems` from redacted Linear status. | offline, synced, viewer configured, queued, blocked external write. | Do not sync Linear from UI without explicit external-write approval. Do not expose tokens or raw queue file contents. |
| Storage | `storageItems` from domain summary. | empty, local fixture, format/compression, verified metadata. | Do not expose absolute storage roots, raw archives, private lake contents, or raw evidence files. |
| Connectors | `connectors` readiness catalog. | fixture-only, ready, blocked, local-only, verified public metadata. | Do not import raw external evidence, credentials, private archives, or collector writes. External evidence lake stays public status metadata only. |
| Cluster | `agentClusterSignals` governance/status signals. | active, gated, blocked, evidence-backed, stale. | Do not start agents, grant authority, or change orchestration from this tab. |
| Household | `householdScopes` selected by user, spouse, or household. | empty, fixture-only, selected scope, owner scoped. | Do not expose private person-level records beyond summarized fixture metrics. Do not write household decisions. |
| Optimize | `recommendations` and legacy `recommendationItems`. | empty, fixture-only, scored, evidence count, warning. | Do not execute savings, card, subscription, commerce, or external account actions from recommendation cards. |

## Migration Rules

- Keep `ShadApp.custom` while Material tab navigation remains stable.
- The current decision is to keep `MaterialApp`, `Scaffold`, `AppBar`,
  `TabBar`, and `TabBarView` as navigation chrome for this phase; see
  `docs/flutter-navigation-chrome-decision.md`.
- Move screen contents behind shared wrappers in `apps/flutter/lib/ui/shadcn_components.dart`
  before changing navigation chrome.
- Add or change colors through `JarvisAstryxTokens` or a semantic wrapper first
  so future agents can read the visual system without hunting for raw hex
  values.
- Treat Connectors as the first low-risk shadcn pilot tab; the decision and
  test target are documented in `docs/flutter-shadcn-low-risk-pilot.md`.
- Keep dashboard migration evidence focused: Finance/Purchases must continue
  mounting `JarvisSurface` and `ShadBadge` in
  `apps/flutter/test/widget_finance_purchases_test.dart`.
- Add focused widget tests before introducing new interactive shadcn controls.
- Keep daemon endpoint contracts and redaction behavior unchanged during UI
  migration.
- Keep generated code-shape budgets green when adding helper widgets.

## Verification

Run these before merging UI contract or migration changes:

```sh
cd apps/flutter
flutter analyze
flutter test
cd ../..
go run ./cmd/mhj code-shape status
```
