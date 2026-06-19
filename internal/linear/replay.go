package linear

import (
	"context"
	"fmt"
	"net/http"
)

func ReplayOffline(ctx context.Context, root string, client *http.Client) ReplayResult {
	entries, run, ok := prepareReplay(root, client)
	if !ok {
		return run.result
	}
	run.result.Mode = "online"
	for _, entry := range entries {
		if run.skipEntry(entry) {
			continue
		}
		if run.replayed[entry.EntryID] {
			run.result.AlreadyReplayedCount++
			continue
		}
		if !run.applyEntry(ctx, entry) {
			break
		}
	}
	if run.result.Message == "" {
		run.result.Message = fmt.Sprintf(
			"Replayed %d write-safe offline Linear actions.",
			run.result.ReplayedCount,
		)
	}
	run.result.Synced = replayFullySynced(run.result)
	return run.result
}

func (run *replayRun) skipEntry(entry queuedOfflineAction) bool {
	return entry.Synced ||
		!replaySafeKind(entry.Kind) ||
		!replayIssueMatchesScope(entry, run.scope)
}

func (run *replayRun) applyEntry(ctx context.Context, entry queuedOfflineAction) bool {
	operation := replayQueuedAction(ctx, run.root, run.client, run.token, entry)
	run.result.HTTPStatus = operation.HTTPStatus
	run.result.RateLimitRemaining = operation.RateLimitRemaining
	if operation.RateLimited {
		run.result.RateLimited = true
		run.result.Message = operation.Message
		return false
	}
	if !operation.Synced {
		run.result.FailedCount++
		run.result.Message = operation.Message
		return false
	}
	run.result.ReplayedCount++
	if err := appendReplayRecord(run.root, entry); err != nil {
		run.result.FailedCount++
		run.result.Message = "Offline replay succeeded but private replay evidence was not recorded."
		return false
	}
	if rateLimitLow(operation.RateLimitRemaining) {
		run.result.RateLimited = true
		run.result.Message = "Linear rate limit remaining is low; offline replay paused."
		return false
	}
	return true
}
