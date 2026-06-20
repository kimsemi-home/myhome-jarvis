package codexsustainability

import "fmt"

func validateProposal(record Record) error {
	if record.ProposalID == "" || record.MonetizationRef == "" {
		return fmt.Errorf("codex sustainability proposal id and monetization ref are required")
	}
	if record.CostPerAcceptedChange <= 0 || record.MedianCycleMinutes <= 0 ||
		record.CacheSavingsUnits < 0 || record.DefectReworkRate < 0 {
		return fmt.Errorf("codex sustainability proposal evidence metrics are required")
	}
	return nil
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
