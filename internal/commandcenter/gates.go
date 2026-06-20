package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/contextpack"

func blockedGates(in inputs) []GateSummary {
	gates := []GateSummary{}
	add := func(cond bool, key string, label string, reason string, count int) {
		if cond {
			gates = append(gates, GateSummary{Key: key, Label: label,
				State: "blocked", Reason: reason, Count: count})
		}
	}
	evidenceDebt := in.Evidence.DanglingEvidenceRefCount + in.Evidence.OpenLearningCount
	incidentDebt := in.Incidents.IncidentDebtCount + in.Incidents.StaleQuarantineCount
	reviewDebt := in.Review.ReviewDebtCount + in.Review.HighRiskOpenCount
	financeConsentDebt := in.FinanceConsent.ConsentDebtCount
	costDebt := in.Cost.ReviewRequiredCount + in.Cost.MissingEvidenceCount
	codexSustainabilityDebt := in.CodexSustainability.ReviewGateCount
	contextPackDebt := contextPackDebtCount(in.ContextPack)
	monetizationDebt := in.Monetization.MonetizationDebtCount
	repoFactoryDebt := in.RepoFactory.MissingTemplateRoleCount +
		in.RepoFactory.MissingCreationGateCount +
		in.RepoFactory.ForbiddenTemplateValueCount
	if in.RepoFactory.RepoCreationBlockedUntilReview && repoFactoryDebt == 0 {
		repoFactoryDebt = 1
	}
	add(!in.PDCA.Ready, "pdca", "PDCA", "pdca artifacts or cycles are not ready",
		in.PDCA.MissingArtifactCount+in.PDCA.InvalidCycleCount)
	add(evidenceDebt > 0, "evidence", "Evidence",
		"evidence links or learning observations need closure", evidenceDebt)
	add(incidentDebt > 0, "incidents", "Incidents",
		"incident or quarantine debt is open", incidentDebt)
	add(authorityBlocked(in.Authority.Outcome), "authority", "Authority",
		in.Authority.ActiveRule, in.Authority.BlockedDecisionCount)
	add(reviewDebt > 0 || in.Review.CapacityState == "overloaded", "review", "Review",
		in.Review.ActiveRule, reviewDebt)
	add(in.FinanceConsent.ReadinessState != "ready_read_only", "finance_consent",
		"Finance Consent", "real finance connectors and shared scopes need consent",
		financeConsentDebt)
	add(in.Cost.BudgetState != "ok" || costDebt > 0, "cost", "Codex Cost",
		"cost budget or evidence review is required", costDebt)
	add(in.CodexSustainability.SustainabilityPosture != "sustainable",
		"codex_sustainability", "Codex Sustainability",
		"usage growth, trend freshness, or optimization evidence needs review",
		codexSustainabilityDebt)
	add(!in.ContextPack.PublicSafe, "context_pack", "Context Pack",
		"cross-repo context, ontology, or authority handoff is incomplete",
		contextPackDebt)
	add(monetizationDebt > 0, "monetization", "Monetization",
		"experiment decisions need evidence, cost, or review closure", monetizationDebt)
	add(!in.RepoFactory.PublicSafe || in.RepoFactory.RepoCreationBlockedUntilReview,
		"repo_factory", "Repo Factory",
		"repo creation requires authority review and public safety evidence",
		repoFactoryDebt)
	return gates
}

func contextPackDebtCount(status contextpack.Status) int {
	count := 0
	if status.SplitCriteriaCount < 5 {
		count++
	}
	if status.ExportedArtifactCount < 6 {
		count++
	}
	if !status.PublicSafe {
		count++
	}
	return count
}

func authorityBlocked(outcome string) bool {
	return outcome != "" && outcome != "allowed" && outcome != "ok"
}
