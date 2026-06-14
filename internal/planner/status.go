package planner

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/knowledge"
	"github.com/kimsemi-home/myhome-jarvis/internal/linear"
)

const generatedRelativePath = "generated/planner.generated.json"

type Policy struct {
	LoopMode                             string            `json:"loop_mode"`
	MaxTaskScope                         string            `json:"max_task_scope"`
	CheckpointRoot                       string            `json:"checkpoint_root"`
	QualityRequired                      bool              `json:"quality_required"`
	LinearOfflineFallback                bool              `json:"linear_offline_fallback"`
	KnowledgeIndexRequiredBeforePlanning bool              `json:"knowledge_index_required_before_planning"`
	KnowledgeIndexDefaultQuery           string            `json:"knowledge_index_default_query"`
	ExternalWriteGate                    ExternalWriteGate `json:"external_write_gate"`
	DefaultNext                          string            `json:"default_next"`
	TaskGraph                            []Task            `json:"task_graph"`
	LinearTemplates                      []LinearTemplate  `json:"linear_templates"`
}

type Task struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Owner     string   `json:"owner"`
	Status    string   `json:"status"`
	DependsOn []string `json:"depends_on"`
}

type LinearTemplate struct {
	Name        string   `json:"name"`
	TitlePrefix string   `json:"title_prefix"`
	Labels      []string `json:"labels"`
}

type ExternalWriteGate struct {
	StandingBoundary        bool   `json:"standing_boundary"`
	ApprovalRequired        bool   `json:"approval_required"`
	MutationSuccessRequired bool   `json:"mutation_success_required"`
	BoundaryTaskID          string `json:"boundary_task_id"`
	EvidencePath            string `json:"evidence_path"`
}

type ExternalWriteGateStatus struct {
	StandingBoundary        bool   `json:"standing_boundary"`
	ApprovalRequired        bool   `json:"approval_required"`
	MutationSuccessRequired bool   `json:"mutation_success_required"`
	BoundaryTaskID          string `json:"boundary_task_id"`
	BoundaryTaskBlocked     bool   `json:"boundary_task_blocked"`
	EvidencePath            string `json:"evidence_path"`
}

type Status struct {
	LoopMode                  string                     `json:"loop_mode"`
	TaskCount                 int                        `json:"task_count"`
	ReadyCount                int                        `json:"ready_count"`
	CompletedCount            int                        `json:"completed_count"`
	BlockedExternalWriteCount int                        `json:"blocked_external_write_count"`
	NextTask                  *Task                      `json:"next_task,omitempty"`
	BlockedExternalWriteTasks []Task                     `json:"blocked_external_write_tasks,omitempty"`
	ExternalWriteGate         ExternalWriteGateStatus    `json:"external_write_gate"`
	LinearWriteEvidence       linear.WriteEvidenceStatus `json:"linear_write_evidence"`
	LinearTemplateCount       int                        `json:"linear_template_count"`
	QualityRequired           bool                       `json:"quality_required"`
	LinearOfflineFallback     bool                       `json:"linear_offline_fallback"`
	KnowledgeIndexRequired    bool                       `json:"knowledge_index_required"`
	KnowledgeEvidence         *knowledge.Evidence        `json:"knowledge_evidence,omitempty"`
	CheckpointRoot            string                     `json:"checkpoint_root"`
	CheckedAt                 string                     `json:"checked_at"`
}

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(filepath.Join(root, filepath.FromSlash(generatedRelativePath)))
	if err != nil {
		return Status{}, err
	}
	if err := validatePolicy(policy); err != nil {
		return Status{}, err
	}
	status := Status{
		LoopMode:               policy.LoopMode,
		TaskCount:              len(policy.TaskGraph),
		LinearTemplateCount:    len(policy.LinearTemplates),
		QualityRequired:        policy.QualityRequired,
		LinearOfflineFallback:  policy.LinearOfflineFallback,
		KnowledgeIndexRequired: policy.KnowledgeIndexRequiredBeforePlanning,
		CheckpointRoot:         policy.CheckpointRoot,
		CheckedAt:              time.Now().UTC().Format(time.RFC3339),
	}
	writeEvidence, err := linear.WriteEvidenceStatusForPath(root, policy.ExternalWriteGate.EvidencePath)
	if err != nil {
		return Status{}, err
	}
	status.LinearWriteEvidence = writeEvidence
	status.ExternalWriteGate = ExternalWriteGateStatus{
		StandingBoundary:        policy.ExternalWriteGate.StandingBoundary,
		ApprovalRequired:        policy.ExternalWriteGate.ApprovalRequired,
		MutationSuccessRequired: policy.ExternalWriteGate.MutationSuccessRequired,
		BoundaryTaskID:          strings.TrimSpace(policy.ExternalWriteGate.BoundaryTaskID),
		EvidencePath:            writeEvidence.EvidencePath,
	}
	if policy.KnowledgeIndexRequiredBeforePlanning {
		query := strings.TrimSpace(policy.KnowledgeIndexDefaultQuery)
		if query == "" {
			query = "planner"
		}
		evidence, err := knowledge.Search(root, query)
		if err != nil {
			return Status{}, err
		}
		summary := knowledge.SummarizeSearch(evidence)
		status.KnowledgeEvidence = &summary
	}
	taskStatuses := make(map[string]string, len(policy.TaskGraph))
	for _, task := range policy.TaskGraph {
		normalized := normalizeStatus(task.Status)
		taskStatuses[task.ID] = normalized
		switch normalized {
		case "ready":
			status.ReadyCount++
		case "completed":
			status.CompletedCount++
		case "blocked_external_write":
			status.BlockedExternalWriteCount++
			status.BlockedExternalWriteTasks = append(status.BlockedExternalWriteTasks, task)
			if task.ID == status.ExternalWriteGate.BoundaryTaskID {
				status.ExternalWriteGate.BoundaryTaskBlocked = true
			}
		}
	}
	for index := range policy.TaskGraph {
		task := policy.TaskGraph[index]
		if normalizeStatus(task.Status) == "ready" && dependenciesSatisfied(task, taskStatuses) {
			status.NextTask = &task
			break
		}
	}
	return status, nil
}

