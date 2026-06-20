package contextpack

func verifyIdentity(policy Policy, declaration Declaration, result *VerifyResult) {
	checkRequired(declaration, result)
	checkEqual("pack_id", declaration.PackID, policy.PackID, result)
	checkEqual("context_pack_version", declaration.ContextPackVersion, policy.Version, result)
	checkEqual("upstream_compatibility_version",
		declaration.UpstreamCompatibilityVersion,
		policy.UpstreamCompatibilityVersion, result)
	checkEqual("ontology_version", declaration.OntologyVersion,
		policy.OntologyVersion, result)
	checkEqual("authority_contract_version",
		declaration.AuthorityContractVersion,
		policy.AuthorityContract.Version, result)
	checkEqual("security_contract_version",
		declaration.SecurityContractVersion,
		policy.SecurityContract.Version, result)
	checkEqual("verification_profile", declaration.VerificationProfile,
		policy.VerificationProfile.Name, result)
}

func checkEqual(field string, actual string, expected string, result *VerifyResult) {
	if actual == expected {
		return
	}
	result.DriftCount++
	result.StaleVersionCount++
	result.Findings = append(result.Findings, field+"_mismatch")
}
