package authority

import "time"

func applyApprovalRecords(
	status *ApprovalDecisionStatus,
	records []ApprovalDecisionRecord,
	now time.Time,
) {
	for _, record := range records {
		summary := approvalScopeSummary(record, now)
		if summary.LeaseState == "active" {
			if !summary.CanUnlockScopeOnly {
				status.UnrelatedAuthorityGranted = true
				status.InvalidRecordCount++
				continue
			}
			status.ActiveApprovalCount++
			status.ScopeSummaries = append(status.ScopeSummaries, summary)
			status.LatestScope = summary.Scope
			status.LatestTarget = summary.Target
			status.LatestLeaseState = summary.LeaseState
			applyScopeGrant(status, summary)
			continue
		}
		status.ExpiredApprovalCount++
	}
}

func applyScopeGrant(
	status *ApprovalDecisionStatus,
	summary ApprovalScopeSummary,
) {
	switch summary.Scope {
	case "repo_creation":
		status.CanCreateRepo = summary.GrantFlags.RepoCreationAllowed
	case "workflow_change":
		status.CanChangeWorkflow = summary.GrantFlags.WorkflowChangesAllowed
	case "external_write":
		status.CanWriteExternal = summary.GrantFlags.ExternalWritesAllowed
	}
}
