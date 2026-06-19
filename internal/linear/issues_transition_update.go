package linear

import (
	"context"
	"net/http"
)

func updateLinearIssueState(ctx context.Context, root string, client *http.Client, token string, issueID string, stateID string, payload map[string]string) OperationResult {
	result := baseOperationResult(root)
	var response struct {
		IssueUpdate struct {
			Success bool   `json:"success"`
			Issue   *Issue `json:"issue"`
		} `json:"issueUpdate"`
	}
	query := `mutation TransitionIssue($issueId: String!, $stateId: String!) { issueUpdate(id: $issueId, input: { stateId: $stateId }) { success issue { id identifier title state { id name type } } } }`
	status, remaining, err := doGraphQL(ctx, client, token, query, map[string]string{"issueId": issueID, "stateId": stateID}, &response)
	result.HTTPStatus = status
	result.RateLimitRemaining = remaining
	if err != nil || !response.IssueUpdate.Success {
		result.Message = "Linear transition failed; queued offline with synced=false."
		if err != nil {
			result.Message += " " + err.Error()
		}
		_ = AppendOfflineAction(root, "linear_transition", result.Message, payload)
		return result
	}
	result.Mode = "online"
	result.Synced = true
	result.Issue = response.IssueUpdate.Issue
	if response.IssueUpdate.Issue != nil {
		result.State = &response.IssueUpdate.Issue.State
	}
	result.Message = "Linear issue transitioned."
	recordTransitionWriteEvidence(root, &result, issueID)
	return result
}

func recordTransitionWriteEvidence(root string, result *OperationResult, issueID string) {
	evidenceIssueKey := issueID
	if result.Issue != nil {
		evidenceIssueKey = result.Issue.Identifier
	}
	if err := AppendWriteEvidence(root, "linear_transition", evidenceIssueKey); err != nil {
		result.Message = "Linear issue transitioned; private write evidence was not recorded."
	}
}
