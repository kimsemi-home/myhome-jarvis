# Flutter Shadcn Low-Risk Pilot

The selected low-risk pilot tab is **Connectors**.

## Why Connectors

- It is read-only and already backed by `JarvisSnapshot.connectors`.
- It has no daemon command execution path.
- It is a natural boundary for public-safe status metadata.
- It can display blocked operations explicitly without requesting credentials,
  importing raw evidence, or writing to collectors.
- It includes the External evidence lake card, which exercises the cross-repo
  public metadata contract.

## Pilot Contract

The Connectors tab may render:

- connector label, category, status, and fixture mode;
- public-safe data classes;
- allowed metadata-only operations;
- blocked or forbidden operations;
- the next safe setup step.

The Connectors tab must not:

- request credentials;
- call external APIs;
- import raw payloads;
- read or display private archives;
- write collectors;
- execute daemon commands.

## Current Implementation

- `apps/flutter/lib/ui/connectors_view.dart` owns the tab list.
- `apps/flutter/lib/ui/connector_tile.dart` renders each read-only card.
- `JarvisSurface`, `JarvisBadge`, and `JarvisBadgeWrap` provide the shared
  shadcn-style surface and badge primitives.
- `apps/flutter/lib/snapshot/sample_connectors.dart` includes the fixture-only
  External evidence lake status card.
- `apps/flutter/test/widget_connectors_test.dart` pins the External evidence
  lake connector fields to the public `upstream_connector` shape from
  `myhome-external-evidence-lake/fixtures/public-ui-status.sample.json`.

Material tab navigation stays in place for this pilot. The migrated scope is
the tab content: cards, badges, labels, and public-safe blocked-state text.

## Verification

Focused widget coverage lives in `apps/flutter/test/widget_connectors_test.dart`
and proves:

- the tab mounts shared shadcn wrappers;
- External evidence lake fixture metadata is visible;
- allowed operations stay metadata-only;
- raw payload import, credentials, private archives, and collector writes remain
  blocked.
