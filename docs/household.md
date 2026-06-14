# Household Views

The first household view is fixture-only.

Scopes:

- `user`: records owned by the primary user fixture owner.
- `spouse`: records owned by the spouse fixture owner.
- `household`: aggregate view across all known fixture owners.

The scope switcher is available in Flutter through the Household tab. It reads
local daemon data only and does not require bank, card, commerce, or identity
credentials.

Validation:

```sh
cargo test -p mhj-core household
go test ./internal/domain ./internal/daemon
cd apps/flutter && flutter test && flutter analyze
```
