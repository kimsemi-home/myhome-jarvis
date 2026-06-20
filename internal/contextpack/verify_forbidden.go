package contextpack

func verifyForbiddenValues(declaration Declaration, result *VerifyResult) {
	values := []string{declaration.PackID, declaration.ContextPackVersion,
		declaration.UpstreamCompatibilityVersion, declaration.OntologyVersion,
		declaration.AuthorityContractVersion,
		declaration.SecurityContractVersion, declaration.VerificationProfile}
	for _, artifact := range declaration.SSOTArtifactVersions {
		values = append(values, artifact.Role, artifact.Path, artifact.Version)
	}
	for _, value := range values {
		if containsUnsafeText(value) {
			result.ForbiddenValueCount++
			result.Findings = append(result.Findings, "forbidden_public_value")
		}
	}
}
