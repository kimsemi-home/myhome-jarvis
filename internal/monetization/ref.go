package monetization

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

func validateRef(policy Policy, ref string) error {
	ref = filepath.ToSlash(strings.TrimSpace(ref))
	if ref == "" ||
		filepath.IsAbs(filepath.FromSlash(ref)) ||
		strings.Contains(ref, "..") {
		return fmt.Errorf("monetization evidence ref is not public-safe: %q", ref)
	}
	for _, prefix := range policy.AllowedEvidencePrefixes {
		if strings.HasPrefix(ref, prefix) {
			return nil
		}
	}
	return fmt.Errorf("monetization evidence ref %q is not allowed", ref)
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
