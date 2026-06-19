package linear

import (
	"context"
	"net/http"
)

func NextIssue(ctx context.Context, root string, client *http.Client) OperationResult {
	result := PullIssues(ctx, root, client)
	if !result.Synced || len(result.Issues) == 0 {
		if result.Message == "" {
			result.Message = "No Linear issue is available."
		}
		return result
	}
	if issue := selectNextIssue(result.Issues, true); issue != nil {
		result.Issue = issue
		result.Message = "Selected next project Linear issue."
		return result
	}
	if issue := selectNextIssue(result.Issues, false); issue != nil {
		result.Issue = issue
		result.Message = "Selected next project Linear issue."
		return result
	}
	result.Message = "Pulled active Linear issues, but none matched the project issue prefix."
	return result
}

func selectNextIssue(issues []Issue, requireStarted bool) *Issue {
	for index := range issues {
		issue := issues[index]
		if !isOpenState(issue.State) || !isProjectIssue(issue) {
			continue
		}
		if requireStarted && !isStartedState(issue.State) {
			continue
		}
		return &issue
	}
	return nil
}
