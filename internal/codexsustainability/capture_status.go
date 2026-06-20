package codexsustainability

import (
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/qualitylog"
)

func newCaptureStatus(policy Policy, quality qualitylog.Status, now time.Time) CaptureStatus {
	return CaptureStatus{
		PolicyPath:                 PolicyRelativePath,
		QualityLedgerPath:          qualitylog.RelativePath,
		SustainabilityLedgerPath:   policy.PrivateEvidenceLedger,
		CaptureState:               "unknown",
		QualityRunCount:            quality.Count,
		EvidenceRef:                qualitylog.RelativePath,
		PublicSafe:                 true,
		Redaction:                  "codex-sustainability-quality-capture-public-status",
		ApprovalState:              "not_approved",
		ApprovalGranted:            false,
		ExternalWritesAllowed:      false,
		SelfApprovalAllowed:        false,
		RawEvidencePublicAllowed:   false,
		PrivateLedgerWriteRequired: true,
		CheckedAt:                  now.Format(time.RFC3339),
	}
}
