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
	t.Setenv("LINEAR_TEAM_ID", "")
	t.Setenv("LINEAR_TEAM_KEY", "")
	client := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			t.Fatal(err)
		}
		bodyText := string(body)
		if !strings.Contains(bodyText, "query PullIssues") {
			t.Fatalf("unexpected body: %s", bodyText)
		}
		for _, expected := range []string{"issues(first: 50)", "team { id key }", "state { id name type }"} {
			if !strings.Contains(bodyText, expected) {
				t.Fatalf("expected %s in %s", expected, bodyText)
			}
		}
		for _, forbidden := range []string{"description", "url", "team { id name }"} {
			if strings.Contains(bodyText, forbidden) {
				t.Fatalf("pull query requested %s in %s", forbidden, bodyText)
			}
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
							"team": {"id": "team-id", "key": "MHJ"},
							"state": {"id": "state-id", "name": "Todo", "type": "unstarted"}
						}, {
							"id": "done-id",
							"identifier": "MHJ-2",
							"title": "Done issue",
							"updatedAt": "2026-06-14T00:00:00.000Z",
							"team": {"id": "team-id", "key": "MHJ"},
							"state": {"id": "done-state-id", "name": "Done", "type": "completed"}
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

func TestPullIssuesFiltersConfiguredTeamKeyAndOpenStates(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	t.Setenv("LINEAR_TEAM_ID", "")
	t.Setenv("LINEAR_TEAM_KEY", "KIM")
	client := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{},
			Body: io.NopCloser(strings.NewReader(`{
				"data": {
					"issues": {
						"nodes": [{
							"id": "wanted-id",
							"identifier": "KIM-6",
							"title": "Scope Linear pull",
							"updatedAt": "2026-06-14T00:00:00.000Z",
							"team": {"id": "team-kim", "key": "KIM"},
							"state": {"id": "todo-state", "name": "Todo", "type": "unstarted"}
						}, {
							"id": "other-team-id",
							"identifier": "OPS-1",
							"title": "Other team",
							"updatedAt": "2026-06-14T00:00:00.000Z",
							"team": {"id": "team-ops", "key": "OPS"},
							"state": {"id": "todo-state", "name": "Todo", "type": "unstarted"}
						}, {
							"id": "done-id",
							"identifier": "KIM-5",
							"title": "Completed team issue",
							"updatedAt": "2026-06-14T00:00:00.000Z",
							"team": {"id": "team-kim", "key": "KIM"},
							"state": {"id": "done-state", "name": "Done", "type": "completed"}
						}]
					}
				}
			}`)),
		}, nil
	})}

	result := PullIssues(context.Background(), t.TempDir(), client)
	if !result.Synced || len(result.Issues) != 1 {
		t.Fatalf("unexpected filtered result: %#v", result)
	}
	if result.Issues[0].Identifier != "KIM-6" {
		t.Fatalf("selected issue = %s, expected KIM-6", result.Issues[0].Identifier)
	}
}

func TestFilterActiveIssuesFiltersConfiguredTeamID(t *testing.T) {
	issues := []Issue{
		{
			Identifier: "KIM-6",
			Team:       TeamStatus{ID: "team-kim", Key: "KIM"},
			State:      StateStatus{Name: "Todo", Type: "unstarted"},
		},
		{
			Identifier: "OPS-1",
			Team:       TeamStatus{ID: "team-ops", Key: "OPS"},
			State:      StateStatus{Name: "Todo", Type: "unstarted"},
		},
	}

	filtered := filterActiveIssues(issues, issueScope{TeamID: "team-kim"})
	if len(filtered) != 1 || filtered[0].Identifier != "KIM-6" {
		t.Fatalf("filtered issues = %#v", filtered)
	}
}

