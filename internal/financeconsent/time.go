package financeconsent

import "time"

func isExpired(value string) bool {
	if value == "" {
		return false
	}
	parsed, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return true
	}
	return parsed.Before(time.Now().UTC())
}
