package linear

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestReplayOfflineFailedEntryRemainsQueued(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	root := t.TempDir()
	if err := AppendOfflineAction(root, "linear_comment", "queued", map[string]string{
		"issue_id": "KIM-16",
		"body":     "status",
	}); err != nil {
		t.Fatal(err)
	}
	client := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		return linearResponse(250, `{"data":{"commentCreate":{"success":false,"comment":null}}}`), nil
	})}

	result := ReplayOffline(context.Background(), root, client)
	if result.Synced || result.FailedCount != 1 || result.ReplayedCount != 0 {
		t.Fatalf("unexpected failed replay result: %#v", result)
	}
	if _, err := os.Stat(filepath.Join(root, OfflineReplayRelativePath)); !os.IsNotExist(err) {
		t.Fatalf("failed replay should not create replay evidence, err=%v", err)
	}
}
