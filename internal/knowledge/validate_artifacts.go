package knowledge

import "fmt"

func validateArtifacts(root string, registry Registry) []string {
	var failures []string
	for _, contract := range registry.GeneratedArtifactContracts {
		if err := requirePublicTarget(root, contract.Path); err != nil {
			failures = append(failures, fmt.Sprintf("artifact contract %q target %q: %v", contract.Name, contract.Path, err))
		}
	}
	return failures
}
