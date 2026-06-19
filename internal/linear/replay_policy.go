package linear

import (
	"encoding/json"
	"strings"
)

func replayIssueKey(entry queuedOfflineAction) string {
	switch entry.Kind {
	case offlineReplayCommentKind:
		var payload commentPayload
		if json.Unmarshal(entry.Payload, &payload) == nil {
			return publicIssueKey(payload.IssueID)
		}
	case offlineReplayTransitionKind:
		var payload transitionPayload
		if json.Unmarshal(entry.Payload, &payload) == nil {
			return publicIssueKey(payload.IssueID)
		}
	}
	return ""
}

func replaySafeKind(kind string) bool {
	switch strings.TrimSpace(kind) {
	case offlineReplayCommentKind, offlineReplayTransitionKind:
		return true
	default:
		return false
	}
}

func replayIssueMatchesScope(entry queuedOfflineAction, scope issueScope) bool {
	teamKey := strings.TrimSpace(scope.TeamKey)
	if teamKey == "" {
		return true
	}
	issueKey := replayIssueKey(entry)
	if issueKey == "" {
		return false
	}
	return strings.HasPrefix(issueKey, strings.ToUpper(teamKey)+"-")
}

func rateLimitLow(remaining int) bool {
	return remaining > 0 && remaining <= defaultReplayRateLimitFloor
}

func replayFullySynced(result ReplayResult) bool {
	return result.FailedCount == 0 &&
		!result.RateLimited &&
		result.ReplayedCount+result.AlreadyReplayedCount == result.EligibleCount
}
