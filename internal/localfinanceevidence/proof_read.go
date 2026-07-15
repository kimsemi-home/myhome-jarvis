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
			return errors.New("local finance execution proof artifact hash changed")
		}
		if err := validateProofBody(body, month, ref); err != nil {
			return err
		}
	}
	return nil
}

func validateProofBody(body []byte, month string, ref ProofRef) error {
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.DisallowUnknownFields()
	var validateErr error
	switch ref.ProofSchema {
	case creditProofSchema:
		var report CreditReport
		if err := decoder.Decode(&report); err != nil {
			return err
		}
		validateErr = validateCreditReport(report, month, ref)
	case portfolioProofSchema:
		var report PortfolioReport
		if err := decoder.Decode(&report); err != nil {
			return err
		}
		validateErr = validatePortfolioReport(report, month, ref)
	case revenueProofSchema:
		var report RevenueReport
		if err := decoder.Decode(&report); err != nil {
			return err
		}
		validateErr = validateRevenueReport(report, month, ref)
	case operatorProofSchema:
		var report OperatorReport
		if err := decoder.Decode(&report); err != nil {
			return err
		}
		validateErr = validateOperatorReport(report, month, ref)
	default:
		return errors.New("local finance execution proof schema is unsupported")
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return errors.New("local finance execution proof contains extra JSON")
	}
	return validateErr
}
