package financeconsent

import "strings"

func evidenceRefsAllowed(refs []string, policy Policy) bool {
	for _, ref := range refs {
		if !evidenceRefAllowed(ref, policy.AllowedEvidencePrefixes) {
			return false
		}
	}
	return true
}

func evidenceRefAllowed(ref string, prefixes []string) bool {
	if ref == "" || strings.HasPrefix(ref, "/") || strings.Contains(ref, "..") {
		return false
	}
	for _, prefix := range prefixes {
		if strings.HasPrefix(ref, prefix) {
			return true
		}
	}
	return false
}
