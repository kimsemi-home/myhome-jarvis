package codexcost

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/codexsustainability"
	"github.com/kimsemi-home/myhome-jarvis/internal/storagearchive"
)

func roiRows(
	policy Policy,
	cost Status,
	sustainability codexsustainability.Status,
	storage storagearchive.Status,
	valueProxy int64,
) []ROIRow {
	rows := make([]ROIRow, 0, len(policy.LoopScopes))
	for _, scope := range policy.LoopScopes {
		costUnits := cost.ByScope[scope]
		rowValue := allocateValueProxy(valueProxy, costUnits, cost.TotalUnits)
		rows = append(rows, ROIRow{
			Scope:                    scope,
			CostUnits:                costUnits,
			CostSharePercent:         costSharePercent(costUnits, cost.TotalUnits),
			Status:                   roiScopeStatus(costUnits),
			ValueProxyUnits:          rowValue,
			CostPerAcceptedChange:    roiCostPerChange(costUnits, sustainability),
			BudgetState:              cost.BudgetState,
			SustainabilityPosture:    sustainability.SustainabilityPosture,
			ReviewGateCount:          sustainability.ReviewGateCount,
			StorageArchivePattern:    storage.CompressionArchivePattern,
			NoiseBudgetReady:         storage.NoiseBudgetReady,
			EvidenceConfigIsEvidence: storage.ConfigIsEvidence,
			Recommendation: roiRecommendation(
				costUnits,
				rowValue,
				cost.BudgetState,
				sustainability,
				storage,
			),
		})
	}
	return rows
}

func countTrackedROIRows(rows []ROIRow) int {
	count := 0
	for _, row := range rows {
		if row.CostUnits > 0 {
			count++
		}
	}
	return count
}
