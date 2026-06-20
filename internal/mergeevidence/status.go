package mergeevidence

import "time"

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	return statusFromPolicy(policy), nil
}

func statusFromPolicy(policy Policy) Status {
	keys := gateKeys(policy.Gates)
	missingGates := missingStrings(keys, requiredGates)
	missingEvidence := missingStrings(policy.RequiredEvidence, requiredEvidence)
	publicSafe := policy.DefaultBehavior == "merge_when_eligible" &&
		policy.MergePreference == "merge_after_checks_pass" &&
		policy.PublicStatusRedacted &&
		!policy.MergeWithoutReviewAllowed &&
		!policy.PersistPrivateEvidence &&
		policy.PostMergeEvidenceRequired &&
		policy.LinearCompletionRequired &&
		policy.MainQualityRunRequired &&
		policy.PrivateDataScanRequired &&
		len(missingGates) == 0 &&
		len(missingEvidence) == 0
	return Status{
		Context:                      policy.Context,
		Version:                      policy.Version,
		PolicyPath:                   PolicyRelativePath,
		DefaultBehavior:              policy.DefaultBehavior,
		MergePreference:              policy.MergePreference,
		PublicSafe:                   publicSafe,
		Redaction:                    "policy-summary-only",
		PostMergeEvidenceRequired:    policy.PostMergeEvidenceRequired,
		LinearCompletionRequired:     policy.LinearCompletionRequired,
		MainQualityRunRequired:       policy.MainQualityRunRequired,
		PrivateDataScanRequired:      policy.PrivateDataScanRequired,
		EligibleGateCount:            len(policy.Gates),
		RequiredEvidenceCount:        len(policy.RequiredEvidence),
		MissingGateCount:             len(missingGates),
		MissingRequiredEvidenceCount: len(missingEvidence),
		MergeReady:                   publicSafe,
		MergeBlockedUntilEvidence:    !publicSafe,
		GateKeys:                     keys,
		RequiredEvidence:             append([]string{}, policy.RequiredEvidence...),
		Commands:                     append([]string{}, policy.Commands...),
		CheckedAt:                    time.Now().UTC().Format(time.RFC3339),
	}
}
