package commandcenter

import "testing"

func TestAuthorityReviewBriefIncludesMergeEvidencePosture(t *testing.T) {
	policy, err := readVisionPolicy(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	status := authorityReviewBriefStatus(policy)
	brief := authorityReviewBriefFromStatus(authorityReviewBriefPlan(), status)
	assertAuthorityReviewMergeEvidence(t, brief.MergeEvidence)
}

func TestAuthorityReviewDecisionPacketIncludesMergeEvidencePosture(t *testing.T) {
	policy, err := readVisionPolicy(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	status := authorityReviewBriefStatus(policy)
	brief := authorityReviewBriefFromStatus(authorityReviewBriefPlan(), status)
	packet := authorityReviewDecisionPacketFromStatus(brief, status)
	assertAuthorityReviewMergeEvidence(t, packet.MergeEvidence)
}

func TestAuthorityReviewMergeEvidenceMustBePublicSafe(t *testing.T) {
	policy, err := readVisionPolicy(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	status := authorityReviewBriefStatus(policy)
	status.MergeEvidence.PublicSafe = false
	brief := authorityReviewBriefFromStatus(authorityReviewBriefPlan(), status)
	if brief.PublicSafe {
		t.Fatalf("brief allowed non-public-safe merge evidence: %#v", brief.MergeEvidence)
	}
	packet := authorityReviewDecisionPacketFromStatus(brief, status)
	if packet.PublicSafe {
		t.Fatalf("packet allowed non-public-safe merge evidence: %#v", packet.MergeEvidence)
	}
}

func assertAuthorityReviewMergeEvidence(t *testing.T, summary MergeEvidenceSummary) {
	t.Helper()
	if !summary.PublicSafe ||
		summary.DefaultBehavior != "merge_when_eligible" ||
		summary.MergePreference != "merge_after_checks_pass" {
		t.Fatalf("merge evidence posture = %#v", summary)
	}
	if !summary.PostMergeEvidenceRequired ||
		!summary.LinearCompletionRequired ||
		!summary.MainQualityRunRequired ||
		!summary.PrivateDataScanRequired {
		t.Fatalf("merge evidence requirements = %#v", summary)
	}
	if summary.MissingGateCount != 0 ||
		summary.MissingRequiredEvidenceCount != 0 ||
		!summary.MergeReady ||
		summary.MergeBlockedUntilEvidence {
		t.Fatalf("merge evidence gates = %#v", summary)
	}
}
