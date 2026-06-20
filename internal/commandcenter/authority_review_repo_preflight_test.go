package commandcenter

import "testing"

func TestAuthorityReviewSurfacesRepoFactoryPreflight(t *testing.T) {
	status, err := StatusForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	assertRepoPreflight(t, status.RepoFactoryPreflight)

	brief, err := AuthorityReviewBriefForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	assertRepoPreflight(t, brief.RepoFactoryPreflight)

	packet, err := AuthorityReviewDecisionPacketForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	assertRepoPreflight(t, packet.RepoFactoryPreflight)
	if packet.CanApplyDecision || packet.ApprovalBoundary.ApprovalGranted {
		t.Fatalf("packet grants authority = %#v", packet)
	}
}

func assertRepoPreflight(t *testing.T, summary RepoFactoryPreflightSummary) {
	t.Helper()
	if !summary.PublicSafe || summary.CreationAllowed ||
		summary.SelfApprovalAllowed {
		t.Fatalf("repo factory preflight authority = %#v", summary)
	}
	if summary.CreationDecision != "blocked_pending_review_evidence" ||
		summary.BlockingGateCount != 1 ||
		summary.NextSafeAction != "await_human_authority_review" {
		t.Fatalf("repo factory preflight state = %#v", summary)
	}
	if !containsString(summary.MissingEvidenceKeys, "authority_review") ||
		containsString(summary.MissingEvidenceKeys, "public_safety_evidence") {
		t.Fatalf("repo factory preflight evidence = %#v", summary)
	}
}
