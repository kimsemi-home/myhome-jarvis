package agentcluster

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type generatedPolicy struct {
	Context                       string    `json:"context"`
	Version                       string    `json:"version"`
	PublicSafe                    bool      `json:"public_safe"`
	RawTranscriptStorageAllowed   bool      `json:"raw_transcript_storage_allowed"`
	PrivateDataInEvidenceAllowed  bool      `json:"private_data_in_evidence_allowed"`
	ExternalAgentExecutionAllowed bool      `json:"external_agent_execution_allowed"`
	SelfApprovalAllowed           bool      `json:"self_approval_allowed"`
	ConfidenceSelfReportAllowed   bool      `json:"confidence_self_report_allowed"`
	AuthorityGateRequired         bool      `json:"authority_gate_required"`
	EvidenceFlow                  []string  `json:"evidence_flow"`
	IncidentLifecycle             []string  `json:"incident_lifecycle"`
	AgentRoles                    []Role    `json:"agent_roles"`
	Sidecars                      []Sidecar `json:"sidecars"`
	DebtTypes                     []string  `json:"debt_types"`
	FailureConditions             []string  `json:"failure_conditions"`
	Signals                       []Signal  `json:"signals"`
	Commands                      []string  `json:"commands"`
}

func readGeneratedPolicy(root string) (generatedPolicy, error) {
	body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(GeneratedPolicyPath)))
	if err != nil {
		return generatedPolicy{}, err
	}
	var policy generatedPolicy
	if err := json.Unmarshal(body, &policy); err != nil {
		return generatedPolicy{}, err
	}
	return policy, nil
}
