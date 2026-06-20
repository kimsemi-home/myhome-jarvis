package commandcenter

import "testing"

func TestAuthorityReviewBriefIncludesCapabilityReadiness(t *testing.T) {
	policy, err := readVisionPolicy(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	status := authorityReviewBriefStatus(policy)
	brief := authorityReviewBriefFromStatus(authorityReviewBriefPlan(), status)
	assertAuthorityReviewCapabilityReadiness(t, brief.CapabilityReadiness)
}

func TestAuthorityReviewDecisionPacketIncludesCapabilityReadiness(t *testing.T) {
	policy, err := readVisionPolicy(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	status := authorityReviewBriefStatus(policy)
	brief := authorityReviewBriefFromStatus(authorityReviewBriefPlan(), status)
	packet := authorityReviewDecisionPacketFromStatus(brief, status)
	assertAuthorityReviewCapabilityReadiness(t, packet.CapabilityReadiness)
	if packet.CanApplyDecision || packet.ApprovalBoundary.ApprovalGranted {
		t.Fatalf("packet grants authority = %#v", packet)
	}
}

func TestAuthorityReviewCapabilityReadinessMustBePublicSafe(t *testing.T) {
	policy, err := readVisionPolicy(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	status := authorityReviewBriefStatus(policy)
	status.MediaReadiness.PublicSafe = false
	brief := authorityReviewBriefFromStatus(authorityReviewBriefPlan(), status)
	if brief.PublicSafe {
		t.Fatalf("brief allowed unsafe capability readiness: %#v", brief.CapabilityReadiness)
	}
	packet := authorityReviewDecisionPacketFromStatus(brief, status)
	if packet.PublicSafe {
		t.Fatalf("packet allowed unsafe capability readiness: %#v", packet.CapabilityReadiness)
	}
}

func assertAuthorityReviewCapabilityReadiness(
	t *testing.T,
	summary CapabilityReadinessSummary,
) {
	t.Helper()
	if !summary.PublicSafe || summary.CapabilityCount != 6 {
		t.Fatalf("capability summary = %#v", summary)
	}
	if summary.Media.State != "ready" || !summary.Media.PlaybackReady {
		t.Fatalf("media readiness = %#v", summary.Media)
	}
	if summary.FinanceConsent.State != "ready" ||
		summary.FinanceConsent.ReadinessState != "ready_read_only" ||
		summary.FinanceConsent.ForbiddenActionEnabledCount != 0 {
		t.Fatalf("finance readiness = %#v", summary.FinanceConsent)
	}
	if summary.Monetization.State != "ready" ||
		summary.Monetization.ExperimentCount == 0 {
		t.Fatalf("monetization readiness = %#v", summary.Monetization)
	}
	if summary.CodexCost.State != "ready" ||
		!summary.CodexCost.PublicSafe ||
		summary.CodexCost.BriefDecision != "allow" ||
		summary.CodexCost.CanApplyExpansion {
		t.Fatalf("codex cost readiness = %#v", summary.CodexCost)
	}
}
