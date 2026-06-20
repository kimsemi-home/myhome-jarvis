package codexsustainability

import (
	"strings"
	"time"
)

func laterRFC3339(left string, right string) string {
	if strings.TrimSpace(right) == "" {
		return left
	}
	if strings.TrimSpace(left) == "" {
		return right
	}
	leftTime, leftErr := time.Parse(time.RFC3339, left)
	rightTime, rightErr := time.Parse(time.RFC3339, right)
	if leftErr != nil || rightErr != nil {
		return right
	}
	if rightTime.After(leftTime) {
		return right
	}
	return left
}

func ageState(now time.Time, at string, maxHours int) string {
	observed, err := time.Parse(time.RFC3339, strings.TrimSpace(at))
	if err != nil || at == "" {
		return "missing"
	}
	if now.Sub(observed) > time.Duration(maxHours)*time.Hour {
		return "stale"
	}
	return "fresh"
}
