package commandcenter

import "testing"

func TestAuthorityReviewBriefIncludesLocalRuntimeReadiness(t *testing.T) {
	policy, err := readVisionPolicy(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	status := authorityReviewBriefStatus(policy)
	brief := authorityReviewBriefFromStatus(authorityReviewBriefPlan(), status)
	assertAuthorityReviewRuntime(t, brief.LocalRuntime, "healthy", 0)
}

func TestAuthorityReviewDecisionPacketIncludesLocalRuntimeReadiness(t *testing.T) {
	policy, err := readVisionPolicy(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	status := authorityReviewBriefStatus(policy)
	brief := authorityReviewBriefFromStatus(authorityReviewBriefPlan(), status)
	packet := authorityReviewDecisionPacketFromStatus(brief, status)
	assertAuthorityReviewRuntime(t, packet.LocalRuntime, "healthy", 0)
}

func TestAuthorityReviewRuntimeCanShowStaleWithoutGranting(t *testing.T) {
	policy, err := readVisionPolicy(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	status := authorityReviewBriefStatus(policy)
	status.LocalRuntime = LocalRuntimeSummary{
		PublicSafe: true, EvidenceRef: "local_runtime:supervisor",
		State: "unhealthy", Stale: true, HealthDebtCount: 3,
		NextSafeAction: "repair_local_runtime_health",
	}
	brief := authorityReviewBriefFromStatus(authorityReviewBriefPlan(), status)
	if !brief.PublicSafe || brief.ApprovalBoundary.ApprovalGranted {
		t.Fatalf("brief grants authority = %#v", brief)
	}
	assertAuthorityReviewRuntime(t, brief.LocalRuntime, "unhealthy", 3)
}

func assertAuthorityReviewRuntime(
	t *testing.T,
	runtime LocalRuntimeSummary,
	state string,
	debt int,
) {
	t.Helper()
	if !runtime.PublicSafe || runtime.RawRuntimePublicAllowed {
		t.Fatalf("runtime public safety = %#v", runtime)
	}
	if runtime.EvidenceRef != "local_runtime:supervisor" ||
		runtime.State != state || runtime.HealthDebtCount != debt {
		t.Fatalf("runtime readiness = %#v", runtime)
	}
}
