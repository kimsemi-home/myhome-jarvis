package orchestrator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/knowledge"
	"github.com/kimsemi-home/myhome-jarvis/internal/linear"
	"github.com/kimsemi-home/myhome-jarvis/internal/planner"
	"github.com/kimsemi-home/myhome-jarvis/internal/security"
)

func TestWriteCheckpointStoresAggregateSecurityStatus(t *testing.T) {
	root := t.TempDir()
	path, err := WriteCheckpoint(root, Checkpoint{
		Task: "loop once",
		LinearStatus: linear.StatusSummary{
			Mode:             "online",
			TokenConfigured:  true,
			Synced:           true,
			QueuePath:        "data/private/linear-offline-queue.jsonl",
			ViewerConfigured: true,
			TeamCount:        1,
			Message:          "ok",
		},
		PlannerStatus: planner.Status{
			LoopMode:                  "closed-loop",
			TaskCount:                 6,
			ReadyCount:                0,
			CompletedCount:            5,
			BlockedExternalWriteCount: 1,
			BlockedExternalWriteTasks: []planner.Task{{
				ID:        "linear_sync",
				Title:     "Sync Linear only after explicit external-write approval",
				Owner:     "go",
				Status:    "blocked_external_write",
				DependsOn: []string{"quality_gate"},
			}},
			LinearTemplateCount:    2,
			QualityRequired:        true,
			LinearOfflineFallback:  true,
			KnowledgeIndexRequired: true,
			KnowledgeEvidence: &knowledge.Evidence{
				Query:        "planner KnowledgeIndex Linear closed loop",
				ConceptCount: 3,
				HitCount:     9,
				LinearIssues: []string{"KIM-14"},
				MustRead:     []string{"generated/concepts.generated.json", "docs/knowledge-index.md"},
				CheckedAt:    "2026-06-15T00:00:00Z",
			},
			CheckpointRoot: "data/private/checkpoints",
			CheckedAt:      "2026-06-15T00:00:00Z",
		},
		SecurityStatus: security.Status{
			OK:                  true,
			CurrentOK:           true,
			HistoryOK:           true,
			CurrentFindingCount: 0,
			HistoryFindingCount: 0,
			CheckedAt:           "2026-06-15T00:00:00Z",
		},
		Result: "checkpoint recorded",
		Next:   "continue",
	})
	if err != nil {
		t.Fatal(err)
	}
	body, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	text := string(body)
	if !strings.Contains(text, `"security_status"`) {
		t.Fatalf("expected aggregate security status in %s", text)
	}
	if !strings.Contains(text, `"planner_status"`) || !strings.Contains(text, `"blocked_external_write_count": 1`) {
		t.Fatalf("expected planner status in %s", text)
	}
	if !strings.Contains(text, `"knowledge_evidence"`) || !strings.Contains(text, `"KIM-14"`) {
		t.Fatalf("expected knowledge evidence in %s", text)
	}
	for _, forbidden := range []string{`"security_report"`, `"findings"`, `"root"`, `"viewer"`, `"teams"`} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("checkpoint leaked %s in %s", forbidden, text)
		}
	}
}

func TestWriteCheckpointUsesCollisionResistantNames(t *testing.T) {
	root := t.TempDir()
	first, err := WriteCheckpoint(root, Checkpoint{Task: "first"})
	if err != nil {
		t.Fatal(err)
	}
	second, err := WriteCheckpoint(root, Checkpoint{Task: "second"})
	if err != nil {
		t.Fatal(err)
	}
	if first == second {
		t.Fatalf("checkpoint paths collided: %s", first)
	}
	for _, path := range []string{first, second} {
		if !strings.Contains(filepath.Base(path), ".") {
			t.Fatalf("checkpoint filename should include sub-second precision: %s", path)
		}
	}
}
