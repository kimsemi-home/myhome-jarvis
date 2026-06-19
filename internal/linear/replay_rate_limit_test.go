package linear

import (
	"context"
	"net/http"
	"testing"
)

func TestReplayOfflinePausesWhenRateLimitIsLow(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	root := t.TempDir()
	for _, issueID := range []string{"KIM-16", "KIM-17"} {
		if err := AppendOfflineAction(root, "linear_comment", "queued", map[string]string{
			"issue_id": issueID,
			"body":     "status",
		}); err != nil {
			t.Fatal(err)
		}
	}
	requests := 0
	client := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		requests++
		return linearResponse(1, `{"data":{"commentCreate":{"success":true,"comment":{"id":"comment-id","createdAt":"2026-06-14T00:00:00.000Z"}}}}`), nil
	})}

	result := ReplayOffline(context.Background(), root, client)
	if result.Synced || !result.RateLimited || result.ReplayedCount != 1 || result.RateLimitRemaining != 1 {
		t.Fatalf("unexpected rate-limited replay result: %#v", result)
	}
	if requests != 1 {
		t.Fatalf("requests = %d", requests)
	}
}
