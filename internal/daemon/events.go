package daemon

import (
	"net/http"
	"time"
)

type requestEvent struct {
	At             string `json:"at"`
	Method         string `json:"method"`
	Path           string `json:"path"`
	Status         int    `json:"status"`
	DurationMillis int64  `json:"duration_millis"`
	Error          string `json:"error,omitempty"`
}

func (server *Server) recordRequestEvent(request *http.Request, status int, started time.Time, err error) {
	if server.events == nil {
		return
	}
	event := requestEvent{
		At:             started.UTC().Format(time.RFC3339Nano),
		Method:         request.Method,
		Path:           request.URL.Path,
		Status:         status,
		DurationMillis: time.Since(started).Milliseconds(),
	}
	if err != nil {
		event.Error = eventErrorLabel(status)
	}
	server.events.add(event)
}

func eventErrorLabel(status int) string {
	switch status {
	case http.StatusUnauthorized:
		return "unauthorized"
	case http.StatusBadRequest:
		return "bad_request"
	default:
		return "request_error"
	}
}
