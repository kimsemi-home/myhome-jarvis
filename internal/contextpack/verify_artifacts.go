package contextpack

func verifyArtifacts(policy Policy, declaration Declaration, result *VerifyResult) {
	declared := artifactMap(declaration.SSOTArtifactVersions)
	for _, artifact := range policy.ExportedArtifacts {
		version, ok := declared[artifact.Path]
		if !ok {
			result.MissingArtifactCount++
			result.Findings = append(result.Findings,
				"artifact_missing:"+artifact.Path)
			continue
		}
		if version != artifact.Version {
			result.DriftCount++
			result.StaleVersionCount++
			result.Findings = append(result.Findings,
				"artifact_version_mismatch:"+artifact.Path)
		}
	}
}
