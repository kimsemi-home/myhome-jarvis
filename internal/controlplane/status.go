package controlplane

import "time"

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	status := newStatus(policy)
	if err := scanLedger(root, policy, &status); err != nil {
		return Status{}, err
	}
	status.ManifestDebtCount = status.InvalidManifestCount + status.VerifierViolationCount
	return status, nil
}

func newStatus(policy Policy) Status {
	return Status{
		PolicyPath:                 PolicyRelativePath,
		ManifestPath:               policy.PrivateManifestLedger,
		VerifierSeparationRequired: policy.VerifierSeparationRequired,
		MinLeaseSeconds:            policy.MinLeaseSeconds,
		MaxLeaseSeconds:            policy.MaxLeaseSeconds,
		ByDecisionKind:             map[string]int{},
		ByAuthorityProfile:         map[string]int{},
		ByLeaseStatus:              map[string]int{},
		CheckedAt:                  time.Now().UTC().Format(time.RFC3339),
	}
}
