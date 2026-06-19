package linear

import (
	"context"
	"testing"
)

func nextIssueFromNodes(t *testing.T, nodes ...string) OperationResult {
	t.Helper()
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	t.Setenv("LINEAR_TEAM_ID", "")
	t.Setenv("LINEAR_TEAM_KEY", "KIM")
	client := linearGraphQLClient(t, "", func(linearGraphQLBody) string {
		return issuesBody(nodes...)
	})
	return NextIssue(context.Background(), t.TempDir(), client)
}

func TestNextIssuePrefersProjectIssue(t *testing.T) {
	result := nextIssueFromNodes(t,
		issueNode("onboarding-id", "KIM-1", "Get familiar with Linear", "team-kim", "KIM", "todo", "Todo", "unstarted"),
		issueNode("project-id", "KIM-7", "[myhome-jarvis] Prefer project Linear issues", "team-kim", "KIM", "started", "In Progress", "started"),
	)
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
	result := nextIssueFromNodes(t,
		issueNode("backlog-id", "KIM-13", "[myhome-jarvis] Include project queue status", "team-kim", "KIM", "backlog", "Backlog", "backlog"),
		issueNode("started-id", "KIM-14", "[myhome-jarvis] Add DDD SSOT", "team-kim", "KIM", "started", "In Progress", "started"),
	)
	if !result.Synced || result.Issue == nil {
		t.Fatalf("unexpected next result: %#v", result)
	}
	if result.Issue.Identifier != "KIM-14" {
		t.Fatalf("next issue = %s, expected KIM-14", result.Issue.Identifier)
	}
}

func TestNextIssueDoesNotSelectUnrelatedActiveIssue(t *testing.T) {
	result := nextIssueFromNodes(t,
		issueNode("onboarding-id", "KIM-1", "Get familiar with Linear", "team-kim", "KIM", "todo", "Todo", "unstarted"),
		issueNode("tools-id", "KIM-3", "Connect your tools", "team-kim", "KIM", "todo", "Todo", "unstarted"),
	)
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
