package commandcenter

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestAuthorityReviewBriefSummarizesGatedHandoff(t *testing.T) {
	policy, err := readVisionPolicy(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	status := authorityReviewBriefStatus(policy)
	brief := authorityReviewBriefFromStatus(authorityReviewBriefPlan(), status)
	if !brief.PublicSafe || brief.Context != "AuthorityReviewBrief" {
		t.Fatalf("brief status = %#v", brief)
	}
	if brief.RequestID == "" || brief.EvidenceRef == "" ||
		brief.QueueState != "pending_human_review" {
		t.Fatalf("brief refs = %#v", brief)
	}
	if !containsString(brief.GatedCapabilityKeys, "shorts_factory_control_plane") ||
		!containsString(brief.BlockedGateKeys, "repo_factory") {
		t.Fatalf("brief gates = %#v", brief)
	}
	if brief.RequiredReviewClassCount != 5 ||
		!containsString(brief.RequiredReviewClasses, "public_repo_change_review") {
		t.Fatalf("review classes = %#v", brief)
	}
	if brief.ReviewRequestStaleAfterHours != 24 ||
		brief.ReviewEscalationAction == "" {
		t.Fatalf("review stale guard = %#v", brief)
	}
	if !brief.RepoFactoryGate.RepoCreationBlockedUntilReview ||
		brief.ApprovalBoundary.ApprovalGranted ||
		brief.ApprovalBoundary.ExternalWritesAllowed ||
		brief.ApprovalBoundary.SelfApprovalAllowed {
		t.Fatalf("approval boundary = %#v", brief)
	}
	if brief.RepoFactoryPreflight.CreationAllowed ||
		brief.RepoFactoryPreflight.BlockingGateCount != 1 ||
		!containsString(brief.RepoFactoryPreflight.MissingEvidenceKeys, "authority_review") ||
		containsString(brief.RepoFactoryPreflight.MissingEvidenceKeys, "public_safety_evidence") {
		t.Fatalf("repo factory preflight = %#v", brief.RepoFactoryPreflight)
	}
	assertReadyExternalEvidenceSummary(t, brief.ExternalEvidence)
}

func TestAuthorityReviewBriefDoesNotExposePrivatePayloadFields(t *testing.T) {
	brief, err := AuthorityReviewBriefForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(brief)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range authorityReviewBriefForbiddenFields() {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("authority review brief leaked %q in %s", forbidden, body)
		}
	}
}
