package contextpack

func publicSafe(policy Policy) bool {
	return policy.PublicStatusRedacted &&
		!policy.RawPrivateContextPublicAllowed &&
		!policy.AuthorityContract.SelfApprovalAllowed &&
		policy.AuthorityContract.PublicSafetyGateRequired &&
		!policy.SecurityContract.PrivatePathsPublicAllowed &&
		!policy.SecurityContract.LocalPathsPublicAllowed &&
		forbiddenValueCount(policy) == 0
}

func forbiddenValueCount(policy Policy) int {
	values := []string{policy.PackID, policy.DeclarationPath, policy.MissionSource,
		policy.BoundedContextSource}
	for _, artifact := range policy.ExportedArtifacts {
		values = append(values, artifact.Path, artifact.Version)
	}
	count := 0
	for _, value := range values {
		if containsUnsafeText(value) {
			count++
		}
	}
	return count
}

func unsafeMarkers() []string {
	oldOwner := "kim" + "jooyoon"
	oldTeam := "kim-joo" + "-yoon"
	return []string{"/" + "users" + "/", oldOwner, oldTeam, "raw_prompt",
		"raw_transcript", "credential", "cookie", "private_evidence",
		"browser_session", "unpublished_monetization"}
}
