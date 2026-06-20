package commandcenter

import "testing"

func TestAuthorityReviewBriefIncludesCodexSustainabilityPosture(t *testing.T) {
	policy, err := readVisionPolicy(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	status := authorityReviewBriefStatus(policy)
	brief := authorityReviewBriefFromStatus(authorityReviewBriefPlan(), status)
	assertAuthorityReviewCodexSustainability(t, brief.CodexSustainability)
}

func TestAuthorityReviewDecisionPacketIncludesCodexSustainabilityPosture(t *testing.T) {
	policy, err := readVisionPolicy(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	status := authorityReviewBriefStatus(policy)
	brief := authorityReviewBriefFromStatus(authorityReviewBriefPlan(), status)
	packet := authorityReviewDecisionPacketFromStatus(brief, status)
	assertAuthorityReviewCodexSustainability(t, packet.CodexSustainability)
}

func TestAuthorityReviewCodexSustainabilityMustBePublicSafe(t *testing.T) {
	policy, err := readVisionPolicy(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	status := authorityReviewBriefStatus(policy)
	status.CodexSustainability.PublicSafe = false
	brief := authorityReviewBriefFromStatus(authorityReviewBriefPlan(), status)
	if brief.PublicSafe {
		t.Fatalf("brief allowed non-public-safe sustainability: %#v", brief.CodexSustainability)
	}
	packet := authorityReviewDecisionPacketFromStatus(brief, status)
	if packet.PublicSafe {
		t.Fatalf("packet allowed non-public-safe sustainability: %#v", packet.CodexSustainability)
	}
}

func assertAuthorityReviewCodexSustainability(
	t *testing.T,
	summary CodexSustainabilitySummary,
) {
	t.Helper()
	if !summary.PublicSafe ||
		summary.SustainabilityPosture != "sustainable" ||
		summary.TrendPosture != "on_trend" ||
		summary.EvidenceFreshness != "fresh" {
		t.Fatalf("codex sustainability posture = %#v", summary)
	}
	if summary.RecordCount == 0 ||
		summary.TrendBaselineCount == 0 ||
		summary.CacheSavingsUnits == 0 ||
		summary.LatestTrendBaselineVersion == "" {
		t.Fatalf("codex sustainability evidence = %#v", summary)
	}
	if summary.ReviewGateCount != 0 ||
		summary.ValidationFailureCount != 0 ||
		summary.HumanReviewDebtCount != 0 ||
		summary.ReworkCount != 0 {
		t.Fatalf("codex sustainability debt = %#v", summary)
	}
}
