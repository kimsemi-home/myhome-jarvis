package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/financeconsent"

func summarizeFinanceConsent(status financeconsent.Status) FinanceConsentSummary {
	return FinanceConsentSummary{
		ReadinessState:              status.ReadinessState,
		FinanceMode:                 status.FinanceMode,
		RecordCount:                 status.RecordCount,
		ActiveConsentCount:          status.ActiveConsentCount,
		MissingRequiredConsentCount: status.MissingRequiredConsentCount,
		ReviewRequiredCount:         status.ReviewRequiredCount,
		MissingEvidenceCount:        status.MissingEvidenceCount,
		InvalidRecordCount:          status.InvalidRecordCount,
		RevokedOrExpiredCount:       status.RevokedOrExpiredCount,
		ForbiddenActionEnabledCount: status.ForbiddenActionEnabledCount,
		ConsentDebtCount:            status.ConsentDebtCount,
	}
}
