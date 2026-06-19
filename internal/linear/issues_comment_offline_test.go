package linear

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAddCommentQueuesOfflineWithoutToken(t *testing.T) {
	root := t.TempDir()
	result := AddComment(context.Background(), root, nil, "MHJ-1", "Started work")
	if result.Synced {
		t.Fatalf("expected offline result: %#v", result)
	}
	data, err := os.ReadFile(filepath.Join(root, "data", "private", "linear-offline-queue.jsonl"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), `"kind":"linear_comment"`) || !strings.Contains(string(data), `"issue_id":"MHJ-1"`) {
		t.Fatalf("offline queue did not contain comment event: %s", string(data))
	}
	status, err := WriteEvidenceStatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.HasSyncedMutation || status.SyncedMutationCount != 0 {
		t.Fatalf("offline comment should not create write evidence: %#v", status)
	}
}
