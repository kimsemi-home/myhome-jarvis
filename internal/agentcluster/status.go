package agentcluster

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const GeneratedPolicyPath = "generated/agent_cluster.generated.json"

type Status struct {
	Context                       string    `json:"context"`
	Version                       string    `json:"version"`
	PublicSafe                    bool      `json:"public_safe"`
	RawTranscriptStorageAllowed   bool      `json:"raw_transcript_storage_allowed"`
	PrivateDataInEvidenceAllowed  bool      `json:"private_data_in_evidence_allowed"`
	ExternalAgentExecutionAllowed bool      `json:"external_agent_execution_allowed"`
	SelfApprovalAllowed           bool      `json:"self_approval_allowed"`
	ConfidenceSelfReportAllowed   bool      `json:"confidence_self_report_allowed"`
	AuthorityGateRequired         bool      `json:"authority_gate_required"`
	EvidenceStageCount            int       `json:"evidence_stage_count"`
	RoleCount                     int       `json:"role_count"`
	SidecarCount                  int       `json:"sidecar_count"`
	DebtTypeCount                 int       `json:"debt_type_count"`
	FailureConditionCount         int       `json:"failure_condition_count"`
	GeneratedPath                 string    `json:"generated_path"`
	EvidenceFlow                  []string  `json:"evidence_flow"`
	IncidentLifecycle             []string  `json:"incident_lifecycle"`
	Roles                         []Role    `json:"roles"`
	Sidecars                      []Sidecar `json:"sidecars"`
	Signals                       []Signal  `json:"signals"`
	Commands                      []string  `json:"commands"`
	Message                       string    `json:"message"`
	CheckedAt                     string    `json:"checked_at"`
}

type Role struct {
	Key           string   `json:"key"`
	Label         string   `json:"label"`
	ReasoningTier string   `json:"reasoning_tier"`
	Authority     string   `json:"authority"`
	MustProduce   []string `json:"must_produce"`
	MustNot       []string `json:"must_not"`
}

type Sidecar struct {
	Key    string   `json:"key"`
	Label  string   `json:"label"`
	Checks []string `json:"checks"`
}

type Signal struct {
	Key      string `json:"key"`
	Label    string `json:"label"`
	Status   string `json:"status"`
	Evidence string `json:"evidence"`
}

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

func StatusForRoot(root string) (Status, error) {
	policy, err := readGeneratedPolicy(root)
	if err != nil {
		return Status{}, err
	}
	if err := validatePolicy(policy); err != nil {
		return Status{}, err
	}
	status := Status{
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
	status.EvidenceStageCount = len(status.EvidenceFlow)
	status.RoleCount = len(status.Roles)
	status.SidecarCount = len(status.Sidecars)
	status.DebtTypeCount = len(normalizeList(policy.DebtTypes))
	status.FailureConditionCount = len(normalizeList(policy.FailureConditions))
	status.Message = "Agent cluster policy is evidence-first and governance-gated; external agents and self-approval are disabled."
	return status, nil
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

func validatePolicy(policy generatedPolicy) error {
	if strings.TrimSpace(policy.Context) != "AgentCluster" {
		return fmt.Errorf("agent cluster context = %q", policy.Context)
	}
	if strings.TrimSpace(policy.Version) == "" {
		return fmt.Errorf("agent cluster version is required")
	}
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
	flow := normalizeOrderedList(policy.EvidenceFlow)
	if err := requireOrdered(flow, "observation", "evidence", "code", "verification_evidence", "knowledge_update"); err != nil {
		return err
	}
	lifecycle := normalizeOrderedList(policy.IncidentLifecycle)
	for _, required := range []string{"observed", "evidence_recorded", "classified", "owner_assigned", "verified", "knowledge_updated"} {
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

func requireOrdered(values []string, required ...string) error {
	last := -1
	for _, item := range required {
		index := indexOf(values, item)
		if index < 0 {
			return fmt.Errorf("agent cluster evidence flow missing %q", item)
		}
		if index <= last {
			return fmt.Errorf("agent cluster evidence flow must keep %q after the previous stage", item)
		}
		last = index
	}
	return nil
}

func sanitizeRoles(roles []Role) []Role {
	clean := make([]Role, 0, len(roles))
	for _, role := range roles {
		role.Key = normalizeToken(role.Key)
		role.Label = publicText(role.Label)
		role.ReasoningTier = strings.TrimSpace(role.ReasoningTier)
		role.Authority = normalizeToken(role.Authority)
		role.MustProduce = normalizeList(role.MustProduce)
		role.MustNot = normalizeList(role.MustNot)
		if role.Key == "" || role.Label == "" {
			continue
		}
		clean = append(clean, role)
	}
	sort.Slice(clean, func(i, j int) bool {
		return clean[i].Key < clean[j].Key
	})
	return clean
}

func sanitizeSidecars(sidecars []Sidecar) []Sidecar {
	clean := make([]Sidecar, 0, len(sidecars))
	for _, sidecar := range sidecars {
		sidecar.Key = normalizeToken(sidecar.Key)
		sidecar.Label = publicText(sidecar.Label)
		sidecar.Checks = normalizeList(sidecar.Checks)
		if sidecar.Key == "" || sidecar.Label == "" {
			continue
		}
		clean = append(clean, sidecar)
	}
	sort.Slice(clean, func(i, j int) bool {
		return clean[i].Key < clean[j].Key
	})
	return clean
}

func sanitizeSignals(signals []Signal) []Signal {
	clean := make([]Signal, 0, len(signals))
	for _, signal := range signals {
		signal.Key = normalizeToken(signal.Key)
		signal.Label = publicText(signal.Label)
		signal.Status = normalizeToken(signal.Status)
		signal.Evidence = publicText(signal.Evidence)
		if signal.Key == "" || signal.Label == "" {
			continue
		}
		clean = append(clean, signal)
	}
	sort.Slice(clean, func(i, j int) bool {
		return clean[i].Key < clean[j].Key
	})
	return clean
}

func normalizeList(values []string) []string {
	seen := map[string]bool{}
	clean := make([]string, 0, len(values))
	for _, value := range values {
		item := normalizeToken(value)
		if item == "" || seen[item] {
			continue
		}
		seen[item] = true
		clean = append(clean, item)
	}
	sort.Strings(clean)
	return clean
}

func normalizeOrderedList(values []string) []string {
	seen := map[string]bool{}
	clean := make([]string, 0, len(values))
	for _, value := range values {
		item := normalizeToken(value)
		if item == "" || seen[item] {
			continue
		}
		seen[item] = true
		clean = append(clean, item)
	}
	return clean
}

func normalizeToken(value string) string {
	return strings.TrimSpace(strings.ToLower(value))
}

func publicText(value string) string {
	value = strings.TrimSpace(value)
	value = strings.ReplaceAll(value, "\n", " ")
	return value
}

func contains(values []string, wanted string) bool {
	return indexOf(values, wanted) >= 0
}

func indexOf(values []string, wanted string) int {
	for index, value := range values {
		if value == wanted {
			return index
		}
	}
	return -1
}
