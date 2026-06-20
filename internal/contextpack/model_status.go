package contextpack

import "time"

type Status struct {
	PolicyPath                    string `json:"policy_path"`
	PackID                        string `json:"pack_id"`
	Version                       string `json:"version"`
	UpstreamCompatibilityVersion  string `json:"upstream_compatibility_version"`
	OntologyVersion               string `json:"ontology_version"`
	DeclarationPath               string `json:"declaration_path"`
	PublicSafe                    bool   `json:"public_safe"`
	SplitCriteriaCount            int    `json:"split_criteria_count"`
	ExportedArtifactCount         int    `json:"exported_artifact_count"`
	AuthorityContractVersion      string `json:"authority_contract_version"`
	SecurityContractVersion       string `json:"security_contract_version"`
	VerificationProfile           string `json:"verification_profile"`
	VerificationRequiredUnitCount int    `json:"verification_required_unit_count"`
	ForbiddenPublicFieldCount     int    `json:"forbidden_public_field_count"`
	CheckedAt                     string `json:"checked_at"`
}

func statusFromPolicy(policy Policy, at time.Time) Status {
	return Status{PolicyPath: PolicyRelativePath, PackID: policy.PackID,
		Version: policy.Version, UpstreamCompatibilityVersion: policy.UpstreamCompatibilityVersion,
		OntologyVersion: policy.OntologyVersion, DeclarationPath: policy.DeclarationPath,
		PublicSafe: publicSafe(policy), SplitCriteriaCount: len(policy.SplitCriteria),
		ExportedArtifactCount:         len(policy.ExportedArtifacts),
		AuthorityContractVersion:      policy.AuthorityContract.Version,
		SecurityContractVersion:       policy.SecurityContract.Version,
		VerificationProfile:           policy.VerificationProfile.Name,
		VerificationRequiredUnitCount: len(policy.VerificationProfile.RequiredUnits),
		ForbiddenPublicFieldCount:     len(policy.ForbiddenPublicFields),
		CheckedAt:                     at.UTC().Format(time.RFC3339)}
}
