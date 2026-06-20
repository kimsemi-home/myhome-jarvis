package authority

import (
	"sort"
	"time"
)

func Assess(policy Policy, inputs Inputs) Status {
	status := Status{
		PolicyPath:                  PolicyRelativePath,
		InputCount:                  len(policy.RequiredInputs),
		DecisionCount:               len(policy.Decisions),
		PublicRepoMode:              policy.PublicRepoHighRiskBlocked,
		ReasoningTierGrantsApproval: policy.ReasoningTierGrantsApproval,
		SelfAuthorityAllowed:        policy.SelfAuthorityAllowed,
		PublicSafetyOK:              inputs.PublicSafety.OK,
		ConfidenceCap:               normalizeToken(inputs.Confidence.LevelCap),
		EvidenceQualityDebtCount:    inputs.EvidenceQuality.ReassessmentDebtCount,
		IncidentDebtCount:           inputs.Incidents.IncidentDebtCount,
		ControlPlaneDebtCount:       inputs.ControlPlane.ManifestDebtCount,
		TranslationDebtCount:        inputs.Translation.OpenDebtCount + inputs.Translation.ForbiddenLossCount,
		HumanReviewDebtCount:        inputs.Review.ReviewDebtCount,
		HumanReviewCapacityState:    normalizeToken(inputs.Review.CapacityState),
		ByRisk:                      map[string]int{},
		CheckedAt:                   time.Now().UTC().Format(time.RFC3339),
	}
	defaultUnknown(&status.ConfidenceCap)
	defaultUnknown(&status.HumanReviewCapacityState)
	status.AuthorityDebtCount = status.EvidenceQualityDebtCount +
		status.IncidentDebtCount +
		status.ControlPlaneDebtCount +
		status.TranslationDebtCount +
		status.HumanReviewDebtCount
	status.Outcome, status.ActiveRule = authorityOutcome(status, inputs)
	for _, decision := range normalizedDecisions(policy.Decisions) {
		status.ByRisk[decision.Risk]++
		if decisionAllowed(status.Outcome, decision) {
			status.AllowedDecisions = append(status.AllowedDecisions, decision.Key)
		} else {
			status.BlockedDecisions = append(status.BlockedDecisions, decision.Key)
		}
	}
	sort.Strings(status.AllowedDecisions)
	sort.Strings(status.BlockedDecisions)
	status.AllowedDecisionCount = len(status.AllowedDecisions)
	status.BlockedDecisionCount = len(status.BlockedDecisions)
	applyProfiles(policy, &status)
	return status
}

func defaultUnknown(value *string) {
	if *value == "" {
		*value = "unknown"
	}
}
