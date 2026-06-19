package linear

import (
	"context"
	"net/http"
	"testing"
)

func TestReplayOfflineSkipsOutOfScopeIssueKeys(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	t.Setenv("LINEAR_TEAM_KEY", "KIM")
	root := t.TempDir()
	if err := AppendOfflineAction(root, "linear_comment", "queued", map[string]string{
		"issue_id": "MHJ-1",
		"body":     "old team",
	}); err != nil {
		t.Fatal(err)
	}
	client := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		t.Fatalf("out-of-scope offline actions should not call Linear")
		return nil, nil
	})}

	result := ReplayOffline(context.Background(), root, client)
	if !result.Synced || result.EligibleCount != 0 || result.SkippedCount != 1 {
		t.Fatalf("unexpected scoped replay result: %#v", result)
	}
}
