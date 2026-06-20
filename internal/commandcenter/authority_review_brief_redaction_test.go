package commandcenter

func authorityReviewBriefForbiddenFields() []string {
	return []string{
		"raw_rationale", "raw_evidence", "raw_prompt", "raw_transcript",
		"private_notes", "reviewer_identity", "linear_url",
		"local_absolute_path", "finance_payload", "private_ledger",
		"token", "secret", "credential", "cookie", "account_id",
		"card_number", "/" + "Users" + "/",
	}
}
