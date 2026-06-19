package evidence

import (
	"strings"
	"time"
)

func updateLastObserved(status *Status, candidate string) {
	status.LastObservedAt = laterRFC3339(status.LastObservedAt, candidate)
}

func laterRFC3339(current string, candidate string) string {
	candidate = strings.TrimSpace(candidate)
	if candidate == "" {
		return current
	}
	if current == "" {
		return candidate
	}
	currentTime, currentErr := time.Parse(time.RFC3339, current)
	candidateTime, candidateErr := time.Parse(time.RFC3339, candidate)
	if currentErr != nil || candidateErr != nil {
		if candidate > current {
			return candidate
		}
		return current
	}
	if candidateTime.After(currentTime) {
		return candidate
	}
	return current
}
