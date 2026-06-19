package authority

import "testing"

func TestAssessKeepsHighRiskDecisionsBlockedInPublicRepoMode(t *testing.T) {
	status := Assess(testPolicy(), clearInputs())

	if status.Outcome != "limited" || status.ActiveRule != "public_repo_high_risk_block" {
		t.Fatalf("status = %#v", status)
	}
	if status.ReasoningTierGrantsApproval || status.SelfAuthorityAllowed {
		t.Fatalf("authority flags = %#v", status)
	}
	if !contains(status.AllowedDecisions, "read_status") || !contains(status.AllowedDecisions, "low_risk_fixture_change") {
		t.Fatalf("allowed decisions = %#v", status.AllowedDecisions)
	}
	for _, blocked := range []string{"major_ontology_change", "security_boundary_change", "production_change", "evidence_pruning", "quarantine_release", "high_risk_automation"} {
		if !contains(status.BlockedDecisions, blocked) {
			t.Fatalf("blocked decisions missing %q in %#v", blocked, status.BlockedDecisions)
		}
	}
}

func TestAssessBlocksOnPublicSafetyOrLowConfidence(t *testing.T) {
	inputs := clearInputs()
	inputs.PublicSafety.OK = false
	status := Assess(testPolicy(), inputs)
	if status.Outcome != "blocked" || status.ActiveRule != "public_safety_not_ok" {
		t.Fatalf("status = %#v", status)
	}

	inputs = clearInputs()
	inputs.Confidence.LevelCap = "low"
	status = Assess(testPolicy(), inputs)
	if status.Outcome != "blocked" || status.ActiveRule != "confidence_cap_low" {
		t.Fatalf("status = %#v", status)
	}
}
