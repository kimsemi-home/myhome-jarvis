package codexsustainability

import (
	"fmt"
	"strings"
)

func validateProposal(record Record) error {
	if record.ProposalID == "" || record.MonetizationRef == "" {
		return fmt.Errorf("codex sustainability proposal id and monetization ref are required")
	}
	if record.CostPerAcceptedChange <= 0 || record.MedianCycleMinutes <= 0 ||
		record.CacheSavingsUnits < 0 || record.DefectReworkRate < 0 ||
		record.DefectReworkRate > 1 {
		return fmt.Errorf("codex sustainability proposal evidence metrics are required")
	}
	if unsafePublicKey(record.ProposalID) || unsafePublicKey(record.MonetizationRef) {
		return fmt.Errorf("codex sustainability proposal public refs are invalid")
	}
	return nil
}

func unsafePublicKey(value string) bool {
	value = strings.ToLower(strings.TrimSpace(value))
	if value == "" || strings.Contains(value, "://") || strings.Contains(value, "/") {
		return true
	}
	for _, marker := range []string{
		"token", "secret", "credential", "cookie", "account_id", "card_number",
		"raw_prompt", "raw_transcript", "private", "linear.app/",
	} {
		if strings.Contains(value, marker) {
			return true
		}
	}
	return false
}

func applyProposal(status *Status, record Record) {
	status.FeatureProposalCount++
	if record.CostPerAcceptedChange > status.maxProposalCostPerAcceptedChange {
		status.maxProposalCostPerAcceptedChange = record.CostPerAcceptedChange
	}
	if record.MedianCycleMinutes > 0 {
		status.cycleMinutes = append(status.cycleMinutes, record.MedianCycleMinutes)
	}
	status.CacheSavingsUnits += record.CacheSavingsUnits
}