func TestNextIssuePrefersProjectIssue(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	t.Setenv("LINEAR_TEAM_ID", "")
	t.Setenv("LINEAR_TEAM_KEY", "KIM")
	client := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{},
			Body: io.NopCloser(strings.NewReader(`{
				"data": {
					"issues": {
						"nodes": [{
							"id": "onboarding-id",
							"identifier": "KIM-1",
							"title": "Get familiar with Linear",
							"updatedAt": "2026-06-14T00:00:00.000Z",
							"team": {"id": "team-kim", "key": "KIM"},
							"state": {"id": "todo-state", "name": "Todo", "type": "unstarted"}
						}, {
							"id": "project-id",
							"identifier": "KIM-7",
							"title": "[myhome-jarvis] Prefer project Linear issues in next selection",
							"updatedAt": "2026-06-14T00:00:00.000Z",
							"team": {"id": "team-kim", "key": "KIM"},
							"state": {"id": "started-state", "name": "In Progress", "type": "started"}
						}]
					}
				}
			}`)),
		}, nil
	})}

	result := NextIssue(context.Background(), t.TempDir(), client)
	if !result.Synced || result.Issue == nil {
		t.Fatalf("unexpected next result: %#v", result)
	}
	if result.Issue.Identifier != "KIM-7" {
		t.Fatalf("next issue = %s, expected KIM-7", result.Issue.Identifier)
	}
	if result.Message != "Selected next project Linear issue." {
		t.Fatalf("message = %q", result.Message)
	}
}

func TestNextIssuePrefersStartedProjectIssue(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	t.Setenv("LINEAR_TEAM_ID", "")
	t.Setenv("LINEAR_TEAM_KEY", "KIM")
	client := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{},
			Body: io.NopCloser(strings.NewReader(`{
				"data": {
					"issues": {
						"nodes": [{
							"id": "backlog-id",
							"identifier": "KIM-13",
							"title": "[myhome-jarvis] Include project queue status in loop checkpoints",
							"updatedAt": "2026-06-14T18:19:53.236Z",
							"team": {"id": "team-kim", "key": "KIM"},
							"state": {"id": "backlog-state", "name": "Backlog", "type": "backlog"}
						}, {
							"id": "started-id",
							"identifier": "KIM-14",
							"title": "[myhome-jarvis] Add DDD SSOT and local KnowledgeIndex thin slice",
							"updatedAt": "2026-06-14T18:17:07.010Z",
							"team": {"id": "team-kim", "key": "KIM"},
							"state": {"id": "started-state", "name": "In Progress", "type": "started"}
						}]
					}
				}
			}`)),
		}, nil
	})}

	result := NextIssue(context.Background(), t.TempDir(), client)
	if !result.Synced || result.Issue == nil {
		t.Fatalf("unexpected next result: %#v", result)
	}
	if result.Issue.Identifier != "KIM-14" {
		t.Fatalf("next issue = %s, expected KIM-14", result.Issue.Identifier)
	}
}

func TestNextIssueDoesNotSelectUnrelatedActiveIssue(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	t.Setenv("LINEAR_TEAM_ID", "")
	t.Setenv("LINEAR_TEAM_KEY", "KIM")
	client := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{},
			Body: io.NopCloser(strings.NewReader(`{
				"data": {
					"issues": {
						"nodes": [{
							"id": "onboarding-id",
							"identifier": "KIM-1",
							"title": "Get familiar with Linear",
							"updatedAt": "2026-06-14T00:00:00.000Z",
							"team": {"id": "team-kim", "key": "KIM"},
							"state": {"id": "todo-state", "name": "Todo", "type": "unstarted"}
						}, {
							"id": "tools-id",
							"identifier": "KIM-3",
							"title": "Connect your tools",
							"updatedAt": "2026-06-14T00:00:00.000Z",
							"team": {"id": "team-kim", "key": "KIM"},
							"state": {"id": "todo-state", "name": "Todo", "type": "unstarted"}
						}]
					}
				}
			}`)),
		}, nil
	})}

	result := NextIssue(context.Background(), t.TempDir(), client)
	if !result.Synced || len(result.Issues) != 2 {
		t.Fatalf("unexpected next result: %#v", result)
	}
	if result.Issue != nil {
		t.Fatalf("unexpected selected issue: %#v", result.Issue)
	}
	expected := "Pulled active Linear issues, but none matched the project issue prefix."
	if result.Message != expected {
		t.Fatalf("message = %q, expected %q", result.Message, expected)
	}
}

