package contextpack

func checkRequired(declaration Declaration, result *VerifyResult) {
	required := map[string]string{
		"pack_id":                        declaration.PackID,
		"context_pack_version":           declaration.ContextPackVersion,
		"upstream_compatibility_version": declaration.UpstreamCompatibilityVersion,
		"ontology_version":               declaration.OntologyVersion,
		"authority_contract_version":     declaration.AuthorityContractVersion,
		"security_contract_version":      declaration.SecurityContractVersion,
		"verification_profile":           declaration.VerificationProfile,
	}
	for _, field := range requiredDeclarationFields {
		if field == "ssot_artifact_versions" {
			if len(declaration.SSOTArtifactVersions) == 0 {
				missingField(field, result)
			}
			continue
		}
		if required[field] == "" {
			missingField(field, result)
		}
	}
}

func missingField(field string, result *VerifyResult) {
	result.MissingFieldCount++
	result.Findings = append(result.Findings, field+"_missing")
}
