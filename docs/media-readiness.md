# Media Readiness

`MediaReadinessBenchmark` measures whether local YouTube and OTT command
planning is ready for daily use without launching a browser or persisting
private media data.

The SSOT source is `lisp/ssot/media-readiness.lisp`; the public generated
policy is `generated/media_readiness.generated.json`.

## What It Measures

- YouTube launch command planning.
- YouTube search command planning with a fixed fixture query.
- OTT launch command planning through the shared safe OTT target map.
- Local launcher availability as a coarse platform probe.

The benchmark records case id, capability, command name, availability, planning
latency, invocation count, expected host, and checked timestamp. It does not
record browsing history, cookies, credentials, account identifiers, raw payloads,
or raw URLs.

## Evidence Surfaces

```sh
go run ./cmd/mhj media-readiness status
```

Daemon evidence is available at `GET /media-readiness/status`.
