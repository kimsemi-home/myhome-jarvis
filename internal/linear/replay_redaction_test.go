package linear

import (
	"context"
	"encoding/json"
	"strings"
	"testing"
)

func TestReplayOfflineSummaryRedactsPayloadAndPaths(t *testing.T) {
	root := t.TempDir()
	if err := AppendOfflineAction(root, "linear_comment", "queued", map[string]string{
		"issue_id": "KIM-16",
		"body":     "private body with https://linear.app/private/issue/KIM-16 and /tmp/local",
	}); err != nil {
		t.Fatal(err)
	}

	payload, err := json.Marshal(ReplayOffline(context.Background(), root, nil))
	if err != nil {
		t.Fatal(err)
	}
	body := string(payload)
	for _, expected := range []string{
		`"queue_path":"data/private/linear-offline-queue.jsonl"`,
		`"replay_path":"data/private/linear-offline-replay.jsonl"`,
		`"eligible_count":1`,
	} {
		if !strings.Contains(body, expected) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{"private body", "linear.app/private", "/tmp/local", root} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("replay summary leaked %s in %s", forbidden, body)
		}
	}
}
