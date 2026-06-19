package linear

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

func appendReplayHappyPathQueue(t *testing.T, root string) {
	t.Helper()
	if err := AppendOfflineAction(root, "linear_comment", "queued", map[string]string{
		"issue_id": "KIM-16",
		"body":     "Started work",
	}); err != nil {
		t.Fatal(err)
	}
	if err := AppendOfflineAction(root, "linear_transition", "queued", map[string]string{
		"issue_id": "KIM-16",
		"state":    "Done",
	}); err != nil {
		t.Fatal(err)
	}
	if err := AppendOfflineAction(root, "linear_create_from_backlog", "queued", map[string]any{
		"raw": "not replay-safe",
	}); err != nil {
		t.Fatal(err)
	}
}

func replayHappyPathClient(t *testing.T, requests *int) *http.Client {
	t.Helper()
	return &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		*requests++
		var body struct {
			Query     string            `json:"query"`
			Variables map[string]string `json:"variables"`
		}
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		switch {
		case strings.Contains(body.Query, "commentCreate"):
			return commentReplayResponse(t, body.Variables), nil
		case strings.Contains(body.Query, "query WorkflowStates"):
			return linearResponse(249, `{"data":{"workflowStates":{"nodes":[{"id":"done-state","name":"Done","type":"completed"}]}}}`), nil
		case strings.Contains(body.Query, "mutation TransitionIssue"):
			return transitionReplayResponse(t, body.Variables), nil
		default:
			t.Fatalf("unexpected GraphQL query: %s", body.Query)
			return nil, nil
		}
	})}
}

func replayNoCallClient(t *testing.T) *http.Client {
	t.Helper()
	return &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		t.Fatalf("already replayed entries should not call Linear again")
		return nil, nil
	})}
}
