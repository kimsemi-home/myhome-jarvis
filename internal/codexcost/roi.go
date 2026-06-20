package codexcost

import (
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/codexsustainability"
	"github.com/kimsemi-home/myhome-jarvis/internal/storagearchive"
)

const valueProxyMethod = "accepted_changes_plus_cache_savings_by_cost_share"

func ROISummaryForRoot(root string) (ROISummary, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return ROISummary{}, err
	}
	cost, err := StatusForRoot(root)
	if err != nil {
		return ROISummary{}, err
	}
	sustainability, err := codexsustainability.StatusForRoot(root)
	if err != nil {
		return ROISummary{}, err
	}
	storage, err := storagearchive.StatusForRoot(root)
	if err != nil {
		return ROISummary{}, err
	}
	return buildROISummary(policy, cost, sustainability, storage), nil
}

func buildROISummary(
	policy Policy,
	cost Status,
	sustainability codexsustainability.Status,
	storage storagearchive.Status,
) ROISummary {
	valueProxy := sustainability.AcceptedChangeCount +
		sustainability.CacheSavingsUnits
	rows := roiRows(policy, cost, sustainability, storage, valueProxy)
	return ROISummary{
		PolicyPath:                PolicyRelativePath,
		LedgerPath:                policy.PrivateUsageLedger,
		ScopeCount:                len(policy.LoopScopes),
		TrackedScopeCount:         countTrackedROIRows(rows),
		TotalUnits:                cost.TotalUnits,
		BudgetState:               cost.BudgetState,
		SustainabilityPosture:     sustainability.SustainabilityPosture,
		TrendPosture:              sustainability.TrendPosture,
		ReviewGateCount:           sustainability.ReviewGateCount,
		AcceptedChangeCount:       sustainability.AcceptedChangeCount,
		CacheSavingsUnits:         sustainability.CacheSavingsUnits,
		ValueProxyUnits:           valueProxy,
		CostPerAcceptedChange:     sustainability.CostPerAcceptedChange,
		ValueProxyMethod:          valueProxyMethod,
		StorageArchivePattern:     storage.CompressionArchivePattern,
		StorageArchiveReady:       storage.ArchiveReady,
		NoiseBudgetReady:          storage.NoiseBudgetReady,
		MaxNoiseRatioPercent:      storage.MaxNoiseRatioPercent,
		ConfigEvidenceField:       storage.ConfigEvidenceField,
		ConfigIsEvidence:          storage.ConfigIsEvidence,
		PrivateLogSourceKeys:      storage.PrivateLogSourceKeys,
		ForbiddenPublicFieldCount: len(policy.ForbiddenPublicFields),
		Rows:                      rows,
		CheckedAt:                 time.Now().UTC().Format(time.RFC3339),
	}
}
