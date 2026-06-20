package contextpack

func testSplitCriteria() []SplitCriterion {
	return []SplitCriterion{
		{Key: "responsibility_overload", Meaning: "too many responsibilities"},
		{Key: "ownership_boundary", Meaning: "ownership diverges"},
		{Key: "independent_release_cadence", Meaning: "separate releases"},
		{Key: "private_data_boundary", Meaning: "private data isolation"},
		{Key: "ci_cost_cache_impact", Meaning: "CI cache impact"},
	}
}

func testArtifacts() []Artifact {
	return []Artifact{
		{Role: "mission", Path: "generated/assistant_vision.generated.json", Version: "v1"},
		{Role: "ontology", Path: "generated/concepts.generated.json", Version: "concept-registry/v1"},
		{Role: "authority", Path: "generated/authority.generated.json", Version: "authority/v1"},
		{Role: "security", Path: "generated/security.generated.json", Version: "security/v1"},
		{Role: "verification", Path: "generated/verification_graph.generated.json", Version: "verification.graph/v1"},
		{Role: "repo_factory", Path: "generated/repo_factory.generated.json", Version: "v1"},
	}
}
