package authority

import "sort"

func applyProfiles(policy Policy, status *Status) {
	for _, profile := range normalizedProfiles(policy.AssistantAuthorityProfiles) {
		status.ProfileCount++
		status.ProfileKeys = append(status.ProfileKeys, profile.Key)
		if profile.RequiresHumanReview {
			status.ReviewRequiredProfileCount++
			status.ReviewRequiredProfiles = append(status.ReviewRequiredProfiles, profile.Key)
		}
		if profile.PublicSafetyGateRequired {
			status.PublicSafetyGatedProfileCount++
			status.PublicSafetyGatedProfiles = append(status.PublicSafetyGatedProfiles, profile.Key)
		}
		if profile.PublicRepoChangeGateRequired {
			status.PublicRepoChangeGatedProfileCount++
		}
		if profile.WorkflowChangeGateRequired {
			status.WorkflowChangeGatedProfileCount++
		}
		if !profile.SelfApprovalAllowed {
			status.SelfApprovalBlockedProfileCount++
		}
	}
	sort.Strings(status.ProfileKeys)
	sort.Strings(status.ReviewRequiredProfiles)
	sort.Strings(status.PublicSafetyGatedProfiles)
}

func normalizedProfiles(profiles []AssistantProfile) []AssistantProfile {
	normalized := make([]AssistantProfile, 0, len(profiles))
	for _, profile := range profiles {
		profile.Key = normalizeToken(profile.Key)
		profile.AuthorityProfile = normalizeToken(profile.AuthorityProfile)
		profile.DataSensitivity = normalizeToken(profile.DataSensitivity)
		if profile.Key != "" {
			normalized = append(normalized, profile)
		}
	}
	sort.Slice(normalized, func(i, j int) bool {
		return normalized[i].Key < normalized[j].Key
	})
	return normalized
}
