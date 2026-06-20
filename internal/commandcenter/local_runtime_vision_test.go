package commandcenter

import "testing"

func TestVisionAuditIncludesLocalRuntimeForSelfImprovement(t *testing.T) {
	policy, err := readVisionPolicy(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	status := visionAuditFixtureStatus(policy)
	status.NextSafeAction = "repair_local_runtime_health"
	status.LocalRuntime = LocalRuntimeSummary{HealthDebtCount: 2}
	status.BlockedGates = append(status.BlockedGates,
		GateSummary{Key: "local_runtime"},
	)
	status.BlockedGateCount = len(status.BlockedGates)
	audit := visionAuditFromStatus(policy, status)
	row := visionAuditRowByKey(audit, "self_improvement_loop")
	if row.State != "blocked" {
		t.Fatalf("self-improvement runtime state = %#v", row)
	}
	if !containsString(row.EvidenceRefs, "local_runtime") ||
		!containsString(row.GateRefs, "local_runtime") {
		t.Fatalf("self-improvement runtime refs = %#v", row)
	}
	if audit.GoalComplete || audit.NextSafeAction != "repair_local_runtime_health" {
		t.Fatalf("vision runtime next action = %#v", audit)
	}
}
