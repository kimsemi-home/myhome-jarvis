package localfinanceevidence

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
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
			return errors.New("credit execution proof artifact hash changed")
		}
		decoder := json.NewDecoder(bytes.NewReader(body))
		decoder.DisallowUnknownFields()
		var report CreditReport
		if err := decoder.Decode(&report); err != nil {
			return err
		}
		if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
			return errors.New("credit execution proof contains extra JSON")
		}
		if err := validateCreditReport(report, month, ref); err != nil {
			return err
		}
	}
	return nil
}
