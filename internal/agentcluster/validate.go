package agentcluster

import (
	"fmt"
	"strings"
)

func validatePolicy(policy generatedPolicy) error {
	if strings.TrimSpace(policy.Context) != "AgentCluster" {
		return fmt.Errorf("agent cluster context = %q", policy.Context)
	}
	if strings.TrimSpace(policy.Version) == "" {
		return fmt.Errorf("agent cluster version is required")
	}
	if err := validateSafetyFlags(policy); err != nil {
		return err
	}
	flow := normalizeOrderedList(policy.EvidenceFlow)
	if err := requireOrdered(flow, "observation", "evidence", "code", "verification_evidence", "knowledge_update"); err != nil {
		return err
	}
	lifecycle := normalizeOrderedList(policy.IncidentLifecycle)
	for _, required := range requiredLifecycleStages {
		if !contains(lifecycle, required) {
			return fmt.Errorf("agent cluster lifecycle missing %q", required)
		}
	}
	if len(policy.AgentRoles) < 4 {
		return fmt.Errorf("agent cluster policy must separate producer, reviewer, verifier, and gate roles")
	}
	if len(policy.Sidecars) < 4 {
		return fmt.Errorf("agent cluster policy must include independent sidecars")
	}
	if !contains(normalizeList(policy.Commands), "mhj agent-cluster status") {
		return fmt.Errorf("agent cluster status command is missing")
	}
	return nil
}

func validateSafetyFlags(policy generatedPolicy) error {
	if !policy.PublicSafe {
		return fmt.Errorf("agent cluster policy must be public-safe")
	}
	if policy.RawTranscriptStorageAllowed {
		return fmt.Errorf("agent cluster policy must not store raw transcripts")
	}
	if policy.PrivateDataInEvidenceAllowed {
		return fmt.Errorf("agent cluster policy must not allow private data in public evidence")
	}
	if policy.ExternalAgentExecutionAllowed {
		return fmt.Errorf("agent cluster policy must not enable external agent execution")
	}
	if policy.SelfApprovalAllowed {
		return fmt.Errorf("agent cluster policy must not allow self-approval")
	}
	if policy.ConfidenceSelfReportAllowed {
		return fmt.Errorf("agent cluster confidence must be assessed externally")
	}
	if !policy.AuthorityGateRequired {
		return fmt.Errorf("agent cluster authority gate is required")
	}
	return nil
}
