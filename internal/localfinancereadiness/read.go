package localfinancereadiness

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func Read(path string) (Manifest, error) {
	file, err := os.Open(path)
	if err != nil {
		return Manifest{}, err
	}
	manifest, err := decodeOne[Manifest](file)
	file.Close()
	if err != nil {
		return Manifest{}, err
	}
	if err := Validate(manifest); err != nil {
		return Manifest{}, err
	}
	base := filepath.Dir(path)
	for _, ref := range manifest.Plans {
		if err := verifyPlan(base, ref); err != nil {
			return Manifest{}, fmt.Errorf("%s readiness plan: %w", ref.Component, err)
		}
	}
	return manifest, nil
}

func verifyPlan(base string, ref Ref) error {
	body, err := os.ReadFile(filepath.Join(base, filepath.FromSlash(ref.Path)))
	if err != nil {
		return err
	}
	if digest(body) != ref.ArtifactSHA256 {
		return fmt.Errorf("artifact SHA-256 mismatch")
	}
	plan, err := decodeOne[Plan](bytes.NewReader(body))
	if err != nil {
		return err
	}
	return validatePlan(plan, ref)
}
