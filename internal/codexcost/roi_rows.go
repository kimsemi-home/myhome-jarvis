package codexcost

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/codexsustainability"
	"github.com/kimsemi-home/myhome-jarvis/internal/storagearchive"
)

func roiRows(
	policy Policy,
	cost Status,
	attribution AttributionStatus,
	sustainability codexsustainability.Status,
	storage storagearchive.Status,
	valueProxy int64,
	costPerChange int64,
) []ROIRow {
	rows := make([]ROIRow, 0, len(policy.LoopScopes))
	for _, scope := range policy.LoopScopes {
		attributedUnits := attribution.ByScope[scope]
		costUnits := rowCostUnits(cost.ByScope[scope], attributedUnits)
		rowValue := allocateValueProxy(valueProxy, costUnits, cost.TotalUnits)
		rows = append(rows, ROIRow{
			Scope:                    scope,
			CostUnits:                costUnits,
			AttributedCostUnits:      attributedUnits,
			AttributionSubjectCount:  attribution.SubjectCountByScope[scope],
			CostSharePercent:         costSharePercent(costUnits, cost.TotalUnits),
			Status:                   roiScopeStatus(cost.ByScope[scope], attributedUnits),
			ValueProxyUnits:          rowValue,
			CostPerAcceptedChange:    roiCostPerChange(costUnits, costPerChange),
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

func rowCostUnits(directUnits int64, attributedUnits int64) int64 {
	if attributedUnits > 0 {
		return attributedUnits
	}
	return directUnits
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
