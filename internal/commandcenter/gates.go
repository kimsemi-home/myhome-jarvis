package commandcenter

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
	costDebt := in.Cost.ReviewRequiredCount + in.Cost.MissingEvidenceCount
	monetizationDebt := in.Monetization.MonetizationDebtCount
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
	add(in.Cost.BudgetState != "ok" || costDebt > 0, "cost", "Codex Cost",
		"cost budget or evidence review is required", costDebt)
	add(monetizationDebt > 0, "monetization", "Monetization",
		"experiment decisions need evidence, cost, or review closure", monetizationDebt)
	return gates
}

func authorityBlocked(outcome string) bool {
	return outcome != "" && outcome != "allowed" && outcome != "ok"
}
