package localfinanceevidence

import (
	"errors"
	"path/filepath"
	"strings"
)

const creditProofSchema = "myhome.ledger-credit-collection-rehearsal/v1"

func validateProofRefs(refs []ProofRef) error {
	if len(refs) != 1 {
		return errors.New("one Ledger credit execution proof is required")
	}
	ref := refs[0]
	clean := filepath.Clean(ref.Path)
	if ref.Component != "ledger" || ref.Capability != "credit-collection-rehearsal" ||
		ref.ProofSchema != creditProofSchema || filepath.IsAbs(ref.Path) || clean == ".." ||
		strings.HasPrefix(clean, ".."+string(filepath.Separator)) ||
		!hashPattern.MatchString(ref.ArtifactSHA256) || !hashPattern.MatchString(ref.ReportHash) {
		return errors.New("Ledger credit execution proof reference is invalid")
	}
	return nil
}
