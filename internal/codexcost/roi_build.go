package codexcost

import (
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/codexsustainability"
	"github.com/kimsemi-home/myhome-jarvis/internal/storagearchive"
)

func buildROISummary(
	policy Policy,
	cost Status,
	attribution AttributionStatus,
	sustainability codexsustainability.Status,
	storage storagearchive.Status,
) ROISummary {
	valueProxy := sustainability.AcceptedChangeCount +
		sustainability.CacheSavingsUnits
	rows := roiRows(policy, cost, attribution, sustainability, storage, valueProxy)
	return ROISummary{
		PolicyPath:                 PolicyRelativePath,
		LedgerPath:                 policy.PrivateUsageLedger,
		AttributionLedgerPath:      policy.PrivateAttributionLedger,
		ScopeCount:                 len(policy.LoopScopes),
		TrackedScopeCount:          countTrackedROIRows(rows),
		TotalUnits:                 cost.TotalUnits,
		AttributedUnits:            attribution.TotalUnits,
		AttributionRecordCount:     attribution.RecordCount,
		AttributionCoveragePercent: attributionCoverage(cost, attribution),
		InvalidAttributionCount:    attribution.InvalidRecordCount,
		BudgetState:                cost.BudgetState,
		SustainabilityPosture:      sustainability.SustainabilityPosture,
		TrendPosture:               sustainability.TrendPosture,
		ReviewGateCount:            sustainability.ReviewGateCount,
		AcceptedChangeCount:        sustainability.AcceptedChangeCount,
		CacheSavingsUnits:          sustainability.CacheSavingsUnits,
		ValueProxyUnits:            valueProxy,
		CostPerAcceptedChange:      sustainability.CostPerAcceptedChange,
		ValueProxyMethod:           valueProxyMethod,
		StorageArchivePattern:      storage.CompressionArchivePattern,
		StorageArchiveReady:        storage.ArchiveReady,
		NoiseBudgetReady:           storage.NoiseBudgetReady,
		MaxNoiseRatioPercent:       storage.MaxNoiseRatioPercent,
		ConfigEvidenceField:        storage.ConfigEvidenceField,
		ConfigIsEvidence:           storage.ConfigIsEvidence,
		PrivateLogSourceKeys:       storage.PrivateLogSourceKeys,
		ForbiddenPublicFieldCount:  len(policy.ForbiddenPublicFields),
		Rows:                       rows,
		CheckedAt:                  time.Now().UTC().Format(time.RFC3339),
	}
}

func attributionCoverage(cost Status, attribution AttributionStatus) int {
	return costSharePercent(attribution.TotalUnits, cost.TotalUnits)
}
