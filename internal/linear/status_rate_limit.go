package linear

import (
	"net/http"
	"strconv"
	"strings"
)

func rateLimitRemaining(header http.Header) int {
	for _, name := range []string{
		"X-RateLimit-Remaining",
		"X-RateLimit-Requests-Remaining",
		"Linear-RateLimit-Remaining",
	} {
		value := strings.TrimSpace(headerValue(header, name))
		if value == "" {
			continue
		}
		parsed, err := strconv.Atoi(value)
		if err == nil {
			return parsed
		}
	}
	return 0
}

func headerValue(header http.Header, name string) string {
	if value := header.Get(name); value != "" {
		return value
	}
	for key, values := range header {
		if !strings.EqualFold(key, name) || len(values) == 0 {
			continue
		}
		return values[0]
	}
	return ""
}
