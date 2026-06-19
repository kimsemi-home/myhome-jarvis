package linear

import (
	"context"
	"net/http"
	"strings"
)

func replayComment(
	ctx context.Context,
	root string,
	client *http.Client,
	token string,
	payload commentPayload,
) ReplayResult {
	issueID := strings.TrimSpace(payload.IssueID)
	body := strings.TrimSpace(payload.Body)
	if issueID == "" || body == "" {
		return failedReplay("Offline Linear comment payload is incomplete.")
	}
	var response struct {
		CommentCreate struct {
			Success bool     `json:"success"`
			Comment *Comment `json:"comment"`
		} `json:"commentCreate"`
	}
	query := `mutation AddComment($issueId: String!, $body: String!) { commentCreate(input: { issueId: $issueId, body: $body }) { success comment { id createdAt } } }`
	httpStatus, remaining, err := doGraphQL(
		ctx,
		client,
		token,
		query,
		map[string]string{"issueId": issueID, "body": body},
		&response,
	)
	if err != nil || !response.CommentCreate.Success {
		return ReplayResult{
			Mode:               "online",
			Synced:             false,
			HTTPStatus:         httpStatus,
			RateLimitRemaining: remaining,
			Message:            "Offline Linear comment replay failed; entry remains synced=false.",
		}
	}
	_ = AppendWriteEvidence(root, offlineReplayCommentKind, issueID)
	return ReplayResult{
		Mode:               "online",
		Synced:             true,
		HTTPStatus:         httpStatus,
		RateLimitRemaining: remaining,
		Message:            "Offline Linear comment replayed.",
	}
}
