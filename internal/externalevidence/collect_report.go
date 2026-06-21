package externalevidence

func newCollectReport(policy Policy, now string) CollectReport {
	return CollectReport{
		PolicyPath:              PolicyRelativePath,
		PublicSafe:              policy.PublicSafe,
		RawPayloadPublicAllowed: policy.RawPayloadPublicAllowed,
		SourceCount:             len(policy.SourceDescriptors),
		ManifestPath:            policy.ManifestPath,
		RawLayerPath:            policy.RawLayerPath,
		BronzeLayerPath:         policy.BronzeLayerPath,
		SilverLayerPath:         policy.SilverLayerPath,
		GoldLayerPath:           policy.GoldLayerPath,
		ArchiveSourceKey:        policy.ArchiveSourceKey,
		RepoSplitRecommendation: policy.RepoSplitAssessment.Recommendation,
		RepoCreationGate:        policy.RepoSplitAssessment.CreationGate,
		CollectionRunState:      "empty",
		CheckedAt:               now,
	}
}

func applyCollectResult(report *CollectReport, result CollectResult) {
	switch result.State {
	case "collected":
		report.CollectedCount++
	case "cached":
		report.CachedCount++
	case "failed":
		report.FailedCount++
	}
}

func collectRunState(report CollectReport) string {
	if report.FailedCount > 0 && report.CollectedCount+report.CachedCount > 0 {
		return "partial"
	}
	if report.FailedCount > 0 {
		return "failed"
	}
	if report.CollectedCount > 0 {
		return "collected"
	}
	if report.CachedCount > 0 {
		return "cached"
	}
	return "empty"
}
