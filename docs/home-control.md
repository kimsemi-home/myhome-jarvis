# Home Control

Home-control commands are deterministic dry-run plans.

Initial commands:

- `open_youtube`
- `open_youtube_search`
- `open_ott`
- `open_url`
- `volume_up`
- `volume_down`
- `volume_set`
- `volume_mute`
- `display_sleep`
- `mac_sleep`
- `movie_mode`
- `sleep_mode`

Unsafe URL schemes such as `javascript:`, `file:`, and `data:` must fail.
Volume levels are restricted to `0..100`.

The Flutter client calls daemon `POST /intent` with `execute=false` to preview
plans. This keeps the UI on the dry-run side of the boundary while showing the
argv plan.

Explicit execution boundary:

- CLI execution requires `MYHOME_EXECUTE=true`.
- Daemon execution requires `MYHOME_EXECUTE=true`, `mhj daemon --execute`, and a
  request body with `execute=true`.
- Execution runs argv arrays directly; it never uses shell interpolation.
- Only `open`, `osascript`, and `pmset` are allowed executables.
- Non-macOS platforms skip execution and return execution metadata instead of
  running commands.
