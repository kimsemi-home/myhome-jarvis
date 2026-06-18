package planner

import (
	"encoding/json"
	"os"
	"testing"
)

func writePlannerPolicyFixture(
	t *testing.T,
	path string,
	checkpointRoot string,
	repoSafetyStatus string,
) {
	t.Helper()
	policy := Policy{
		LoopMode:                             "closed-loop",
		MaxTaskScope:                         "one file",
		CheckpointRoot:                       checkpointRoot,
		QualityRequired:                      true,
		LinearOfflineFallback:                true,
		ExternalWriteGate:                    plannerTestExternalWriteGate(),
		DefaultNext:                          "ready",
		TaskGraph:                            plannerTestTasks(repoSafetyStatus),
		LinearTemplates:                      []LinearTemplate{},
		KnowledgeIndexDefaultQuery:           "",
		KnowledgeIndexRequiredBeforePlanning: false,
	}
	data, err := json.Marshal(policy)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatal(err)
	}
}

func plannerTestExternalWriteGate() ExternalWriteGate {
	return ExternalWriteGate{
		StandingBoundary:        true,
		ApprovalRequired:        true,
		MutationSuccessRequired: true,
		BoundaryTaskID:          "linear_sync",
		EvidencePath:            "data/private/linear-write-evidence.jsonl",
	}
}

func plannerTestTasks(repoSafetyStatus string) []Task {
	return []Task{
		{
			ID:        "repo_safety",
			Title:     "Repo safety",
			Owner:     "go",
			Status:    repoSafetyStatus,
			DependsOn: []string{},
		},
		{
			ID:        "linear_sync",
			Title:     "Linear sync",
			Owner:     "go",
			Status:    "blocked_external_write",
			DependsOn: []string{"repo_safety"},
		},
	}
}
