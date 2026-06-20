package commandcenter

import "testing"

func TestStatusIncludesCodexCostScaling(t *testing.T) {
	status, err := StatusForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	scaling := status.CodexCostScaling
	if !scaling.PublicSafe ||
		scaling.Decision == "" ||
		scaling.NextSafeAction == "" ||
		scaling.Recommendation == "" {
		t.Fatalf("codex cost scaling = %#v", scaling)
	}
	if scaling.CanApplyExpansion ||
		scaling.GrantingScalingOptionCount != 0 ||
		scaling.ScalingOptionCount == 0 {
		t.Fatalf("scaling grants expansion = %#v", scaling)
	}
	if scaling.RemainingToWarningUnits <= 0 ||
		scaling.RemainingToReviewUnits <= scaling.RemainingToWarningUnits ||
		scaling.WarningUsedPercent < 0 ||
		scaling.ReviewUsedPercent < 0 {
		t.Fatalf("scaling headroom = %#v", scaling)
	}
	if !scaling.StorageArchiveReady ||
		!scaling.NoiseBudgetReady ||
		!scaling.ConfigIsEvidence ||
		scaling.ArchiveManifestBudgetBreaches != 0 ||
		scaling.ArchiveManifestInvalidEntries != 0 {
		t.Fatalf("scaling storage evidence = %#v", scaling)
	}
	if !hasPillar(scaling.RecommendedScalingOptionKeys, scaling.NextSafeAction) {
		t.Fatalf("scaling recommendations = %#v", scaling)
	}
}
