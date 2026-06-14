# Home Control

Home-control commands are deterministic dry-run plans.

Initial commands:

- `open_youtube`
- `open_youtube_search`
- `open_ott`
- `open_netflix`
- `open_disney_plus`
- `open_tving`
- `open_wavve`
- `open_coupang_play`
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
The service-specific OTT shortcuts are zero-payload convenience commands over
the same safe target map used by `open_ott`; they do not introduce scraping,
cookies, credentials, downloads, or DRM bypass behavior.

The Flutter client calls daemon `POST /intent` with `execute=false` to preview
plans. This keeps the UI on the dry-run side of the boundary while showing the
argv plan.

CLI and daemon command intents append a private redacted audit event under
`data/private/audit/command-intents.jsonl`. The audit records command/source,
dry-run and execute-gate metadata, counts, success, and coarse error category
only. It does not record payload JSON, argv arrays, URLs, headers, tokens, raw
errors, or local absolute paths.

Explicit execution boundary:

- CLI execution requires `MYHOME_EXECUTE=true`.
- Daemon execution requires `MYHOME_EXECUTE=true`, `mhj daemon --execute`, and a
  request body with `execute=true`.
- Execution runs argv arrays directly; it never uses shell interpolation.
- Only `open`, `osascript`, and `pmset` are allowed executables.
- Non-macOS platforms skip execution and return execution metadata instead of
  running commands.
