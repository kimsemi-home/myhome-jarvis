package authority

import "fmt"

func normalizeApprovalGrantFlags(
	request ApprovalDecisionRequest,
) (ApprovalGrantFlags, error) {
	if request.ApprovalGranted == nil ||
		request.RepoCreationAllowed == nil ||
		request.WorkflowChangesAllowed == nil ||
		request.ExternalWritesAllowed == nil ||
		request.SelfApprovalAllowed == nil {
		return ApprovalGrantFlags{}, fmt.Errorf("approval requires exact grant flags")
	}
	flags := ApprovalGrantFlags{
		ApprovalGranted:        *request.ApprovalGranted,
		RepoCreationAllowed:    *request.RepoCreationAllowed,
		WorkflowChangesAllowed: *request.WorkflowChangesAllowed,
		ExternalWritesAllowed:  *request.ExternalWritesAllowed,
		SelfApprovalAllowed:    *request.SelfApprovalAllowed,
	}
	if !flags.ApprovalGranted || flags.SelfApprovalAllowed {
		return ApprovalGrantFlags{}, fmt.Errorf("approval flags are not allowed")
	}
	return flags, nil
}

func approvalFlagsMatchScope(record ApprovalDecisionRecord) bool {
	flags := record.GrantFlags
	switch record.Scope {
	case "repo_creation":
		return flags.RepoCreationAllowed &&
			!flags.WorkflowChangesAllowed && !flags.ExternalWritesAllowed
	case "workflow_change":
		return flags.WorkflowChangesAllowed &&
			!flags.RepoCreationAllowed && !flags.ExternalWritesAllowed
	case "external_write":
		return flags.ExternalWritesAllowed &&
			!flags.RepoCreationAllowed && !flags.WorkflowChangesAllowed
	default:
		return false
	}
}
