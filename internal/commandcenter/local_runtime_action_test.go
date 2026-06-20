package commandcenter

import "testing"

func TestNextSafeActionRepairsLocalRuntimeFirst(t *testing.T) {
	status := Status{
		PublicSafe: true,
		AuthorityReview: AuthorityReviewSummary{
			NextSafeAction: "await_human_authority_review",
		},
		BlockedGates: []GateSummary{
			{Key: "authority"},
			{Key: "local_runtime"},
		},
	}
	if got := nextSafeAction(status); got != "repair_local_runtime_health" {
		t.Fatalf("next safe action = %q", got)
	}
}

func TestWorkItemUsesLocalRuntimeEvidenceWhenRuntimeIsNext(t *testing.T) {
	status := Status{
		NextSafeAction: "repair_local_runtime_health",
		LocalRuntime:   LocalRuntimeSummary{EvidenceRef: "local_runtime:supervisor"},
		AuthorityReview: AuthorityReviewSummary{
			EvidenceRef: "authority_review_request:test",
		},
		BlockedGates: []GateSummary{
			{Key: "authority"}, {Key: "local_runtime"},
		},
	}
	item := summarizeWorkItem(status)
	if item.NextSafeAction != "repair_local_runtime_health" ||
		item.DecisionKey != "repair_local_runtime_health" {
		t.Fatalf("work item action = %#v", item)
	}
	if item.WorkItemRef != "universal_work_item:local_runtime" ||
		item.WorkItemState != "runtime_health_debt" {
		t.Fatalf("work item runtime state = %#v", item)
	}
	if item.EvidenceRef != "local_runtime:supervisor" ||
		!containsString(item.BlockedGateKeys, "local_runtime") {
		t.Fatalf("work item runtime gate = %#v", item)
	}
	if item.ApprovalGranted || item.ExternalWritesAllowed || item.SelfApprovalAllowed {
		t.Fatalf("work item granted authority = %#v", item)
	}
}
