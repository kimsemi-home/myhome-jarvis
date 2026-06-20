package monetization

import (
	"strings"
	"time"
)

func updateLastObserved(status *Status, candidate string) {
	if candidate > status.LastObservedAt {
		status.LastObservedAt = candidate
	}
}

func normalizeRecordedAt(value string, now time.Time) (string, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return now.UTC().Format(time.RFC3339), nil
	}
	parsed, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return "", err
	}
	return parsed.UTC().Format(time.RFC3339), nil
}
