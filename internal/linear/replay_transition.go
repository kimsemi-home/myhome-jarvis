package linear

import (
	"context"
	"net/http"
	"strings"
)

func replayTransition(
	ctx context.Context,
	root string,
	client *http.Client,
	token string,
	payload transitionPayload,
) ReplayResult {
	issueID := strings.TrimSpace(payload.IssueID)
	stateName := strings.TrimSpace(payload.State)
	if issueID == "" || stateName == "" {
		return failedReplay("Offline Linear transition payload is incomplete.")
	}
	stateID, httpStatus, remaining, err := findWorkflowStateID(ctx, client, token, stateName)
	if err != nil {
		return ReplayResult{
			Mode:               "online",
			Synced:             false,
			HTTPStatus:         httpStatus,
			RateLimitRemaining: remaining,
			Message:            "Offline Linear transition lookup failed; entry remains synced=false.",
		}
	}
	if rateLimitLow(remaining) {
		return ReplayResult{
			Mode:               "online",
			Synced:             false,
			HTTPStatus:         httpStatus,
			RateLimitRemaining: remaining,
			RateLimited:        true,
			Message:            "Linear rate limit remaining is low; offline replay paused before transition mutation.",
		}
	}
	return replayTransitionMutation(ctx, root, client, token, issueID, stateID)
}

func replayTransitionMutation(
	ctx context.Context,
	root string,
	client *http.Client,
	token string,
	issueID string,
	stateID string,
) ReplayResult {
	var response struct {
		IssueUpdate struct {
			Success bool   `json:"success"`
			Issue   *Issue `json:"issue"`
		} `json:"issueUpdate"`
	}
	query := `mutation TransitionIssue($issueId: String!, $stateId: String!) { issueUpdate(id: $issueId, input: { stateId: $stateId }) { success issue { id identifier title state { id name type } } } }`
	httpStatus, remaining, err := doGraphQL(
		ctx,
		client,
		token,
		query,
		map[string]string{"issueId": issueID, "stateId": stateID},
		&response,
	)
	return transitionMutationResult(root, issueID, httpStatus, remaining, err, response)
}
