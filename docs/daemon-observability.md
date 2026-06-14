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

## Validation

Current validation covers:

- successful request recording
- bounded buffer behavior
- handler error recording
- query data redaction from recorded paths
- Flutter snapshot loading of the event count
