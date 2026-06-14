package daemon

import (
	"net/http"
	"sync"
	"time"
)

const maxRequestEvents = 100

type requestEvent struct {
	At             string `json:"at"`
	Method         string `json:"method"`
	Path           string `json:"path"`
	Status         int    `json:"status"`
	DurationMillis int64  `json:"duration_millis"`
	Error          string `json:"error,omitempty"`
}

type eventLog struct {
	mu     sync.Mutex
	max    int
	events []requestEvent
}

func newEventLog(max int) *eventLog {
	if max <= 0 {
		max = maxRequestEvents
	}
	return &eventLog{max: max}
}

func (log *eventLog) add(event requestEvent) {
	log.mu.Lock()
	defer log.mu.Unlock()
	if len(log.events) == log.max {
		copy(log.events, log.events[1:])
		log.events[len(log.events)-1] = event
		return
	}
	log.events = append(log.events, event)
}

func (log *eventLog) snapshot() []requestEvent {
	log.mu.Lock()
	defer log.mu.Unlock()
	events := make([]requestEvent, len(log.events))
	copy(events, log.events)
	return events
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (recorder *statusRecorder) WriteHeader(status int) {
	recorder.status = status
	recorder.ResponseWriter.WriteHeader(status)
}

func (recorder *statusRecorder) Write(body []byte) (int, error) {
	if recorder.status == 0 {
		recorder.status = http.StatusOK
	}
	return recorder.ResponseWriter.Write(body)
}

func (recorder *statusRecorder) statusCode() int {
	if recorder.status == 0 {
		return http.StatusOK
	}
	return recorder.status
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
