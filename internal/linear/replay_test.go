package linear

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestReplayOfflineReplaysWriteSafeActionsOnce(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	root := t.TempDir()
	if err := AppendOfflineAction(root, "linear_comment", "queued", map[string]string{"issue_id": "KIM-16", "body": "Started work"}); err != nil {
		t.Fatal(err)
	}
	if err := AppendOfflineAction(root, "linear_transition", "queued", map[string]string{"issue_id": "KIM-16", "state": "Done"}); err != nil {
		t.Fatal(err)
	}
	if err := AppendOfflineAction(root, "linear_create_from_backlog", "queued", map[string]any{"raw": "not replay-safe"}); err != nil {
		t.Fatal(err)
	}

	requests := 0
	client := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		requests++
		var body struct {
			Query     string            `json:"query"`
			Variables map[string]string `json:"variables"`
		}
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		switch {
		case strings.Contains(body.Query, "commentCreate"):
			if body.Variables["issueId"] != "KIM-16" || body.Variables["body"] != "Started work" {
				t.Fatalf("unexpected comment variables: %#v", body.Variables)
			}
			return linearResponse(250, `{"data":{"commentCreate":{"success":true,"comment":{"id":"comment-id","createdAt":"2026-06-14T00:00:00.000Z"}}}}`), nil
		case strings.Contains(body.Query, "query WorkflowStates"):
			return linearResponse(249, `{"data":{"workflowStates":{"nodes":[{"id":"done-state","name":"Done","type":"completed"}]}}}`), nil
		case strings.Contains(body.Query, "mutation TransitionIssue"):
			if body.Variables["issueId"] != "KIM-16" || body.Variables["stateId"] != "done-state" {
				t.Fatalf("unexpected transition variables: %#v", body.Variables)
			}
			return linearResponse(248, `{"data":{"issueUpdate":{"success":true,"issue":{"id":"issue-id","identifier":"KIM-16","title":"Replay","state":{"id":"done-state","name":"Done","type":"completed"}}}}}`), nil
		default:
			t.Fatalf("unexpected GraphQL query: %s", body.Query)
			return nil, nil
		}
	})}

	result := ReplayOffline(context.Background(), root, client)
	if !result.Synced || result.ReplayedCount != 2 || result.EligibleCount != 2 || result.SkippedCount != 1 {
		t.Fatalf("unexpected replay result: %#v", result)
	}
	if requests != 3 {
		t.Fatalf("requests = %d", requests)
	}
	replayData, err := os.ReadFile(filepath.Join(root, OfflineReplayRelativePath))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(replayData), "Started work") || strings.Contains(string(replayData), "not replay-safe") {
		t.Fatalf("replay evidence leaked raw payload: %s", string(replayData))
	}
	evidence, err := WriteEvidenceStatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if evidence.SyncedMutationCount != 2 || evidence.LatestSyncedMutation == nil || evidence.LatestSyncedMutation.IssueKey != "KIM-16" {
		t.Fatalf("write evidence = %#v", evidence)
	}

	secondClient := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		t.Fatalf("already replayed entries should not call Linear again")
		return nil, nil
	})}
	second := ReplayOffline(context.Background(), root, secondClient)
	if !second.Synced || second.ReplayedCount != 0 || second.AlreadyReplayedCount != 2 {
		t.Fatalf("unexpected second replay result: %#v", second)
	}
}

func TestReplayOfflinePausesWhenRateLimitIsLow(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	root := t.TempDir()
	for _, issueID := range []string{"KIM-16", "KIM-17"} {
		if err := AppendOfflineAction(root, "linear_comment", "queued", map[string]string{"issue_id": issueID, "body": "status"}); err != nil {
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

func TestReplayOfflineSkipsOutOfScopeIssueKeys(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	t.Setenv("LINEAR_TEAM_KEY", "KIM")
	root := t.TempDir()
	if err := AppendOfflineAction(root, "linear_comment", "queued", map[string]string{"issue_id": "MHJ-1", "body": "old team"}); err != nil {
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

func TestReplayOfflineFailedEntryRemainsQueued(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	root := t.TempDir()
	if err := AppendOfflineAction(root, "linear_comment", "queued", map[string]string{"issue_id": "KIM-16", "body": "status"}); err != nil {
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

func linearResponse(remaining int, body string) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"X-RateLimit-Remaining": []string{strconv.Itoa(remaining)}},
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}
