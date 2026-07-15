package localfinanceevidence

import (
	"errors"
	"path/filepath"
	"strings"
)

const creditProofSchema = "myhome.ledger-credit-collection-rehearsal/v1"
const portfolioProofSchema = "myhome.portfolio-readonly-collection-rehearsal/v1"
const revenueProofSchema = "myhome.revenue-youtube-analytics-rehearsal/v1"

var requiredProofs = map[string]struct {
	capability string
	schema     string
}{
	"ledger":    {"credit-collection-rehearsal", creditProofSchema},
	"portfolio": {"readonly-collection-rehearsal", portfolioProofSchema},
	"revenue":   {"youtube-revenue-collection-rehearsal", revenueProofSchema},
}

func validateProofRefs(refs []ProofRef) error {
	if len(refs) != len(requiredProofs) {
		return errors.New("Ledger, Portfolio, and Revenue execution proofs are required")
	}
	seen := map[string]bool{}
	for _, ref := range refs {
		expected, ok := requiredProofs[ref.Component]
		clean := filepath.Clean(ref.Path)
		if !ok || seen[ref.Component] || ref.Capability != expected.capability || ref.ProofSchema != expected.schema ||
			filepath.IsAbs(ref.Path) || clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) ||
			!hashPattern.MatchString(ref.ArtifactSHA256) || !hashPattern.MatchString(ref.ReportHash) {
			return errors.New("local finance execution proof reference is invalid")
		}
		seen[ref.Component] = true
	}
	return nil
}
