# Flutter

The first Flutter client lives in `apps/flutter`.

Current scope:

- Dart-only Flutter skeleton.
- Status, command, Linear, and storage tabs.
- Dry-run command rows with editable payload fields for the initial
  home-control surface.
- Daemon snapshot client for `/health`, `/commands`, `/linear/status`, and
  `/metrics`.
- Domain summary rendering from `/domain/summary` for finance, commerce, and
  storage fixture status.
- Dry-run preview client for `/intent`; command buttons show the daemon plan
  before any execution boundary exists.
- Widget and client tests for the first local operations screens.

Platform runner files are intentionally deferred until packaging or device
integration is needed.

Validation:

```sh
cd apps/flutter
flutter test
flutter analyze
```
