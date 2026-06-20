package repofactory

var requiredTemplateRoles = []string{
	"generated_ci",
	"security_scan",
	"private_data_policy",
	"bootstrap_checklist",
	"codex_project",
	"context_pack_declaration",
}

var requiredCreationGates = []string{
	"authority_review",
	"public_safety_evidence",
	"generated_ci",
	"private_data_policy",
	"bootstrap_checklist",
}

var requiredChecklistItems = []string{
	"choose public-safe repository name",
	"generate quality workflow from SSOT",
	"run mhj security check",
	"run mhj security history",
	"record authority review evidence",
	"record public safety evidence",
	"declare consumed context pack version",
}

var requiredSummaryFields = []string{
	"policy_path",
	"template_file_count",
	"creation_gate_count",
	"bootstrap_check_count",
	"authority_review_required",
	"public_safety_evidence_required",
	"public_safe",
	"missing_template_role_count",
	"missing_creation_gate_count",
	"forbidden_template_value_count",
	"checked_at",
}
