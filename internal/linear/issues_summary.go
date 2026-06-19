package linear

import "strings"

func SummarizeOperation(result OperationResult) OperationSummary {
	summary := OperationSummary{
		Mode:               result.Mode,
		Synced:             result.Synced,
		QueuePath:          privateRelativePath(result.QueuePath),
		HTTPStatus:         result.HTTPStatus,
		RateLimitRemaining: result.RateLimitRemaining,
		Message:            result.Message,
		IssueCount:         len(result.Issues),
	}
	if len(result.Issues) > 0 {
		summary.Issues = summarizeIssues(result.Issues)
	}
	if result.Issue != nil {
		issue := summarizeIssue(*result.Issue)
		summary.Issue = &issue
	}
	if result.Comment != nil {
		summary.CommentCreated = true
		summary.CommentCreatedAt = result.Comment.CreatedAt
	}
	if result.State != nil {
		summary.StateType = strings.TrimSpace(result.State.Type)
	}
	return summary
}

func summarizeIssues(issues []Issue) []IssueSummary {
	summaries := make([]IssueSummary, 0, len(issues))
	for _, issue := range issues {
		summaries = append(summaries, summarizeIssue(issue))
	}
	return summaries
}

func summarizeIssue(issue Issue) IssueSummary {
	return IssueSummary{
		Identifier: strings.TrimSpace(issue.Identifier),
		Title:      strings.TrimSpace(issue.Title),
		UpdatedAt:  strings.TrimSpace(issue.UpdatedAt),
		StateType:  strings.TrimSpace(issue.State.Type),
	}
}
