# Daemon Observability

The daemon keeps a bounded in-memory request event log for local debugging and
closed-loop safety checks.

## Event Contract

`GET /events` returns the most recent daemon request events:

```json
{
  "count": 1,
  "events": [
    {
      "at": "2026-06-14T12:00:00Z",
      "method": "GET",
      "path": "/health",
      "status": 200,
      "duration_millis": 1
    }
  ]
}
```

The log records only:

- UTC timestamp
- HTTP method
- URL path without query parameters
- response status
- duration in milliseconds
- coarse error category when a request fails

The log does not record request bodies, response bodies, headers, bearer
tokens, query strings, raw handler error text, raw private data, or local
filesystem paths.

## Limits

- The event log is process-local and resets when the daemon restarts.
- The buffer keeps the newest 100 events.
- `GET /metrics` exposes `event_count` so clients can show a lightweight
  observability signal without fetching all events.
- `GET /metrics` also exposes redacted Go runtime counters: goroutine count,
  heap allocation bytes, heap system bytes, stack in-use bytes, and GC count.
  It does not expose environment variables, local roots, request data, or raw
  command output.
- The HTTP server uses bounded resource defaults: 5s read-header timeout, 15s
  read timeout, 30s write timeout, 60s idle timeout, and 1 MiB max header bytes.
  These bounds protect the local daemon from slow or idle clients while keeping
  ordinary LAN requests comfortable.

## Validation

Current validation covers:

- successful request recording
- bounded buffer behavior
- handler error recording
- query data redaction from recorded paths
- redacted runtime counters from `/metrics`
- bounded HTTP server timeout and header-size defaults
- Flutter snapshot loading of the event count
