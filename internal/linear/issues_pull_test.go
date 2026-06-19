package linear

import (
	"context"
	"strings"
	"testing"
)

func TestPullIssuesUsesDirectGraphQL(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	t.Setenv("LINEAR_TEAM_ID", "")
	t.Setenv("LINEAR_TEAM_KEY", "")
	client := linearGraphQLClient(t, "4998", func(body linearGraphQLBody) string {
		if !strings.Contains(body.Query, "query PullIssues") {
			t.Fatalf("unexpected query: %s", body.Query)
		}
		requireBodyContains(t, body.Query, "issues(first: 50)", "team { id key }", "state { id name type }")
		requireBodyOmits(t, body.Query, "description", "url", "team { id name }")
		return issuesBody(
			issueNode("issue-id", "MHJ-1", "Build local daemon", "team-id", "MHJ", "state-id", "Todo", "unstarted"),
			issueNode("done-id", "MHJ-2", "Done issue", "team-id", "MHJ", "done-state-id", "Done", "completed"),
		)
	})

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
	client := linearGraphQLClient(t, "", func(linearGraphQLBody) string {
		return issuesBody(
			issueNode("wanted-id", "KIM-6", "Scope Linear pull", "team-kim", "KIM", "todo", "Todo", "unstarted"),
			issueNode("other-id", "OPS-1", "Other team", "team-ops", "OPS", "todo", "Todo", "unstarted"),
			issueNode("done-id", "KIM-5", "Completed team issue", "team-kim", "KIM", "done", "Done", "completed"),
		)
	})

	result := PullIssues(context.Background(), t.TempDir(), client)
	if !result.Synced || len(result.Issues) != 1 {
		t.Fatalf("unexpected filtered result: %#v", result)
	}
	if result.Issues[0].Identifier != "KIM-6" {
		t.Fatalf("selected issue = %s, expected KIM-6", result.Issues[0].Identifier)
	}
}