func TestProjectIssueTitlePrefixMatchesGeneratedPolicy(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "..", "generated", "linear.generated.json"))
	if err != nil {
		t.Fatal(err)
	}
	var policy struct {
		ProjectIssueTitlePrefix    string   `json:"project_issue_title_prefix"`
		NextPrefersProjectIssue    bool     `json:"next_prefers_project_issues"`
		NextRequiresProjectIssue   bool     `json:"next_requires_project_issue"`
		BacklogSeedDedupesByTitle  bool     `json:"backlog_seed_dedupes_by_title"`
		BacklogSeedCurrentProject  bool     `json:"backlog_seed_current_project_only"`
		BacklogSeedQueriesExisting bool     `json:"backlog_seed_queries_existing_titles"`
		OfflineReplayEvidence      string   `json:"offline_replay_evidence"`
		OfflineReplayRateFloor     int      `json:"offline_replay_rate_limit_floor"`
		Commands                   []string `json:"commands"`
		OfflineReplaySafeKinds     []string `json:"offline_replay_safe_action_kinds"`
	}
	if err := json.Unmarshal(data, &policy); err != nil {
		t.Fatal(err)
	}
	if policy.ProjectIssueTitlePrefix != projectIssueTitlePrefix {
		t.Fatalf("project prefix = %q, expected %q", policy.ProjectIssueTitlePrefix, projectIssueTitlePrefix)
	}
	if !policy.NextPrefersProjectIssue {
		t.Fatal("generated policy must keep next_prefers_project_issues enabled")
	}
	if !policy.NextRequiresProjectIssue {
		t.Fatal("generated policy must keep next_requires_project_issue enabled")
	}
	if !policy.BacklogSeedDedupesByTitle {
		t.Fatal("generated policy must keep backlog_seed_dedupes_by_title enabled")
	}
	if !policy.BacklogSeedCurrentProject {
		t.Fatal("generated policy must keep backlog_seed_current_project_only enabled")
	}
	if !policy.BacklogSeedQueriesExisting {
		t.Fatal("generated policy must keep backlog_seed_queries_existing_titles enabled")
	}
	if policy.OfflineReplayEvidence != OfflineReplayRelativePath {
		t.Fatalf("offline replay evidence path = %q, expected %q", policy.OfflineReplayEvidence, OfflineReplayRelativePath)
	}
	if policy.OfflineReplayRateFloor != defaultReplayRateLimitFloor {
		t.Fatalf("offline replay rate floor = %d, expected %d", policy.OfflineReplayRateFloor, defaultReplayRateLimitFloor)
	}
	if !containsString(policy.Commands, "mhj linear replay-offline") {
		t.Fatalf("generated commands missing replay-offline: %#v", policy.Commands)
	}
	for _, kind := range []string{offlineReplayCommentKind, offlineReplayTransitionKind} {
		if !containsString(policy.OfflineReplaySafeKinds, kind) {
			t.Fatalf("generated safe replay kinds missing %s: %#v", kind, policy.OfflineReplaySafeKinds)
		}
	}
}

