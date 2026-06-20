package authority

import (
	"fmt"
	"strings"
	"time"
)

func normalizeReviewRecordTime(value string, now time.Time) (string, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return now.UTC().Format(time.RFC3339), nil
	}
	parsed, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return "", fmt.Errorf("authority review record time must be RFC3339")
	}
	return parsed.UTC().Format(time.RFC3339), nil
}
