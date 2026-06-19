package knowledge

import "fmt"

func validateContexts(registry Registry, contexts map[string]bool) []string {
	var failures []string
	for _, context := range registry.BoundedContexts {
		name := clean(context.Name)
		if name == "" {
			failures = append(failures, "bounded context name is required")
			continue
		}
		if contexts[name] {
			failures = append(failures, fmt.Sprintf("duplicate bounded context %q", name))
		}
		contexts[name] = true
	}
	return failures
}
