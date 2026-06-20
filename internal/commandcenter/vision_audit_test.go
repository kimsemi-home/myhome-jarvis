package commandcenter

import "testing"

func TestVisionAuditShowsIncompleteGoalWhenCapabilityGated(t *testing.T) {
	policy, err := readVisionPolicy(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	audit := visionAuditFromStatus(policy, visionAuditFixtureStatus(policy))
	if audit.GoalComplete {
		t.Fatalf("goal should still be incomplete: %#v", audit)
	}
	if audit.RequirementCount != 6 ||
		audit.ReadyRequirementCount != 4 ||
		audit.GatedRequirementCount != 2 {
		t.Fatalf("requirement counts = %#v", audit)
	}
	if audit.NextSafeAction != "await_human_authority_review" {
		t.Fatalf("next action = %#v", audit)
	}
	if audit.EvidenceRetention.CompressionArchivePattern != "compress_then_archive" ||
		!audit.EvidenceRetention.NoiseBudgetReady ||
		audit.EvidenceRetention.ConfigEvidenceField != "evidence_noise_budget" ||
		audit.EvidenceRetention.ConfigEvidenceSHA256 == "" {
		t.Fatalf("evidence retention audit = %#v", audit.EvidenceRetention)
	}
	if audit.EvidenceRetention.MaxNoiseRatioPercent > 25 ||
		!audit.EvidenceRetention.BreachBlocksArchive ||
		!containsString(audit.EvidenceRetention.ConfigHashInputs, "evidence_noise_budget") {
		t.Fatalf("evidence retention noise config = %#v", audit.EvidenceRetention)
	}
	row := visionAuditRowByKey(audit, "self_improvement_loop")
	if row.State != "gated" || !containsString(row.GateRefs, "authority") ||
		!containsString(row.EvidenceRefs, "storage_archive") {
		t.Fatalf("self-improvement row = %#v", row)
	}
}

func TestVisionGoalCompleteRequiresAllReadyAndNoOpenGates(t *testing.T) {
	status := Status{
		PublicSafe:       true,
		NextSafeAction:   "continue_closed_loop_planning",
		BlockedGates:     nil,
		BlockedGateCount: 0,
		Vision: VisionSummary{
			CapabilityCount:  2,
			ReadyPillarCount: 2,
		},
		StorageArchive: readyStorageArchiveSummary(),
	}
	if !visionGoalComplete(status) {
		t.Fatalf("expected complete status = %#v", status)
	}
	status.BlockedGateCount = 1
	if visionGoalComplete(status) {
		t.Fatalf("open gate should keep goal incomplete")
	}
	status.BlockedGateCount = 0
	status.StorageArchive.NoiseBudgetReady = false
	if visionGoalComplete(status) {
		t.Fatalf("noise budget debt should keep goal incomplete")
	}
}
