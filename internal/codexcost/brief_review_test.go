package codexcost

import "testing"

func TestBriefRequiresReviewWhenArchiveGovernanceIsMissing(t *testing.T) {
	brief := buildBrief(ROISummary{
		PolicyPath:                 PolicyRelativePath,
		BudgetState:                "ok",
		TotalUnits:                 100,
		AttributionCoveragePercent: 100,
		SustainabilityPosture:      "sustainable",
		TrendPosture:               "on_trend",
		StorageArchivePattern:      "compress_then_archive",
		StorageArchiveReady:        true,
		NoiseBudgetReady:           false,
		ConfigIsEvidence:           true,
		CheckedAt:                  "2026-06-19T00:00:00Z",
	}, Status{
		WarningUnitThreshold: 100000,
		ReviewUnitThreshold:  500000,
	})
	if brief.Decision != "review_required" ||
		brief.NextSafeAction != "hold_paid_or_external_expansion" ||
		!contains(brief.Reasons, "noise_budget_not_ready") {
		t.Fatalf("brief = %#v", brief)
	}
}
