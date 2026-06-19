package linear

import (
	"context"
	"fmt"
	"net/http"
)

func CreateFromBacklog(ctx context.Context, root string, client *http.Client) OperationResult {
	result := baseOperationResult(root)
	seeds := backlogSeeds()
	token, err := loadToken(root)
	if err != nil {
		queueBacklogOffline(root, &result, "No Linear token found. Backlog seed queued offline with synced=false.", seeds)
		return result
	}
	teamID, ok := resolveBacklogTeamID(ctx, root, client, token.Value, &result, seeds)
	if !ok {
		return result
	}
	existingTitles, status, remaining, err := listIssueTitles(ctx, client, token.Value)
	result.HTTPStatus = status
	result.RateLimitRemaining = remaining
	if err != nil {
		message := "Linear issue title lookup failed; backlog seed queued offline with synced=false. " + err.Error()
		queueBacklogOffline(root, &result, message, seeds)
		return result
	}
	pendingSeeds := missingBacklogSeeds(seeds, existingTitles)
	if len(pendingSeeds) == 0 {
		result.Mode = "online"
		result.Synced = true
		result.Message = fmt.Sprintf("Created 0 Linear backlog seed issues; skipped %d existing seeds.", len(seeds))
		return result
	}
	return createMissingBacklogSeeds(ctx, root, client, token.Value, teamID, seeds, pendingSeeds)
}

func queueBacklogOffline(root string, result *OperationResult, message string, seeds []backlogSeed) {
	result.Message = message
	payload := map[string]any{"issues": seeds}
	_ = AppendOfflineAction(root, "linear_create_from_backlog", result.Message, payload)
}
