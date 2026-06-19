package commands

import (
	"fmt"
	"strings"
)

func openOTTPlan(command string, payload []byte) (Plan, error) {
	var body struct {
		Service string `json:"service"`
	}
	if err := decodePayload(payload, &body); err != nil {
		return Plan{}, err
	}
	service := strings.ToLower(strings.TrimSpace(body.Service))
	target, ok := ottURLs()[service]
	if !ok {
		return Plan{}, fmt.Errorf("%w: unsupported ott service %q", errInvalidPayload, body.Service)
	}
	return openURLPlan(command, target), nil
}

func openURLPayloadPlan(command string, payload []byte) (Plan, error) {
	var body struct {
		URL string `json:"url"`
	}
	if err := decodePayload(payload, &body); err != nil {
		return Plan{}, err
	}
	target, err := validateHTTPURL(body.URL)
	if err != nil {
		return Plan{}, err
	}
	return openURLPlan(command, target), nil
}
