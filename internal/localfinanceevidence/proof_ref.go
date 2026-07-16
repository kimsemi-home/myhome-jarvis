package localfinanceevidence

import (
	"errors"
	"path/filepath"
	"strings"
)

const creditProofSchema = "myhome.ledger-credit-collection-rehearsal/v5"
const portfolioProofSchema = "myhome.portfolio-readonly-collection-rehearsal/v2"
const revenueProofSchema = "myhome.revenue-youtube-analytics-rehearsal/v2"
const operatorProofSchema = "myhome.finance-operator-monthly-rehearsal/v1"
const shortsProofSchema = "shorts.youtube-loopback-fixture-report/v2"
const shortsActivationProofSchema = "shorts.youtube-activation-boundary-rehearsal/v2"

var requiredProofs = map[string]struct {
	capability string
	schema     string
}{
	"ledger":             {"credit-collection-rehearsal", creditProofSchema},
	"ledger-batch-apply": {"credit-batch-apply-rehearsal", creditBatchApplyProofSchema},
	"portfolio":          {"readonly-collection-rehearsal", portfolioProofSchema},
	"revenue":            {"youtube-revenue-collection-rehearsal", revenueProofSchema},
	"finance-operator":   {"monthly-orchestration-rehearsal", operatorProofSchema},
	"shorts":             {"youtube-private-upload-rehearsal", shortsProofSchema},
	"shorts-activation":  {"youtube-activation-boundary-rehearsal", shortsActivationProofSchema},
}

func validateProofRefs(refs []ProofRef) error {
	if len(refs) != len(requiredProofs) {
		return errors.New("Ledger, Portfolio, Revenue, Finance Operator, Shorts upload, and Shorts activation proofs are required")
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