func TestCreateFromBacklogSkipsExistingSeedTitles(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	t.Setenv("LINEAR_TEAM_ID", "team-id")
	root := t.TempDir()
	seeds := backlogSeeds()
	missingTitle := seeds[1].Title
	requests := 0
	createdTitles := []string{}
	client := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		requests++
		var body struct {
			Query     string         `json:"query"`
			Variables map[string]any `json:"variables"`
		}
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		switch {
		case strings.Contains(body.Query, "query ExistingIssueTitles"):
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     http.Header{"X-RateLimit-Remaining": []string{"4000"}},
				Body: io.NopCloser(strings.NewReader(`{
					"data": {
						"issues": {
							"nodes": [{
								"title": "[myhome-jarvis] Track approved Linear write evidence"
							}, {
								"title": "[myhome-jarvis] Include project queue status in loop checkpoints"
							}]
						}
					}
				}`)),
			}, nil
		case strings.Contains(body.Query, "mutation IssueCreate"):
			title, ok := body.Variables["title"].(string)
			if !ok {
				t.Fatalf("title variable missing: %#v", body.Variables)
			}
			createdTitles = append(createdTitles, title)
			if title != missingTitle {
				t.Fatalf("created title = %q, expected %q", title, missingTitle)
			}
			if body.Variables["teamId"] != "team-id" || int(body.Variables["priority"].(float64)) != 3 {
				t.Fatalf("unexpected variables: %#v", body.Variables)
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     http.Header{"X-RateLimit-Remaining": []string{"3999"}},
				Body: io.NopCloser(strings.NewReader(`{
					"data": {
						"issueCreate": {
							"success": true,
							"issue": {
								"id": "issue-id",
								"identifier": "KIM-11",
								"title": "[myhome-jarvis] Reconcile planner external-write gate",
								"state": {"id": "state-id", "name": "Todo", "type": "unstarted"}
							}
						}
					}
				}`)),
			}, nil
		default:
			t.Fatalf("unexpected GraphQL request: %s", body.Query)
			return nil, nil
		}
	})}

	result := CreateFromBacklog(context.Background(), root, client)
	if !result.Synced || len(result.Issues) != 1 {
		t.Fatalf("unexpected result: %#v", result)
	}
	if requests != 2 || len(createdTitles) != 1 {
		t.Fatalf("requests=%d created=%#v", requests, createdTitles)
	}
	if !strings.Contains(result.Message, "Created 1") || !strings.Contains(result.Message, "skipped 2") {
		t.Fatalf("message did not include created/skipped counts: %q", result.Message)
	}
	status, err := WriteEvidenceStatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.SyncedMutationCount != 1 || status.LatestSyncedMutation == nil {
		t.Fatalf("write evidence status = %#v", status)
	}
	if status.LatestSyncedMutation.Action != "linear_create_from_backlog" || status.LatestSyncedMutation.IssueKey != "KIM-11" {
		t.Fatalf("write evidence = %#v", status.LatestSyncedMutation)
	}
}

func TestCreateFromBacklogSkipsAllExistingSeedTitles(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	t.Setenv("LINEAR_TEAM_ID", "team-id")
	client := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		var body struct {
			Query string `json:"query"`
		}
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(body.Query, "query ExistingIssueTitles") {
			t.Fatalf("unexpected GraphQL request: %s", body.Query)
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"X-RateLimit-Remaining": []string{"4000"}},
			Body: io.NopCloser(strings.NewReader(`{
				"data": {
					"issues": {
						"nodes": [{
							"title": "[myhome-jarvis] Track approved Linear write evidence"
						}, {
							"title": "[myhome-jarvis] Reconcile planner external-write gate"
						}, {
							"title": "[myhome-jarvis] Include project queue status in loop checkpoints"
						}]
					}
				}
			}`)),
		}, nil
	})}

	result := CreateFromBacklog(context.Background(), t.TempDir(), client)
	if !result.Synced || len(result.Issues) != 0 {
		t.Fatalf("unexpected result: %#v", result)
	}
	if result.RateLimitRemaining != 4000 {
		t.Fatalf("rate remaining = %d", result.RateLimitRemaining)
	}
	if result.Message != "Created 0 Linear backlog seed issues; skipped 3 existing seeds." {
		t.Fatalf("message = %q", result.Message)
	}
}

func TestOperationSummaryRedactsLinearIssueDetails(t *testing.T) {
	result := OperationResult{
		Mode:               "online",
		Synced:             true,
		QueuePath:          filepath.Join(t.TempDir(), "data", "private", "linear-offline-queue.jsonl"),
		HTTPStatus:         http.StatusOK,
		RateLimitRemaining: 42,
		Message:            "Selected next open Linear issue.",
		Issues: []Issue{{
			ID:          "issue-id",
			Identifier:  "MHJ-1",
			Title:       "Build local daemon",
			Description: "raw acceptance text",
			URL:         "https://linear.app/private/issue/MHJ-1",
			UpdatedAt:   "2026-06-14T00:00:00.000Z",
			Team:        TeamStatus{ID: "team-id", Name: "Private Team"},
			State:       StateStatus{ID: "state-id", Name: "Todo", Type: "unstarted"},
		}},
	}
	result.Issue = &result.Issues[0]

	payload, err := json.Marshal(SummarizeOperation(result))
	if err != nil {
		t.Fatal(err)
	}
	body := string(payload)
	for _, expected := range []string{
		`"queue_path":"data/private/linear-offline-queue.jsonl"`,
		`"identifier":"MHJ-1"`,
		`"state_type":"unstarted"`,
	} {
		if !strings.Contains(body, expected) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{
		`description`,
		`url`,
		`team`,
		`issue-id`,
		`team-id`,
		`state-id`,
		`Private Team`,
		`raw acceptance text`,
		`"queue_path":"/`,
		`"queue_path":"\\`,
	} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("summary leaked %s in %s", forbidden, body)
		}
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
	status, err := WriteEvidenceStatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.HasSyncedMutation || status.SyncedMutationCount != 0 {
		t.Fatalf("offline comment should not create write evidence: %#v", status)
	}
}

