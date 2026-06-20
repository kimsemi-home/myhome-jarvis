package contextpack

import (
	"fmt"
	"strings"
)

func requireAll(label string, values []string, required []string) error {
	for _, item := range required {
		if !contains(values, item) {
			return fmt.Errorf("context pack %s %q is missing", label, item)
		}
	}
	return nil
}

func contains(values []string, value string) bool {
	for _, item := range values {
		if item == value {
			return true
		}
	}
	return false
}

func containsUnsafeText(value string) bool {
	value = strings.ToLower(value)
	for _, marker := range unsafeMarkers() {
		if strings.Contains(value, marker) {
			return true
		}
	}
	return false
}
