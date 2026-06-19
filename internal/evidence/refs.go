package evidence

import (
	"fmt"
	"path/filepath"
	"strings"
)

func normalizeEvidenceRef(policy Policy, ref string) (string, error) {
	normalized := filepath.ToSlash(strings.TrimSpace(ref))
	if normalized == "" ||
		filepath.IsAbs(filepath.FromSlash(normalized)) ||
		strings.Contains(normalized, "..") {
		return "", fmt.Errorf("evidence graph found invalid evidence ref")
	}
	for _, prefix := range policy.AllowedEvidencePrefixes {
		if strings.HasPrefix(normalized, prefix) {
			return normalized, nil
		}
	}
	return "", fmt.Errorf("evidence graph found evidence ref outside allowed prefixes")
}
