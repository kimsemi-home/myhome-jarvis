package authority

import (
	"testing"
	"time"
)

func TestNormalizeApprovalDecisionRequiresScopedHumanLease(t *testing.T) {
	now := time.Date(2026, 6, 21, 6, 0, 0, 0, time.UTC)
	record, err := normalizeApprovalDecisionRequest(
		testPolicy(),
		approvalFixtureRequest(now),
		approvalFixturePacket(now),
		now,
	)
	if err != nil {
		t.Fatal(err)
	}
	if record.Scope != "repo_creation" ||
		!record.GrantFlags.RepoCreationAllowed ||
		record.GrantFlags.WorkflowChangesAllowed ||
		record.GrantFlags.ExternalWritesAllowed ||
		record.GrantFlags.SelfApprovalAllowed ||
		record.ReviewerIsRequester {
		t.Fatalf("approval record = %#v", record)
	}
}

func TestNormalizeApprovalDecisionRejectsUnsafeInputs(t *testing.T) {
	now := time.Date(2026, 6, 21, 6, 0, 0, 0, time.UTC)
	for name, mutate := range map[string]func(*ApprovalDecisionRequest){
		"self reviewer": func(request *ApprovalDecisionRequest) {
			trueValue := true
			request.ReviewerIsRequester = &trueValue
		},
		"stale packet": func(request *ApprovalDecisionRequest) {
			request.DecisionPacketCheckedAt =
				now.Add(-25 * time.Hour).Format(time.RFC3339)
		},
		"expired lease": func(request *ApprovalDecisionRequest) {
			request.ExpiresAt = now.Add(-time.Minute).Format(time.RFC3339)
		},
		"private target": func(request *ApprovalDecisionRequest) {
			request.Target = "data/private/external-evidence"
		},
		"wrong target": func(request *ApprovalDecisionRequest) {
			request.Target = "kimsemi-home/not-the-approved-repo"
		},
		"mixed grant": func(request *ApprovalDecisionRequest) {
			trueValue := true
			request.WorkflowChangesAllowed = &trueValue
		},
	} {
		t.Run(name, func(t *testing.T) {
			request := approvalFixtureRequest(now)
			mutate(&request)
			_, err := normalizeApprovalDecisionRequest(
				testPolicy(),
				request,
				approvalFixturePacket(now),
				now,
			)
			if err == nil {
				t.Fatal("expected approval decision rejection")
			}
		})
	}
}
