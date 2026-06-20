package monetization

import "fmt"

func requireAll(label string, values []string, required []string) error {
	normalized := normalizeList(values)
	for _, value := range required {
		if !contains(normalized, value) {
			return fmt.Errorf("monetization %s %q is missing", label, value)
		}
	}
	return nil
}
