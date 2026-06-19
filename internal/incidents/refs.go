package incidents

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

func validateRef(policy Policy, ref string) error {
	ref = filepath.ToSlash(strings.TrimSpace(ref))
	if ref == "" {
		return errMissingEvidence
	}
	if filepath.IsAbs(filepath.FromSlash(ref)) || strings.Contains(ref, "..") {
		return fmt.Errorf("incident evidence ref must be repo-relative")
	}
	for _, prefix := range policy.AllowedEvidencePrefixes {
		if strings.HasPrefix(ref, prefix) {
			return nil
		}
	}
	return fmt.Errorf("incident evidence ref %q is outside allowed prefixes", ref)
}

func normalizeRefs(values []string) []string {
	seen := map[string]bool{}
	refs := make([]string, 0, len(values))
	for _, value := range values {
		ref := filepath.ToSlash(strings.TrimSpace(value))
		if ref == "" || seen[ref] {
			continue
		}
		seen[ref] = true
		refs = append(refs, ref)
	}
	sort.Strings(refs)
	return refs
}
