package financeconsent

import (
	"path/filepath"
	"sort"
	"strings"
)

func evidenceRefsAllowed(refs []string, policy Policy) bool {
	for _, ref := range refs {
		if !evidenceRefAllowed(ref, policy.AllowedEvidencePrefixes) {
			return false
		}
	}
	return true
}

func evidenceRefAllowed(ref string, prefixes []string) bool {
	ref = filepath.ToSlash(strings.TrimSpace(ref))
	if ref == "" ||
		filepath.IsAbs(filepath.FromSlash(ref)) ||
		strings.Contains(ref, "..") {
		return false
	}
	for _, prefix := range prefixes {
		if strings.HasPrefix(ref, prefix) {
			return true
		}
	}
	return false
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
