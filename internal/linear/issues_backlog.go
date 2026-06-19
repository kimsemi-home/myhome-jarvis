package linear

func missingBacklogSeeds(seeds []backlogSeed, existingTitles map[string]struct{}) []backlogSeed {
	missing := make([]backlogSeed, 0, len(seeds))
	for _, seed := range seeds {
		if _, exists := existingTitles[normalizedIssueTitle(seed.Title)]; exists {
			continue
		}
		missing = append(missing, seed)
	}
	return missing
}

func backlogSeeds() []backlogSeed {
	return []backlogSeed{
		{
			Title:       "[myhome-jarvis] Track approved Linear write evidence",
			Description: "Acceptance: approved Linear mutations append private redacted evidence with issue key, action, and sync status only; default surfaces avoid raw descriptions, workspace URLs, identities, UUIDs, tokens, absolute paths, and local checkout paths.",
			Priority:    3,
		},
		{
			Title:       "[myhome-jarvis] Reconcile planner external-write gate",
			Description: "Acceptance: planner status distinguishes the standing external-write boundary from user-approved Linear work evidence without marking sync success unless the Linear API mutation succeeds.",
			Priority:    3,
		},
		{
			Title:       "[myhome-jarvis] Include project queue status in loop checkpoints",
			Description: "Acceptance: closed-loop checkpoint evidence includes redacted project queue availability from Linear summaries while keeping raw team names, workspace URLs, descriptions, UUIDs, tokens, and absolute paths out of public surfaces.",
			Priority:    3,
		},
	}
}
