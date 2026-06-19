package linear

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

func resolveBacklogTeamID(ctx context.Context, root string, client *http.Client, token string, result *OperationResult, seeds []backlogSeed) (string, bool) {
	teamID := strings.TrimSpace(envLinearTeamID())
	if teamID != "" {
		return teamID, true
	}
	teams, status, remaining, err := listTeams(ctx, client, token)
	result.HTTPStatus = status
	result.RateLimitRemaining = remaining
	if err == nil && len(teams) > 0 {
		return teams[0].ID, true
	}
	message := "Linear team lookup failed; backlog seed queued offline with synced=false."
	if err != nil {
		message += " " + err.Error()
	}
	queueBacklogOffline(root, result, message, seeds)
	return "", false
}

func createMissingBacklogSeeds(ctx context.Context, root string, client *http.Client, token string, teamID string, seeds []backlogSeed, pendingSeeds []backlogSeed) OperationResult {
	result := baseOperationResult(root)
	for index, seed := range pendingSeeds {
		issue, status, remaining, err := createBacklogIssue(ctx, client, token, teamID, seed)
		result.HTTPStatus = status
		result.RateLimitRemaining = remaining
		if err != nil {
			message := "Linear issue creation failed; remaining seed queued offline with synced=false. " + err.Error()
			queueBacklogOffline(root, &result, message, pendingSeeds[index:])
			return result
		}
		if issue != nil {
			result.Issues = append(result.Issues, *issue)
		}
	}
	result.Mode = "online"
	result.Synced = true
	result.Message = fmt.Sprintf("Created %d Linear backlog seed issues; skipped %d existing seeds.", len(result.Issues), len(seeds)-len(pendingSeeds))
	if !recordBacklogWriteEvidence(root, result.Issues) {
		result.Message += " Private write evidence was not fully recorded."
	}
	return result
}
