package commandcenter

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestAuthorityReviewDecisionPacketIncludesStorageEvidence(t *testing.T) {
	policy, err := readVisionPolicy(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	status := authorityReviewBriefStatus(policy)
	brief := authorityReviewBriefFromStatus(authorityReviewBriefPlan(), status)
	packet := authorityReviewDecisionPacketFromStatus(brief, status)
	if !packet.PublicSafe || packet.Context != "AuthorityReviewDecisionPacket" {
		t.Fatalf("packet status = %#v", packet)
	}
	if packet.DecisionPacketState != "review_only" || packet.CanApplyDecision {
		t.Fatalf("packet grants authority = %#v", packet)
	}
	if packet.StorageEvidence.CompressionArchivePattern != "compress_then_archive" ||
		packet.StorageEvidence.ConfigEvidenceField != "evidence_noise_budget" ||
		!packet.StorageEvidence.ConfigIsEvidence {
		t.Fatalf("storage evidence = %#v", packet.StorageEvidence)
	}
	if !containsString(packet.StorageEvidence.ConfigHashInputs, "evidence_noise_budget") ||
		!containsString(packet.StorageEvidence.DedupeKeyFields, "evidence_ref") {
		t.Fatalf("storage evidence config = %#v", packet.StorageEvidence)
	}
	if packet.RequiredReviewClassCount != 5 ||
		!containsString(packet.GatedCapabilityKeys, "shorts_factory_control_plane") {
		t.Fatalf("review handoff = %#v", packet)
	}
}

func TestAuthorityReviewDecisionPacketOptionsAreNonGranting(t *testing.T) {
	packet, err := AuthorityReviewDecisionPacketForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	if len(packet.DecisionOptions) != 3 {
		t.Fatalf("decision options = %#v", packet.DecisionOptions)
	}
	for _, option := range packet.DecisionOptions {
		if option.ThisPacketGrantsApproval || option.AllowsExternalWrites ||
			option.AllowsRepoCreation || option.AllowsWorkflowChanges ||
			option.AllowsSelfApproval {
			t.Fatalf("granting option = %#v", option)
		}
	}
}

func TestAuthorityReviewDecisionPacketDoesNotExposePrivatePayloadFields(t *testing.T) {
	packet, err := AuthorityReviewDecisionPacketForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(packet)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range authorityReviewBriefForbiddenFields() {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("decision packet leaked %q in %s", forbidden, body)
		}
	}
}
