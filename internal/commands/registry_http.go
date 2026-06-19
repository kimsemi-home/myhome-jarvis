package commands

import (
	"fmt"
	"net/url"
	"strings"
)

func openURLPlan(name string, target string) Plan {
	return Plan{
		Name:   name,
		DryRun: true,
		Invocations: []Invocation{
			{Label: name, Argv: []string{"open", target}, URL: target},
		},
	}
}

func validateHTTPURL(raw string) (string, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return "", fmt.Errorf("%w: url is required", errInvalidPayload)
	}
	parsed, err := url.Parse(trimmed)
	if err != nil {
		return "", fmt.Errorf("%w: %v", errInvalidPayload, err)
	}
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return "", fmt.Errorf("%w: only http and https URLs are allowed", errInvalidPayload)
	}
	if parsed.Host == "" {
		return "", fmt.Errorf("%w: URL host is required", errInvalidPayload)
	}
	return parsed.String(), nil
}
