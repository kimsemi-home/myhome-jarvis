package codexcost

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/codexsustainability"
	"github.com/kimsemi-home/myhome-jarvis/internal/storagearchive"
)

func roiRecommendation(
	costUnits int64,
	valueProxy int64,
	budgetState string,
	sustainability codexsustainability.Status,
	storage storagearchive.Status,
) string {
	if costUnits <= 0 {
		return "no_usage_yet"
	}
	if !storage.ArchiveReady ||
		!storage.NoiseBudgetReady ||
		!storage.ConfigIsEvidence {
		return "fix_evidence_archive"
	}
	if budgetState == "review_required" ||
		sustainability.ReviewGateCount > 0 ||
		sustainability.SustainabilityPosture == "blocked" ||
		sustainability.SustainabilityPosture == "review_required" {
		return "review_before_scaling"
	}
	if budgetState == "warning" {
		return "monitor_roi_before_scaling"
	}
	if valueProxy > costUnits {
		return "cache_value_supports_scaling"
	}
	return "monitor"
}
