package linear

import "net/http"

func prepareReplay(root string, client *http.Client) ([]queuedOfflineAction, replayRun, bool) {
	run := replayRun{
		result: ReplayResult{
			Mode:       "offline",
			Synced:     false,
			QueuePath:  privateRelativePath(filepathJoinSlash(root, OfflineQueueRelativePath)),
			ReplayPath: privateRelativePath(filepathJoinSlash(root, OfflineReplayRelativePath)),
		},
		root:   root,
		client: client,
		scope:  configuredIssueScope(),
	}
	entries, err := readQueuedOfflineActions(root)
	if err != nil {
		run.result.Message = "Offline queue could not be read; replay skipped."
		return nil, run, false
	}
	run.result.QueuedCount = len(entries)
	if len(entries) == 0 {
		run.result.Synced = true
		run.result.Message = "No offline Linear actions are queued."
		return entries, run, false
	}
	replayed, err := readReplayedEntryIDs(root)
	if err != nil {
		run.result.Message = "Offline replay evidence could not be read; replay skipped."
		return entries, run, false
	}
	run.replayed = replayed
	measureReplayEligibility(entries, &run)
	if run.result.EligibleCount == 0 {
		run.result.Synced = true
		run.result.Message = "No in-scope write-safe offline Linear actions are queued for replay."
		return entries, run, false
	}
	token, err := loadToken(root)
	if err != nil {
		run.result.Message = "No Linear token found. Offline replay skipped and queued entries remain synced=false."
		return entries, run, false
	}
	run.token = token.Value
	return entries, run, true
}

func measureReplayEligibility(entries []queuedOfflineAction, run *replayRun) {
	for _, entry := range entries {
		if entry.Synced {
			continue
		}
		if replaySafeKind(entry.Kind) && replayIssueMatchesScope(entry, run.scope) {
			run.result.EligibleCount++
			continue
		}
		run.result.SkippedCount++
	}
}
