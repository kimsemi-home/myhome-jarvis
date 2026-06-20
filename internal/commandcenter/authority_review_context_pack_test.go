package commandcenter

import "testing"

func TestAuthorityReviewBriefIncludesContextPackPosture(t *testing.T) {
	policy, err := readVisionPolicy(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	status := authorityReviewBriefStatus(policy)
	brief := authorityReviewBriefFromStatus(authorityReviewBriefPlan(), status)
	assertAuthorityReviewContextPack(t, brief.ContextPack)
}

func TestAuthorityReviewDecisionPacketIncludesContextPackPosture(t *testing.T) {
	policy, err := readVisionPolicy(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	status := authorityReviewBriefStatus(policy)
	brief := authorityReviewBriefFromStatus(authorityReviewBriefPlan(), status)
	packet := authorityReviewDecisionPacketFromStatus(brief, status)
	assertAuthorityReviewContextPack(t, packet.ContextPack)
}

func TestAuthorityReviewContextPackMustBePublicSafe(t *testing.T) {
	policy, err := readVisionPolicy(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	status := authorityReviewBriefStatus(policy)
	status.ContextPack.PublicSafe = false
	brief := authorityReviewBriefFromStatus(authorityReviewBriefPlan(), status)
	if brief.PublicSafe {
		t.Fatalf("brief allowed non-public-safe context pack: %#v", brief.ContextPack)
	}
	packet := authorityReviewDecisionPacketFromStatus(brief, status)
	if packet.PublicSafe {
		t.Fatalf("packet allowed non-public-safe context pack: %#v", packet.ContextPack)
	}
}

func assertAuthorityReviewContextPack(t *testing.T, summary ContextPackSummary) {
	t.Helper()
	if !summary.PublicSafe ||
		summary.PackID != "myhome-jarvis/context-pack" ||
		summary.Version != "v1" ||
		summary.UpstreamCompatibilityVersion != "myhome-jarvis/context-pack/v1" ||
		summary.OntologyVersion != "concept-registry/v1" {
		t.Fatalf("context pack identity = %#v", summary)
	}
	if summary.SplitCriteriaCount != 5 ||
		summary.ExportedArtifactCount != 6 ||
		summary.AuthorityContractVersion != "authority/v1" ||
		summary.SecurityContractVersion != "security/v1" ||
		summary.VerificationProfile != "quality" ||
		summary.VerificationRequiredUnitCount != 5 {
		t.Fatalf("context pack contract = %#v", summary)
	}
}
