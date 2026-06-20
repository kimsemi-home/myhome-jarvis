package contextpack

func splitKeys(policy Policy) []string {
	keys := make([]string, 0, len(policy.SplitCriteria))
	for _, item := range policy.SplitCriteria {
		keys = append(keys, item.Key)
	}
	return keys
}

func artifactRoles(artifacts []Artifact) []string {
	roles := make([]string, 0, len(artifacts))
	for _, item := range artifacts {
		roles = append(roles, item.Role)
	}
	return roles
}

func artifactMap(artifacts []Artifact) map[string]string {
	versions := map[string]string{}
	for _, item := range artifacts {
		versions[item.Path] = item.Version
	}
	return versions
}
