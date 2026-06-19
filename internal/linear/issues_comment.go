package linear

import (
	"context"
	"net/http"
	"strings"
)

func AddComment(ctx context.Context, root string, client *http.Client, issueID string, body string) OperationResult {
	issueID = strings.TrimSpace(issueID)
	body = strings.TrimSpace(body)
	payload := map[string]string{"issue_id": issueID, "body": body}
	result := baseOperationResult(root)
	if issueID == "" || body == "" {
		result.Message = "Issue id and comment body are required."
		_ = AppendOfflineAction(root, "linear_comment_invalid", result.Message, payload)
		return result
	}
	token, err := loadToken(root)
	if err != nil {
		result.Message = "No Linear token found. Comment queued offline with synced=false."
		_ = AppendOfflineAction(root, "linear_comment", result.Message, payload)
		return result
	}
	return createLinearComment(ctx, root, client, token.Value, payload)
}

func createLinearComment(ctx context.Context, root string, client *http.Client, token string, payload map[string]string) OperationResult {
	result := baseOperationResult(root)
	var response struct {
		CommentCreate struct {
			Success bool     `json:"success"`
			Comment *Comment `json:"comment"`
		} `json:"commentCreate"`
	}
	query := `mutation AddComment($issueId: String!, $body: String!) { commentCreate(input: { issueId: $issueId, body: $body }) { success comment { id createdAt } } }`
	variables := map[string]string{"issueId": payload["issue_id"], "body": payload["body"]}
	status, remaining, err := doGraphQL(ctx, client, token, query, variables, &response)
	result.HTTPStatus = status
	result.RateLimitRemaining = remaining
	if err != nil || !response.CommentCreate.Success {
		result.Message = "Linear comment failed; queued offline with synced=false."
		if err != nil {
			result.Message += " " + err.Error()
		}
		_ = AppendOfflineAction(root, "linear_comment", result.Message, payload)
		return result
	}
	result.Mode = "online"
	result.Synced = true
	result.Comment = response.CommentCreate.Comment
	result.Message = "Linear comment created."
	if err := AppendWriteEvidence(root, "linear_comment", payload["issue_id"]); err != nil {
		result.Message = "Linear comment created; private write evidence was not recorded."
	}
	return result
}
