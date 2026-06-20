package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/evidence"

func summarizeEvidenceIntegrity(status evidence.IntegrityStatus) EvidenceIntegritySummary {
	return EvidenceIntegritySummary{
		PublicSafe:               status.PublicSafe,
		CheckedEvidenceRefCount:  status.CheckedEvidenceRefCount,
		PresentEvidenceRefCount:  status.PresentEvidenceRefCount,
		DanglingEvidenceRefCount: status.DanglingEvidenceRefCount,
		NextSafeAction:           status.NextSafeAction,
	}
}
