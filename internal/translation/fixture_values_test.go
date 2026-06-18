package translation

func translationTestContexts() []string {
	return []string{"HomeControl", "HouseholdFinance", "CommerceIntelligence", "ConnectorReadiness", "AgentCluster", "StorageLake", "SecurityPolicy", "AgentOps", "KnowledgeIndex"}
}

func translationTestManifestFields() []string {
	return []string{"source_context", "target_context", "source_version", "target_version", "preserved_rules", "known_losses", "owner", "evidence_refs"}
}

func translationTestLossLevels() []string {
	return []string{"l0_none", "l1_note", "l2_degraded", "l3_review_required", "l4_forbidden"}
}

func translationTestLossCategories() []string {
	return []string{"none", "mapping_gap", "version_drift", "field_drop", "precision_loss", "review_needed", "authority", "security_boundary", "user_consent", "deletion_semantics", "audit_record", "legal_obligation", "financial_commitment"}
}

func translationTestForbiddenLossCategories() []string {
	return []string{"authority", "security_boundary", "user_consent", "deletion_semantics", "audit_record", "legal_obligation", "financial_commitment"}
}
