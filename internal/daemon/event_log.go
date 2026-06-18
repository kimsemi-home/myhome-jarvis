package daemon

import "sync"

const maxRequestEvents = 100

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
