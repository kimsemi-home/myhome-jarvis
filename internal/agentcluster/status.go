package agentcluster

import (
	"strings"
	"time"
)

const GeneratedPolicyPath = "generated/agent_cluster.generated.json"

func StatusForRoot(root string) (Status, error) {
	policy, err := readGeneratedPolicy(root)
	if err != nil {
		return Status{}, err
	}
	if err := validatePolicy(policy); err != nil {
		return Status{}, err
	}
	status := statusFromPolicy(policy)
	status.EvidenceStageCount = len(status.EvidenceFlow)
	status.RoleCount = len(status.Roles)
	status.SidecarCount = len(status.Sidecars)
	status.DebtTypeCount = len(normalizeList(policy.DebtTypes))
	status.FailureConditionCount = len(normalizeList(policy.FailureConditions))
	status.Message = "Agent cluster policy is evidence-first and governance-gated; external agents and self-approval are disabled."
	return status, nil
}

func statusFromPolicy(policy generatedPolicy) Status {
	return Status{
		Context:                       strings.TrimSpace(policy.Context),
		Version:                       strings.TrimSpace(policy.Version),
		PublicSafe:                    policy.PublicSafe,
		RawTranscriptStorageAllowed:   policy.RawTranscriptStorageAllowed,
		PrivateDataInEvidenceAllowed:  policy.PrivateDataInEvidenceAllowed,
		ExternalAgentExecutionAllowed: policy.ExternalAgentExecutionAllowed,
		SelfApprovalAllowed:           policy.SelfApprovalAllowed,
		ConfidenceSelfReportAllowed:   policy.ConfidenceSelfReportAllowed,
		AuthorityGateRequired:         policy.AuthorityGateRequired,
		EvidenceFlow:                  normalizeOrderedList(policy.EvidenceFlow),
		IncidentLifecycle:             normalizeOrderedList(policy.IncidentLifecycle),
		Roles:                         sanitizeRoles(policy.AgentRoles),
		Sidecars:                      sanitizeSidecars(policy.Sidecars),
		Signals:                       sanitizeSignals(policy.Signals),
		Commands:                      normalizeList(policy.Commands),
		GeneratedPath:                 GeneratedPolicyPath,
		CheckedAt:                     time.Now().UTC().Format(time.RFC3339),
	}
}
