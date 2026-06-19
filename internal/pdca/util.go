package pdca

import "fmt"

func requireAll(label string, have []string, want []string) error {
	for _, item := range want {
		if !contains(have, item) {
			return fmt.Errorf("PDCA %s %q is missing", label, item)
		}
	}
	return nil
}

func contains(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

func later(a string, b string) string {
	if b > a {
		return b
	}
	return a
}
