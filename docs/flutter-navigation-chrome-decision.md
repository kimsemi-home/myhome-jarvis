# Flutter Navigation Chrome Decision

Date: 2026-07-07

## Decision

Keep `MaterialApp`, `Scaffold`, `AppBar`, `TabBar`, and `TabBarView` as the
Flutter navigation chrome under `ShadApp.custom` for the current migration
phase.

`shadcn_ui` remains the component baseline for screen surfaces, badges, icon
actions, payload controls, dialogs, and other agent-readable widgets. The
navigation shell should move only after a later pass can replace it without
changing tab semantics, daemon snapshot loading, redaction boundaries, or
widget-test coverage.

Meta Astryx is the design-system reference for the shadcn token bridge. Because
Astryx is currently React/StyleX based, the Flutter app mirrors its neutral
theme values into `JarvisAstryxTokens` instead of importing Astryx npm packages
at runtime.

## Why This Stays Material For Now

- The existing tab shell is stable and already maps directly to the
  `JarvisSnapshot` screen contracts.
- `ShadApp.custom` gives the app shadcn tokens and components without forcing a
  full app-shell rewrite.
- The current Material `TabBar` keeps keyboard, scroll, and tab semantics
  predictable while the migration focuses on read-only and dry-run screen
  surfaces.
- The refresh action already crosses the boundary through `JarvisIconAction`,
  which wraps a shadcn icon button without changing daemon behavior.
- Keeping navigation stable lowers risk while Connectors remains the
  low-risk shadcn pilot tab.

## Next Migration Gate

Revisit navigation chrome only when all of these are true:

- a shadcn-style tab or shell wrapper is selected for Flutter and documented;
- Status, Commands, Connectors, and one dashboard tab have focused widget tests
  covering mounted shadcn controls and user-visible states;
- the replacement preserves all tab labels, refresh behavior, loading
  indicator behavior, offline fallback behavior, and redaction contracts;
- `flutter analyze`, `flutter test`, `go run ./cmd/mhj code-shape status`, and
  `go run ./cmd/mhj security check` stay green.

## Test Pin

`apps/flutter/test/widget_shadcn_test.dart` intentionally asserts both sides of
the bridge:

- `ShadApp` and `ShadAppBuilder` mount at the app root.
- `JarvisShadTheme` resolves its core shadcn colors and radius from
  `JarvisAstryxTokens`.
- `MaterialApp`, `Scaffold`, `AppBar`, `TabBar`, and `TabBarView` remain the
  current navigation chrome.
- `JarvisIconAction` and `ShadIconButton` keep the refresh action on the shadcn
  component path.
- `apps/flutter/test/widget_status_test.dart` pins the Status tab by asserting
  shadcn surfaces, badges, and user-visible status states.
- `apps/flutter/test/widget_commands_test.dart` pins the Commands tab by
  asserting shadcn surfaces, badges, dry-run state, blocked execution state,
  editable payload fields, selects, and inputs.
- `apps/flutter/test/widget_integrations_test.dart` pins Linear and Storage
  shadcn surfaces, badges, queued/verified state, local fixture state, and
  storage format state.
- `apps/flutter/test/widget_cluster_household_test.dart` pins Cluster shadcn
  surfaces, badges, and active/gated/tracked state tones.
- `apps/flutter/test/widget_finance_purchases_test.dart` pins one dashboard
  gate by asserting Finance/Purchases shadcn surfaces, badges, and visible
  dashboard states.