func TestAddCommentUsesVariables(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	commentBody := "Line one\nLine two with \"quotes\""
	root := t.TempDir()
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

	result := AddComment(context.Background(), root, client, "MHJ-1", commentBody)
	if !result.Synced || result.Comment == nil || result.Comment.ID != "comment-id" {
		t.Fatalf("unexpected result: %#v", result)
	}
	status, err := WriteEvidenceStatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.SyncedMutationCount != 1 || status.LatestSyncedMutation == nil {
		t.Fatalf("write evidence status = %#v", status)
	}
	if status.LatestSyncedMutation.Action != "linear_comment" || status.LatestSyncedMutation.IssueKey != "MHJ-1" || !status.LatestSyncedMutation.Synced {
		t.Fatalf("write evidence = %#v", status.LatestSyncedMutation)
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

func TestTransitionRecordsApprovedWriteEvidence(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	root := t.TempDir()
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
		case strings.Contains(body.Query, "query WorkflowStates"):
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     http.Header{},
				Body: io.NopCloser(strings.NewReader(`{
					"data": {
						"workflowStates": {
							"nodes": [{
								"id": "state-id",
								"name": "Done",
								"type": "completed"
							}]
						}
					}
				}`)),
			}, nil
		case strings.Contains(body.Query, "mutation TransitionIssue"):
			if body.Variables["issueId"] != "MHJ-1" || body.Variables["stateId"] != "state-id" {
				t.Fatalf("unexpected variables: %#v", body.Variables)
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     http.Header{},
				Body: io.NopCloser(strings.NewReader(`{
					"data": {
						"issueUpdate": {
							"success": true,
							"issue": {
								"id": "issue-id",
								"identifier": "MHJ-1",
								"title": "Done issue",
								"state": {"id": "state-id", "name": "Done", "type": "completed"}
							}
						}
					}
				}`)),
			}, nil
		default:
			t.Fatalf("unexpected GraphQL request: %s", body.Query)
			return nil, nil
		}
	})}

	result := TransitionIssue(context.Background(), root, client, "MHJ-1", "Done")
	if !result.Synced || result.State == nil || result.State.Type != "completed" {
		t.Fatalf("unexpected result: %#v", result)
	}
	if requests != 2 {
		t.Fatalf("requests = %d", requests)
	}
	status, err := WriteEvidenceStatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.SyncedMutationCount != 1 || status.LatestSyncedMutation == nil {
		t.Fatalf("write evidence status = %#v", status)
	}
	if status.LatestSyncedMutation.Action != "linear_transition" || status.LatestSyncedMutation.IssueKey != "MHJ-1" {
		t.Fatalf("write evidence = %#v", status.LatestSyncedMutation)
	}
}

func TestWriteEvidenceRedactsNonIssueKeys(t *testing.T) {
	root := t.TempDir()
	if err := AppendWriteEvidence(root, "linear_comment", "550e8400-e29b-41d4-a716-446655440000"); err != nil {
		t.Fatal(err)
	}
	status, err := WriteEvidenceStatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.LatestSyncedMutation == nil {
		t.Fatalf("expected latest write evidence: %#v", status)
	}
	if status.LatestSyncedMutation.IssueKey != "" {
		t.Fatalf("non-issue key leaked into evidence: %#v", status.LatestSyncedMutation)
	}
	payload, err := os.ReadFile(filepath.Join(root, WriteEvidenceRelativePath))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(payload), "550e8400") {
		t.Fatalf("raw id leaked into evidence file: %s", string(payload))
	}
}

func containsString(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}
