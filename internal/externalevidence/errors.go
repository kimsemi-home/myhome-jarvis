package externalevidence

import "strings"

func errorCategory(err error) string {
	if err == nil {
		return ""
	}
	value := err.Error()
	for _, allowed := range []string{
		"request_build", "network", "http_status", "payload_limit",
		"payload_too_large", "read_body", "write_private_layer",
	} {
		if strings.Contains(value, allowed) {
			return allowed
		}
	}
	return "internal"
}
