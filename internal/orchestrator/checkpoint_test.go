package orchestrator

import (
	"os"
	"strings"
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/linear"
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
	for _, forbidden := range []string{`"security_report"`, `"findings"`, `"root"`, `"viewer"`, `"teams"`} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("checkpoint leaked %s in %s", forbidden, text)
		}
	}
}
