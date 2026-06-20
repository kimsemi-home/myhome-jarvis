package contextpack

var requiredSplitCriteria = []string{
	"responsibility_overload", "ownership_boundary",
	"independent_release_cadence", "private_data_boundary",
	"ci_cost_cache_impact",
}

var requiredExportRoles = []string{
	"mission", "ontology", "authority", "security", "verification",
	"repo_factory",
}

var requiredDeclarationFields = []string{
	"pack_id", "context_pack_version", "upstream_compatibility_version",
	"ontology_version", "authority_contract_version",
	"security_contract_version", "verification_profile",
	"ssot_artifact_versions",
}

var requiredCommands = []string{
	"mhj context-pack status",
	"mhj context-pack verify",
}
