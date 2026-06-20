package repofactory

import "time"

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	return statusFromPolicy(policy), nil
}

func statusFromPolicy(policy Policy) Status {
	roles := templateRoles(policy.TemplateFiles)
	keys := gateKeys(policy.CreationGates)
	missingRoles := containsAll(roles, requiredTemplateRoles)
	missingGates := containsAll(keys, requiredCreationGates)
	forbiddenCount := forbiddenTemplateValueCount(policy)
	checklistMissing := containsAll(policy.BootstrapChecklist, requiredChecklistItems)
	status := Status{
		PolicyPath:                     PolicyRelativePath,
		TemplateFileCount:              len(policy.TemplateFiles),
		CreationGateCount:              len(policy.CreationGates),
		BootstrapCheckCount:            len(policy.BootstrapChecklist),
		AuthorityReviewRequired:        policy.AuthorityReviewRequired,
		PublicSafetyEvidenceRequired:   policy.PublicSafetyEvidenceRequired,
		CodexProjectRequired:           policy.CodexProjectRequired,
		CreationAllowedWithoutReview:   policy.RepoCreationAllowedWithoutReview,
		MissingTemplateRoleCount:       len(missingRoles),
		MissingCreationGateCount:       len(missingGates),
		ForbiddenTemplateValueCount:    forbiddenCount,
		TemplateRoles:                  roles,
		CreationGateKeys:               keys,
		BootstrapChecklistReady:        len(checklistMissing) == 0,
		GeneratedCIPresent:             contains(roles, "generated_ci"),
		SecurityScanPresent:            contains(roles, "security_scan"),
		PrivateDataPolicyPresent:       contains(roles, "private_data_policy"),
		BootstrapChecklistPresent:      contains(roles, "bootstrap_checklist"),
		CodexProjectTemplatePresent:    contains(roles, "codex_project"),
		RepoCreationBlockedUntilReview: !policy.RepoCreationAllowedWithoutReview,
		CheckedAt:                      time.Now().UTC().Format(time.RFC3339),
	}
	status.PublicSafe = status.publicSafe()
	return status
}

func (status Status) publicSafe() bool {
	return status.AuthorityReviewRequired &&
		status.PublicSafetyEvidenceRequired &&
		status.CodexProjectRequired &&
		!status.CreationAllowedWithoutReview &&
		status.BootstrapChecklistReady &&
		status.MissingTemplateRoleCount == 0 &&
		status.MissingCreationGateCount == 0 &&
		status.ForbiddenTemplateValueCount == 0
}

func forbiddenTemplateValueCount(policy Policy) int {
	count := 0
	for _, file := range policy.TemplateFiles {
		if forbiddenTemplateValue(file, policy.ForbiddenPublicFragments) {
			count++
		}
	}
	for _, item := range policy.BootstrapChecklist {
		if containsPrivateMarker(item) {
			count++
		}
	}
	return count
}
