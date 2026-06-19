package linear

func transitionMutationResult(
	root string,
	issueID string,
	httpStatus int,
	remaining int,
	err error,
	response struct {
		IssueUpdate struct {
			Success bool   `json:"success"`
			Issue   *Issue `json:"issue"`
		} `json:"issueUpdate"`
	},
) ReplayResult {
	if err != nil || !response.IssueUpdate.Success {
		return ReplayResult{
			Mode:               "online",
			Synced:             false,
			HTTPStatus:         httpStatus,
			RateLimitRemaining: remaining,
			Message:            "Offline Linear transition replay failed; entry remains synced=false.",
		}
	}
	evidenceIssueKey := issueID
	if response.IssueUpdate.Issue != nil {
		evidenceIssueKey = response.IssueUpdate.Issue.Identifier
	}
	_ = AppendWriteEvidence(root, offlineReplayTransitionKind, evidenceIssueKey)
	return ReplayResult{
		Mode:               "online",
		Synced:             true,
		HTTPStatus:         httpStatus,
		RateLimitRemaining: remaining,
		Message:            "Offline Linear transition replayed.",
	}
}