func ReadPolicy(path string) (Policy, error) {
	file, err := os.Open(path)
	if err != nil {
		return Policy{}, err
	}
	defer file.Close()

	var policy Policy
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&policy); err != nil {
		return Policy{}, err
	}
	return policy, nil
}

func validatePolicy(policy Policy) error {
	if strings.TrimSpace(policy.LoopMode) == "" {
		return errors.New("planner loop mode is required")
	}
	if len(policy.TaskGraph) == 0 {
		return errors.New("planner task graph is required")
	}
	if err := validateCheckpointRoot(policy.CheckpointRoot); err != nil {
		return err
	}
	if err := validateExternalWriteGate(policy.ExternalWriteGate); err != nil {
		return err
	}
	ids := map[string]bool{}
	hasExternalWriteBoundary := false
	gateTaskBlocked := false
	for _, task := range policy.TaskGraph {
		id := strings.TrimSpace(task.ID)
		if id == "" {
			return errors.New("planner task id is required")
		}
		if ids[id] {
			return fmt.Errorf("duplicate planner task id %q", id)
		}
		ids[id] = true
		if normalizeStatus(task.Status) == "blocked_external_write" {
			hasExternalWriteBoundary = true
			if id == strings.TrimSpace(policy.ExternalWriteGate.BoundaryTaskID) {
				gateTaskBlocked = true
			}
		}
		if !validStatus(task.Status) {
			return fmt.Errorf("planner task %q has invalid status %q", task.ID, task.Status)
		}
	}
	if !hasExternalWriteBoundary {
		return errors.New("planner task graph must include an external-write boundary")
	}
	if !gateTaskBlocked {
		return fmt.Errorf("planner external-write gate task %q must be blocked_external_write", policy.ExternalWriteGate.BoundaryTaskID)
	}
	if policy.KnowledgeIndexRequiredBeforePlanning && strings.TrimSpace(policy.KnowledgeIndexDefaultQuery) == "" {
		return errors.New("planner knowledge index default query is required")
	}
	for _, task := range policy.TaskGraph {
		for _, dependency := range task.DependsOn {
			if !ids[dependency] {
				return fmt.Errorf("planner task %q depends on unknown task %q", task.ID, dependency)
			}
		}
	}
	return nil
}

func validateExternalWriteGate(gate ExternalWriteGate) error {
	if !gate.StandingBoundary {
		return errors.New("planner external-write gate must keep a standing boundary")
	}
	if !gate.ApprovalRequired {
		return errors.New("planner external-write gate must require approval")
	}
	if !gate.MutationSuccessRequired {
		return errors.New("planner external-write gate must require mutation success")
	}
	if strings.TrimSpace(gate.BoundaryTaskID) == "" {
		return errors.New("planner external-write gate boundary task id is required")
	}
	return validatePrivateRepoPath("planner external-write gate evidence path", gate.EvidencePath)
}

func validateCheckpointRoot(value string) error {
	return validatePrivateRepoPath("planner checkpoint root", value)
}

func validatePrivateRepoPath(name string, value string) error {
	value = strings.TrimSpace(value)
	if value == "" {
		return fmt.Errorf("%s is required", name)
	}
	native := filepath.FromSlash(value)
	if filepath.IsAbs(native) {
		return fmt.Errorf("%s must be repo-relative", name)
	}
	clean := filepath.ToSlash(filepath.Clean(native))
	if clean == "." || clean == ".." || strings.HasPrefix(clean, "../") {
		return fmt.Errorf("%s must stay inside the repo", name)
	}
	if !strings.HasPrefix(clean, "data/private/") {
		return fmt.Errorf("%s must stay under data/private", name)
	}
	return nil
}

func dependenciesSatisfied(task Task, taskStatuses map[string]string) bool {
	for _, dependency := range task.DependsOn {
		switch taskStatuses[dependency] {
		case "ready", "completed":
		default:
			return false
		}
	}
	return true
}

func normalizeStatus(status string) string {
	return strings.TrimSpace(strings.ToLower(status))
}

func validStatus(status string) bool {
	switch normalizeStatus(status) {
	case "ready", "completed", "blocked", "blocked_external_write":
		return true
	default:
		return false
	}
}
