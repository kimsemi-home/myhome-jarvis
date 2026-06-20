package commandcenter

func visionAuditFixtureStatus(policy visionPolicy) Status {
	return Status{
		PublicSafe:       true,
		NextSafeAction:   "await_human_authority_review",
		BlockedGateCount: 2,
		BlockedGates: []GateSummary{
			{Key: "authority"},
			{Key: "repo_factory"},
		},
		Vision: visionAuditFixtureSummary(policy),
		MediaReadiness: MediaReadinessSummary{
			PublicSafe:    true,
			PlaybackReady: true,
		},
		FinanceConsent: FinanceConsentSummary{
			ReadinessState: "ready_read_only",
		},
		RepoFactory: RepoFactorySummary{
			PublicSafe:                     true,
			RepoCreationBlockedUntilReview: true,
		},
		Monetization: MonetizationSummary{
			ExperimentCount: 1,
		},
		Cost: CostSummary{
			BudgetState: "ok",
		},
		PDCA: PDCASummary{
			Ready: true,
		},
		Authority: AuthoritySummary{
			Outcome: "blocked",
		},
		StorageArchive: readyStorageArchiveSummary(),
	}
}

func visionAuditFixtureSummary(policy visionPolicy) VisionSummary {
	return VisionSummary{
		CapabilityCount:    len(policy.CapabilityPillars),
		PillarKeys:         visionPillarKeys(policy.CapabilityPillars),
		ReadyPillarCount:   4,
		GatedPillarCount:   2,
		BlockedPillarCount: 0,
		ReadyPillarKeys: []string{
			"local_media_concierge", "household_finance_copilot",
			"monetization_console", "codex_cost_governor",
		},
		GatedPillarKeys: []string{
			"shorts_factory_control_plane", "self_improvement_loop",
		},
		BlockedPillarKeys: []string{},
	}
}
