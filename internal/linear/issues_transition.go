package linear

import (
	"context"
	"net/http"
	"strings"
)

func TransitionIssue(ctx context.Context, root string, client *http.Client, issueID string, stateName string) OperationResult {
	issueID = strings.TrimSpace(issueID)
	stateName = strings.TrimSpace(stateName)
	payload := map[string]string{"issue_id": issueID, "state": stateName}
	result := baseOperationResult(root)
	if issueID == "" || stateName == "" {
		result.Message = "Issue id and target state are required."
		_ = AppendOfflineAction(root, "linear_transition_invalid", result.Message, payload)
		return result
	}
	token, err := loadToken(root)
	if err != nil {
		result.Message = "No Linear token found. Transition queued offline with synced=false."
		_ = AppendOfflineAction(root, "linear_transition", result.Message, payload)
		return result
	}

	stateID, status, remaining, err := findWorkflowStateID(ctx, client, token.Value, stateName)
	result.HTTPStatus = status
	result.RateLimitRemaining = remaining
	if err != nil {
		message := "Linear transition lookup failed; queued offline with synced=false. " + err.Error()
		queueTransitionOffline(root, &result, message, payload)
		return result
	}
	return updateLinearIssueState(ctx, root, client, token.Value, issueID, stateID, payload)
}

func queueTransitionOffline(root string, result *OperationResult, message string, payload map[string]string) {
	result.Message = message
	_ = AppendOfflineAction(root, "linear_transition", result.Message, payload)
}
