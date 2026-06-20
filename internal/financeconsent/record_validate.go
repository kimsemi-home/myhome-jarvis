package financeconsent

type recordResult struct {
	Valid            bool
	MissingEvidence  bool
	ReviewRequired   bool
	RevokedOrExpired bool
	Active           bool
}

func validateRecord(record Record, policy Policy) recordResult {
	result := recordResult{Valid: true}
	if record.At == "" || record.ConsentKind == "" || record.SubjectScope == "" {
		result.Valid = false
	}
	if !contains(policy.ConsentKinds, record.ConsentKind) ||
		!contains(policy.ConsentStatuses, record.Status) ||
		!contains(policy.ReviewStatuses, record.ReviewStatus) ||
		!contains(policy.AuthorityProfiles, record.AuthorityProfile) {
		result.Valid = false
	}
	if len(record.EvidenceRefs) == 0 || !evidenceRefsAllowed(record.EvidenceRefs, policy) {
		result.MissingEvidence = true
	}
	if record.ReviewStatus == "requested" {
		result.ReviewRequired = true
	}
	if record.Status == "revoked" || record.Status == "expired" || isExpired(record.ExpiresAt) {
		result.RevokedOrExpired = true
	}
	result.Active = result.Valid && !result.MissingEvidence &&
		record.Status == "granted" && record.ReviewStatus == "approved" &&
		!result.RevokedOrExpired
	return result
}
