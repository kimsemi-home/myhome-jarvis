package knowledge

import "fmt"

func validateHarnesses(root string, registry Registry, contexts map[string]bool) []string {
	var failures []string
	if len(registry.HarnessCaseContracts) == 0 {
		failures = append(failures, "harness case contracts are required")
	}
	for _, harness := range registry.HarnessCaseContracts {
		failures = append(failures, validateHarness(root, harness, contexts)...)
	}
	return failures
}

func validateHarness(root string, harness HarnessCase, contexts map[string]bool) []string {
	var failures []string
	if clean(harness.Name) == "" {
		failures = append(failures, "harness case name is required")
	}
	if !contexts[harness.BoundedContext] {
		failures = append(failures, fmt.Sprintf("harness case %q references unknown bounded context %q", harness.Name, harness.BoundedContext))
	}
	if clean(harness.Command) == "" {
		failures = append(failures, fmt.Sprintf("harness case %q must declare command", harness.Name))
	}
	if err := requirePublicTarget(root, harness.EvidenceTarget); err != nil {
		failures = append(failures, fmt.Sprintf("harness case %q target %q: %v", harness.Name, harness.EvidenceTarget, err))
	}
	return failures
}
