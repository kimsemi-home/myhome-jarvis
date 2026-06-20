package codexcost

func scalingStorage(brief Brief) ScalingStorage {
	return ScalingStorage{
		Pattern:                  brief.StorageArchivePattern,
		Ready:                    brief.StorageArchiveReady,
		NoiseBudgetReady:         brief.NoiseBudgetReady,
		MaxNoiseRatioPercent:     brief.MaxNoiseRatioPercent,
		ManifestEntryCount:       brief.ArchiveManifestEntryCount,
		ManifestBudgetBreaches:   brief.ArchiveManifestBudgetBreaches,
		ManifestInvalidEntries:   brief.ArchiveManifestInvalidEntries,
		ManifestCompressionRatio: brief.ArchiveManifestCompressionRatio,
		ConfigIsEvidence:         brief.ConfigIsEvidence,
	}
}

func remainingUnits(threshold int64, used int64) int64 {
	if threshold <= used {
		return 0
	}
	return threshold - used
}

func usedPercent(used int64, threshold int64) int {
	if threshold <= 0 || used <= 0 {
		return 0
	}
	return int((used * 100) / threshold)
}

func scalingOptions(decision string) []ScalingOption {
	reviewRecommended := decision == "review_required"
	warnRecommended := decision == "warn"
	return []ScalingOption{
		scalingOption(
			"continue_local_first_loop", "Continue local loop",
			decision == "allow", "no_paid_or_external_change", false, false,
		),
		scalingOption(
			"estimate_next_loop_with_guard", "Run cost guard",
			warnRecommended, "requires_guard_before_scaling", true, false,
		),
		scalingOption(
			"hold_paid_or_external_expansion", "Hold expansion",
			reviewRecommended, "keeps_paid_and_external_expansion_blocked", true, true,
		),
	}
}

func scalingOption(
	key string,
	label string,
	recommended bool,
	effect string,
	requiresGuard bool,
	requiresReview bool,
) ScalingOption {
	return ScalingOption{
		Key:                           key,
		Label:                         label,
		Recommended:                   recommended,
		Effect:                        effect,
		RequiresGuard:                 requiresGuard,
		RequiresHumanReview:           requiresReview,
		RequiresSeparateRecordCommand: requiresGuard || requiresReview,
		ThisPacketGrantsSpend:         false,
		AllowsPaidExpansion:           false,
		AllowsExternalTooling:         false,
		AllowsWorkflowChanges:         false,
		AllowsSelfApproval:            false,
	}
}
