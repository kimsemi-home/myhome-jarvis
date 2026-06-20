package authority

func applyReviewProfilePlans(policy Policy, plan *ReviewPlanStatus) {
	for _, profile := range normalizedProfiles(policy.AssistantAuthorityProfiles) {
		item := ReviewProfilePlan{
			ProfileKey:                   profile.Key,
			AuthorityProfile:             profile.AuthorityProfile,
			ReviewClass:                  reviewClassForProfile(profile),
			RequiresHumanReview:          profile.RequiresHumanReview,
			PublicSafetyGateRequired:     profile.PublicSafetyGateRequired,
			PublicRepoChangeGateRequired: profile.PublicRepoChangeGateRequired,
			WorkflowChangeGateRequired:   profile.WorkflowChangeGateRequired,
			VerifierSeparationRequired:   profile.VerifierSeparationRequired,
			SelfApprovalAllowed:          profile.SelfApprovalAllowed,
		}
		plan.Profiles = append(plan.Profiles, item)
		applyProfileReviewCounts(profile, plan)
	}
}

func applyProfileReviewCounts(profile AssistantProfile, plan *ReviewPlanStatus) {
	if profile.RequiresHumanReview {
		plan.RequiredReviewClasses = append(plan.RequiredReviewClasses, "human_review")
	}
	if profile.PublicSafetyGateRequired {
		plan.PublicSafetyReviewProfileCount++
		plan.RequiredReviewClasses = append(plan.RequiredReviewClasses, "public_safety_review")
	}
	if profile.PublicRepoChangeGateRequired {
		plan.PublicRepoReviewProfileCount++
		plan.RequiredReviewClasses = append(plan.RequiredReviewClasses, "public_repo_change_review")
	}
	if profile.WorkflowChangeGateRequired {
		plan.WorkflowReviewProfileCount++
		plan.RequiredReviewClasses = append(plan.RequiredReviewClasses, "workflow_change_review")
	}
	if profile.VerifierSeparationRequired {
		plan.VerifierSeparationRequiredCount++
	}
	if profile.ExternalWritesAllowed {
		plan.ExternalWritesAllowedProfileCount++
	}
}
