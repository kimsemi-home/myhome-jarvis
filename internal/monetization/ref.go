package monetization

import (
	"fmt"
	"path/filepath"
	"strings"
)

func validateRef(policy Policy, ref string) error {
	ref = strings.TrimSpace(ref)
	if ref == "" || filepath.IsAbs(ref) || strings.Contains(ref, "..") {
		return fmt.Errorf("monetization evidence ref is not public-safe: %q", ref)
	}
	for _, prefix := range policy.AllowedEvidencePrefixes {
		if strings.HasPrefix(ref, prefix) {
			return nil
		}
	}
	return fmt.Errorf("monetization evidence ref %q is not allowed", ref)
}
