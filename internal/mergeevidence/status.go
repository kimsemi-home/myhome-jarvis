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
	publicSafe := policy.PublicStatusRedacted &&
		!policy.MergeWithoutReviewAllowed &&
		!policy.PersistPrivateEvidence &&
		len(missingGates) == 0 &&
		len(missingEvidence) == 0
	return Status{
		Context:                      policy.Context,
		Version:                      policy.Version,
		PolicyPath:                   PolicyRelativePath,
		DefaultBehavior:              policy.DefaultBehavior,
		PublicSafe:                   publicSafe,
		Redaction:                    "policy-summary-only",
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
