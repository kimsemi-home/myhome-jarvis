package linear

import (
	"context"
	"testing"
)

func TestReplayOfflineReplaysWriteSafeActionsOnce(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	root := t.TempDir()
	appendReplayHappyPathQueue(t, root)

	requests := 0
	result := ReplayOffline(context.Background(), root, replayHappyPathClient(t, &requests))
	if !result.Synced || result.ReplayedCount != 2 || result.EligibleCount != 2 || result.SkippedCount != 1 {
		t.Fatalf("unexpected replay result: %#v", result)
	}
	if requests != 3 {
		t.Fatalf("requests = %d", requests)
	}

	assertReplayEvidenceRedacted(t, root)
	evidence, err := WriteEvidenceStatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if evidence.SyncedMutationCount != 2 ||
		evidence.LatestSyncedMutation == nil ||
		evidence.LatestSyncedMutation.IssueKey != "KIM-16" {
		t.Fatalf("write evidence = %#v", evidence)
	}

	second := ReplayOffline(context.Background(), root, replayNoCallClient(t))
	if !second.Synced || second.ReplayedCount != 0 || second.AlreadyReplayedCount != 2 {
		t.Fatalf("unexpected second replay result: %#v", second)
	}
}
