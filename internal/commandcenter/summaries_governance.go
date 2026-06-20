package commandcenter

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/authority"
	"github.com/kimsemi-home/myhome-jarvis/internal/codexcost"
	"github.com/kimsemi-home/myhome-jarvis/internal/monetization"
	"github.com/kimsemi-home/myhome-jarvis/internal/review"
)

func summarizeAuthority(status authority.Status) AuthoritySummary {
	return AuthoritySummary{
		Outcome:                       status.Outcome,
		ActiveRule:                    status.ActiveRule,
		BlockedDecisionCount:          status.BlockedDecisionCount,
		AuthorityDebtCount:            status.AuthorityDebtCount,
		PublicRepoMode:                status.PublicRepoMode,
		PublicSafetyOK:                status.PublicSafetyOK,
		SelfAuthorityAllowed:          status.SelfAuthorityAllowed,
		SelfApprovalBlockedProfiles:   status.SelfApprovalBlockedProfileCount,
		ReviewRequiredProfileCount:    status.ReviewRequiredProfileCount,
		PublicSafetyGatedProfileCount: status.PublicSafetyGatedProfileCount,
	}
}

func summarizeReview(status review.Status) ReviewSummary {
	return ReviewSummary{
		CapacityState:     status.CapacityState,
		ActiveRule:        status.ActiveRule,
		OpenCount:         status.OpenCount,
		HighRiskOpenCount: status.HighRiskOpenCount,
		ReviewDebtCount:   status.ReviewDebtCount,
	}
}

func summarizeCost(status codexcost.Status) CostSummary {
	return CostSummary{
		BudgetState:          status.BudgetState,
		TotalUnits:           status.TotalUnits,
		WarningUnitThreshold: status.WarningUnitThreshold,
		ReviewUnitThreshold:  status.ReviewUnitThreshold,
		ReviewRequiredCount:  status.ReviewRequiredCount,
		MissingEvidenceCount: status.MissingEvidenceCount,
	}
}

func summarizeMonetization(status monetization.Status) MonetizationSummary {
	return MonetizationSummary{
		ExperimentCount:           status.ExperimentCount,
		DecisionCount:             status.DecisionCount,
		ReviewRequiredCount:       status.ReviewRequiredCount,
		MonetizationDebtCount:     status.MonetizationDebtCount,
		MissingEvidenceCount:      status.MissingEvidenceCount,
		MissingCostEstimateCount:  status.MissingCostEstimateCount,
		ExpectedValueUnknownCount: status.ExpectedValueUnknownCount,
	}
}
