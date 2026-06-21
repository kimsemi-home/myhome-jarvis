package externalevidence

import "time"

func StatusForRoot(root string) (Status, error) {
	policy, err := readPolicy(root)
	if err != nil {
		return Status{}, err
	}
	if err := ValidatePolicy(policy); err != nil {
		return Status{}, err
	}
	manifest, err := readManifestSummary(root, policy.ManifestPath)
	if err != nil {
		return Status{}, err
	}
	return statusFromPolicy(policy, manifest), nil
}

func statusFromPolicy(policy Policy, manifest manifestSummary) Status {
	assessment := policy.RepoSplitAssessment
	return Status{
		PolicyPath:                     PolicyRelativePath,
		SchemaVersion:                  policy.SchemaVersion,
		PublicSafe:                     policy.PublicSafe,
		ExternalCollectionAllowed:      policy.ExternalNetworkCollectionAllowed,
		CredentialsAllowed:             policy.CredentialsAllowed,
		CookiesAllowed:                 policy.CookiesAllowed,
		RawPayloadPublicAllowed:        policy.RawPayloadPublicAllowed,
		PrivateRoot:                    policy.PrivateRoot,
		ManifestPath:                   policy.ManifestPath,
		RawLayerPath:                   policy.RawLayerPath,
		BronzeLayerPath:                policy.BronzeLayerPath,
		SilverLayerPath:                policy.SilverLayerPath,
		GoldLayerPath:                  policy.GoldLayerPath,
		ArchiveSourceKey:               policy.ArchiveSourceKey,
		StorageArchiveSourcePath:       policy.StorageArchiveSourcePath,
		SourceCount:                    len(policy.SourceDescriptors),
		SourceClasses:                  policy.SourceClasses,
		PreprocessingRules:             policy.PreprocessingRules,
		ManifestPresent:                manifest.Present,
		ManifestRecordCount:            manifest.Count,
		LatestCollectedAt:              manifest.LatestAt,
		RepoSplitRecommendation:        assessment.Recommendation,
		FutureRepoCandidate:            assessment.FutureRepoCandidate,
		RepoCreationGate:               assessment.CreationGate,
		SplitTriggerCount:              len(assessment.SplitTriggers),
		CurrentRepoResponsibilityCount: len(assessment.CurrentRepoResponsibilities),
		FutureRepoResponsibilityCount:  len(assessment.FutureRepoResponsibilities),
		PublicRepoRules:                assessment.PublicRepoRules,
		Commands:                       policy.Commands,
		CheckedAt:                      time.Now().UTC().Format(time.RFC3339),
	}
}
