package financeconsent

func resultForRecord(record Record, status Status) RecordResult {
	return RecordResult{
		ConsentKind:                 record.ConsentKind,
		SubjectScope:                record.SubjectScope,
		Status:                      record.Status,
		ReviewStatus:                record.ReviewStatus,
		AuthorityProfile:            record.AuthorityProfile,
		EvidenceRefCount:            len(record.EvidenceRefs),
		ReadinessState:              status.ReadinessState,
		ActiveConsentCount:          status.ActiveConsentCount,
		MissingRequiredConsentCount: status.MissingRequiredConsentCount,
		ConsentDebtCount:            status.ConsentDebtCount,
		RecordedAt:                  record.At,
	}
}
