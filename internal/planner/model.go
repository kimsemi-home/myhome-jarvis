package planner

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/knowledge"
	"github.com/kimsemi-home/myhome-jarvis/internal/linear"
)

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
