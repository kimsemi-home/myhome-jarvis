package orchestrator

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/linear"
	"github.com/kimsemi-home/myhome-jarvis/internal/planner"
)

func redactedCheckpointFixture() Checkpoint {
	return Checkpoint{
		Task:           "loop once",
		LinearStatus:   redactedLinearStatus(),
		LinearNext:     redactedLinearNext(),
		PlannerStatus:  redactedPlannerStatus(),
		SecurityStatus: redactedSecurityStatus(),
		Result:         "checkpoint recorded",
		Next:           "continue",
	}
}

func redactedLinearStatus() linear.StatusSummary {
	return linear.StatusSummary{
		Mode:             "online",
		TokenConfigured:  true,
		Synced:           true,
		QueuePath:        "data/private/linear-offline-queue.jsonl",
		ViewerConfigured: true,
		TeamCount:        1,
		Message:          "ok",
	}
}

func redactedLinearNext() *linear.OperationSummary {
	return &linear.OperationSummary{
		Mode:       "online",
		Synced:     true,
		QueuePath:  "data/private/linear-offline-queue.jsonl",
		HTTPStatus: 200,
		Message:    "Selected next project Linear issue.",
		IssueCount: 1,
		Issue: &linear.IssueSummary{
			Identifier: "KIM-13",
			Title:      "[myhome-jarvis] Include project queue status in loop checkpoints",
			StateType:  "started",
		},
	}
}

func redactedPlannerStatus() planner.Status {
	return planner.Status{
		LoopMode:                  "closed-loop",
		TaskCount:                 6,
		CompletedCount:            5,
		BlockedExternalWriteCount: 1,
		BlockedExternalWriteTasks: []planner.Task{redactedBlockedTask()},
		LinearTemplateCount:       2,
		QualityRequired:           true,
		LinearOfflineFallback:     true,
		KnowledgeIndexRequired:    true,
		KnowledgeEvidence:         redactedKnowledgeEvidence(),
		CheckpointRoot:            "data/private/checkpoints",
		CheckedAt:                 "2026-06-15T00:00:00Z",
	}
}

func redactedBlockedTask() planner.Task {
	return planner.Task{
		ID: "linear_sync", Title: "Sync Linear only after explicit external-write approval",
		Owner: "go", Status: "blocked_external_write", DependsOn: []string{"quality_gate"},
	}
}
