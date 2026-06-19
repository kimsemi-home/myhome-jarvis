package linear

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestTransitionQueuesOfflineWithoutToken(t *testing.T) {
	root := t.TempDir()
	result := TransitionIssue(context.Background(), root, nil, "MHJ-1", "In Progress")
	if result.Synced {
		t.Fatalf("expected offline result: %#v", result)
	}
	data, err := os.ReadFile(filepath.Join(root, "data", "private", "linear-offline-queue.jsonl"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), `"kind":"linear_transition"`) || !strings.Contains(string(data), `"state":"In Progress"`) {
		t.Fatalf("offline queue did not contain transition event: %s", string(data))
	}
}
