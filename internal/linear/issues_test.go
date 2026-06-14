package linear

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPullIssuesUsesDirectGraphQL(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	client := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(body), "query PullIssues") {
			t.Fatalf("unexpected body: %s", string(body))
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"X-RateLimit-Remaining": []string{"4998"}},
			Body: io.NopCloser(strings.NewReader(`{
				"data": {
					"issues": {
						"nodes": [{
							"id": "issue-id",
							"identifier": "MHJ-1",
							"title": "Build local daemon",
							"description": "Acceptance text",
							"url": "https://linear.app/example/issue/MHJ-1",
							"updatedAt": "2026-06-14T00:00:00.000Z",
							"team": {"id": "team-id", "name": "Home"},
							"state": {"id": "state-id", "name": "Todo", "type": "unstarted"}
						}]
					}
				}
			}`)),
		}, nil
	})}

	result := PullIssues(context.Background(), t.TempDir(), client)
	if !result.Synced || len(result.Issues) != 1 {
		t.Fatalf("unexpected result: %#v", result)
	}
	if result.Issues[0].Identifier != "MHJ-1" || result.RateLimitRemaining != 4998 {
		t.Fatalf("unexpected issue/rate data: %#v", result)
	}
}

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
}

func TestAddCommentUsesVariables(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	commentBody := "Line one\nLine two with \"quotes\""
	client := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		var body struct {
			Query     string            `json:"query"`
			Variables map[string]string `json:"variables"`
		}
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(body.Query, "commentCreate") {
			t.Fatalf("expected commentCreate mutation, got %s", body.Query)
		}
		if body.Variables["body"] != commentBody || body.Variables["issueId"] != "MHJ-1" {
			t.Fatalf("unexpected variables: %#v", body.Variables)
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{},
			Body: io.NopCloser(strings.NewReader(`{
				"data": {
					"commentCreate": {
						"success": true,
						"comment": {"id": "comment-id", "body": "ok", "createdAt": "2026-06-14T00:00:00.000Z"}
					}
				}
			}`)),
		}, nil
	})}

	result := AddComment(context.Background(), t.TempDir(), client, "MHJ-1", commentBody)
	if !result.Synced || result.Comment == nil || result.Comment.ID != "comment-id" {
		t.Fatalf("unexpected result: %#v", result)
	}
}

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
