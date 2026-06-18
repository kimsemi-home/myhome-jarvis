package translation

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func validateReferencedManifest(root string, policy Policy, manifestPath string) error {
	manifestPath = filepath.ToSlash(strings.TrimSpace(manifestPath))
	if manifestPath == "" {
		return fmt.Errorf("translation manifest path is required")
	}
	if filepath.IsAbs(filepath.FromSlash(manifestPath)) || strings.Contains(manifestPath, "..") {
		return fmt.Errorf("translation manifest path must be repo-relative")
	}
	prefix := strings.TrimRight(policy.PrivateManifestRoot, "/") + "/"
	if !strings.HasPrefix(manifestPath, prefix) {
		return fmt.Errorf("translation manifest path must stay under private manifest root")
	}
	if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(manifestPath))); err != nil {
		return err
	}
	return nil
}

func validateEvidenceRef(policy Policy, ref string) error {
	ref = filepath.ToSlash(strings.TrimSpace(ref))
	if ref == "" {
		return fmt.Errorf("translation evidence ref is required")
	}
	if filepath.IsAbs(filepath.FromSlash(ref)) || strings.Contains(ref, "..") {
		return fmt.Errorf("translation evidence ref must be repo-relative")
	}
	for _, prefix := range policy.AllowedEvidencePrefixes {
		if strings.HasPrefix(ref, prefix) {
			return nil
		}
	}
	return fmt.Errorf("translation evidence ref is outside allowed prefixes")
}
