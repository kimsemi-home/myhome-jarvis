package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/codexcost"

func summarizeCodexCostBrief(brief codexcost.Brief) CodexCostBriefSummary {
	return CodexCostBriefSummary{
		PublicSafe:                      brief.PublicSafe,
		Decision:                        brief.Decision,
		Recommendation:                  brief.Recommendation,
		NextSafeAction:                  brief.NextSafeAction,
		BudgetState:                     brief.BudgetState,
		TotalUnits:                      brief.TotalUnits,
		AttributionCoveragePercent:      brief.AttributionCoveragePercent,
		AcceptedChangeCount:             brief.AcceptedChangeCount,
		CacheSavingsUnits:               brief.CacheSavingsUnits,
		ValueProxyUnits:                 brief.ValueProxyUnits,
		CostPerAcceptedChange:           brief.CostPerAcceptedChange,
		StorageArchivePattern:           brief.StorageArchivePattern,
		StorageArchiveReady:             brief.StorageArchiveReady,
		NoiseBudgetReady:                brief.NoiseBudgetReady,
		MaxNoiseRatioPercent:            brief.MaxNoiseRatioPercent,
		ArchiveManifestEntryCount:       brief.ArchiveManifestEntryCount,
		ArchiveManifestBudgetBreaches:   brief.ArchiveManifestBudgetBreaches,
		ArchiveManifestInvalidEntries:   brief.ArchiveManifestInvalidEntries,
		ArchiveManifestCompressionRatio: brief.ArchiveManifestCompressionRatio,
		ConfigIsEvidence:                brief.ConfigIsEvidence,
		ForbiddenPublicFieldCount:       brief.ForbiddenPublicFieldCount,
	}
}
