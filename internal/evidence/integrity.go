package evidence

import "time"

func IntegrityForRoot(root string) (IntegrityStatus, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return IntegrityStatus{}, err
	}
	status := newIntegrityStatus(policy, time.Now().UTC())
	if err := inspectLearningIntegrity(root, policy, &status); err != nil {
		return IntegrityStatus{}, err
	}
	if status.DanglingEvidenceRefCount > 0 {
		status.NextSafeAction = "repair_private_learning_refs"
	}
	return status, nil
}

func newIntegrityStatus(policy Policy, checked time.Time) IntegrityStatus {
	status := IntegrityStatus{
		PolicyPath:     PolicyRelativePath,
		PrivateRoot:    policy.PrivateRoot,
		PublicSafe:     true,
		Redaction:      "prefix-counts-only",
		PrefixCounts:   make([]IntegrityPrefix, 0, len(policy.AllowedEvidencePrefixes)),
		NextSafeAction: "none",
		CheckedAt:      checked.Format(time.RFC3339),
	}
	for _, prefix := range policy.AllowedEvidencePrefixes {
		status.PrefixCounts = append(status.PrefixCounts, IntegrityPrefix{Prefix: prefix})
	}
	return status
}
