package evidencequality

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

var errMissingEvidenceRef = errors.New("evidence quality snapshot evidence ref is required")

func validateRef(policy Policy, ref string) error {
	ref = filepath.ToSlash(strings.TrimSpace(ref))
	if ref == "" {
		return errMissingEvidenceRef
	}
	if filepath.IsAbs(filepath.FromSlash(ref)) || strings.Contains(ref, "..") {
		return fmt.Errorf("evidence quality ref must be repo-relative")
	}
	for _, prefix := range policy.AllowedEvidencePrefixes {
		if strings.HasPrefix(ref, prefix) {
			return nil
		}
	}
	return fmt.Errorf("evidence quality ref is outside allowed prefixes")
}
