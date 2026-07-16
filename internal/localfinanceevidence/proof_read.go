package localfinanceevidence

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"path/filepath"
)

func validateProofFiles(root, month string, refs []ProofRef) error {
	for _, ref := range refs {
		body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(ref.Path)))
		if err != nil {
			return err
		}
		sum := sha256.Sum256(body)
		if hex.EncodeToString(sum[:]) != ref.ArtifactSHA256 {
			return errors.New("local finance execution proof artifact hash changed")
		}
		if err := validateProofBody(body, month, ref); err != nil {
			return err
		}
	}
	return nil
}
