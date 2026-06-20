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
	merge mergeAcceptance,
) ROISummary {
	accepted := acceptedChangeEvidenceFor(sustainability, merge)
	valueProxy := accepted.AcceptedCount +
		sustainability.CacheSavingsUnits
	costPerChange := roiCostPerAcceptedChange(
		cost.TotalUnits,
		accepted.AcceptedCount,
		sustainability.CostPerAcceptedChange,
	)
	rows := roiRows(policy, cost, attribution, sustainability, storage,
		valueProxy, costPerChange)
	return ROISummary{
		PolicyPath:                      PolicyRelativePath,
		LedgerPath:                      policy.PrivateUsageLedger,
		AttributionLedgerPath:           policy.PrivateAttributionLedger,
		ScopeCount:                      len(policy.LoopScopes),
		TrackedScopeCount:               countTrackedROIRows(rows),
		TotalUnits:                      cost.TotalUnits,
		AttributedUnits:                 attribution.CoverageUnits,
		AttributionEntryUnits:           attribution.EntryUnits,
		AttributionRecordCount:          attribution.RecordCount,
		AttributionCostRefCount:         attribution.CostRefCount,
		AttributionCoveragePercent:      attributionCoverage(cost, attribution),
		InvalidAttributionCount:         attribution.InvalidRecordCount,
		BudgetState:                     cost.BudgetState,
		SustainabilityPosture:           sustainability.SustainabilityPosture,
		TrendPosture:                    sustainability.TrendPosture,
		ReviewGateCount:                 sustainability.ReviewGateCount,
		AcceptedChangeCount:             accepted.AcceptedCount,
		LedgerAcceptedChangeCount:       accepted.LedgerCount,
		MergeAcceptedChangeCount:        accepted.MergeCount,
		AcceptedChangeEvidenceSource:    accepted.Source,
		AcceptedChangeLogLimit:          accepted.LogLimit,
		CacheSavingsUnits:               sustainability.CacheSavingsUnits,
		ValueProxyUnits:                 valueProxy,
		CostPerAcceptedChange:           costPerChange,
		ValueProxyMethod:                valueProxyMethod,
		StorageArchivePattern:           storage.CompressionArchivePattern,
		StorageArchiveReady:             storage.ArchiveReady,
		NoiseBudgetReady:                storage.NoiseBudgetReady,
		MaxNoiseRatioPercent:            storage.MaxNoiseRatioPercent,
		ArchiveManifestEntryCount:       storage.ManifestEntryCount,
		ArchiveManifestBudgetBreaches:   storage.ManifestBudgetBreachCount,
		ArchiveManifestInvalidEntries:   storage.ManifestInvalidEntryCount,
		ArchiveManifestCompressionRatio: storage.ManifestCompressionRatio,
		ConfigEvidenceField:             storage.ConfigEvidenceField,
		ConfigIsEvidence:                storage.ConfigIsEvidence,
		PrivateLogSourceKeys:            storage.PrivateLogSourceKeys,
		ForbiddenPublicFieldCount:       len(policy.ForbiddenPublicFields),
		Rows:                            rows,
		CheckedAt:                       time.Now().UTC().Format(time.RFC3339),
	}
}

func attributionCoverage(cost Status, attribution AttributionStatus) int {
	return costSharePercent(attribution.CoverageUnits, cost.TotalUnits)
}
