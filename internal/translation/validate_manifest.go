package translation

import (
	"fmt"
	"strings"
)

func validateManifest(policy Policy, current manifest) error {
	if !contains(policy.AllowedContexts, strings.TrimSpace(current.SourceContext)) {
		return fmt.Errorf("source context is not allowed")
	}
	if !contains(policy.AllowedContexts, strings.TrimSpace(current.TargetContext)) {
		return fmt.Errorf("target context is not allowed")
	}
	if strings.TrimSpace(current.SourceVersion) == "" ||
		strings.TrimSpace(current.TargetVersion) == "" {
		return fmt.Errorf("translation manifest versions are required")
	}
	if len(current.PreservedRules) == 0 ||
		len(current.KnownLosses) == 0 ||
		strings.TrimSpace(current.Owner) == "" ||
		len(current.EvidenceRefs) == 0 {
		return fmt.Errorf("translation manifest required fields are incomplete")
	}
	for _, ref := range current.EvidenceRefs {
		if err := validateEvidenceRef(policy, ref); err != nil {
			return err
		}
	}
	return nil
}
