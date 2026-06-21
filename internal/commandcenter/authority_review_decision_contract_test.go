package commandcenter

import "testing"

func TestAuthorityReviewDecisionPacketIncludesDecisionContract(t *testing.T) {
	packet, err := AuthorityReviewDecisionPacketForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	contract := packet.DecisionContract
	if !contract.PublicSafe || !contract.ReviewOnly || contract.CanApplyDecision {
		t.Fatalf("decision contract boundary = %#v", contract)
	}
	if len(contract.ReadyCapabilitiesNonBlocking) !=
		len(packet.CapabilityReadiness.ReadyCapabilityKeys) ||
		len(contract.ContractItems) != len(packet.GatedCapabilityKeys) {
		t.Fatalf("decision contract capabilities = %#v", contract)
	}
	for _, key := range packet.CapabilityReadiness.ReadyCapabilityKeys {
		if !containsString(contract.ReadyCapabilitiesNonBlocking, key) {
			t.Fatalf("ready capability %s missing in %#v", key, contract)
		}
	}
	assertContractItem(t, contract, "shorts_factory_control_plane",
		"public_repo_and_workflow_control", "public_repo_change_review")
	assertContractItem(t, contract, "self_improvement_loop",
		"closed_loop_authority_and_external_write_boundary", "workflow_change_review")
	if contract.ForbiddenGrantFlags.ApprovalGranted ||
		contract.ForbiddenGrantFlags.ExternalWritesAllowed ||
		contract.ForbiddenGrantFlags.RepoCreationAllowed ||
		contract.ForbiddenGrantFlags.WorkflowChangesAllowed ||
		contract.ForbiddenGrantFlags.SelfApprovalAllowed {
		t.Fatalf("decision contract grants authority = %#v", contract.ForbiddenGrantFlags)
	}
}

func assertContractItem(
	t *testing.T,
	contract AuthorityReviewDecisionContract,
	key string,
	scope string,
	reviewClass string,
) {
	t.Helper()
	for _, item := range contract.ContractItems {
		if item.CapabilityKey != key {
			continue
		}
		if item.Scope != scope || item.RequiredReviewClass != reviewClass ||
			!item.HumanDecisionRecordRequired ||
			!containsString(item.RequiredEvidenceKeys, "storage_evidence") ||
			!containsString(item.RequiredEvidenceKeys, "external_evidence") ||
			!containsString(item.RequiredEvidenceKeys, "capability_readiness") ||
			item.ThisPacketGrantsApproval || item.AllowsExternalWrites ||
			item.AllowsRepoCreation || item.AllowsWorkflowChanges ||
			item.AllowsSelfApproval {
			t.Fatalf("contract item %s = %#v", key, item)
		}
		return
	}
	t.Fatalf("missing contract item %s in %#v", key, contract.ContractItems)
}
