package repogovernance

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func LoadManifest(root string) (Manifest, error) {
	var manifest Manifest
	if err := decode(filepath.Join(root, "traceability.json"), &manifest); err != nil {
		return manifest, err
	}
	return manifest, VerifyManifest(root, manifest)
}

func VerifyManifest(root string, manifest Manifest) error {
	if manifest.SchemaVersion != "repo.traceability/v1" || len(manifest.Groups) == 0 {
		return errors.New("traceability header is invalid")
	}
	for _, group := range manifest.Groups {
		if group.ID == "" || group.ChangePolicy != "bidirectional" || len(group.DocumentSources) == 0 || len(group.GeneratedDocuments) == 0 || len(group.Code) == 0 || len(group.Tests) == 0 {
			return fmt.Errorf("trace group %q is incomplete", group.ID)
		}
		paths := append(append(append([]string{}, group.DocumentSources...), group.GeneratedDocuments...), append(group.Code, group.Tests...)...)
		for _, rel := range paths {
			if filepath.IsAbs(rel) || filepath.Clean(rel) != rel || strings.HasPrefix(rel, "..") {
				return fmt.Errorf("unsafe trace path %q", rel)
			}
			if _, err := os.Stat(filepath.Join(root, rel)); err != nil {
				return fmt.Errorf("missing trace path %q", rel)
			}
		}
	}
	return nil
}
