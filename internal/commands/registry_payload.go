package commands

import (
	"encoding/json"
	"fmt"
	"strings"
)

func normalizeName(name string) string {
	return strings.ReplaceAll(strings.TrimSpace(strings.ToLower(name)), "-", "_")
}

func decodePayload(payload []byte, target any) error {
	if len(strings.TrimSpace(string(payload))) == 0 {
		payload = []byte("{}")
	}
	if err := json.Unmarshal(payload, target); err != nil {
		return fmt.Errorf("%w: %v", errInvalidPayload, err)
	}
	return nil
}
