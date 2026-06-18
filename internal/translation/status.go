package translation

import "time"

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	status := newStatus(policy)
	if err := inspectLedger(root, policy, &status); err != nil {
		return Status{}, err
	}
	if err := inspectManifests(root, policy, &status); err != nil {
		return Status{}, err
	}
	status.OpenDebtCount = status.OpenLossCount +
		status.InvalidLossCount +
		status.InvalidManifestCount +
		status.MissingManifestCount
	return status, nil
}

func newStatus(policy Policy) Status {
	return Status{
		PolicyPath:      PolicyRelativePath,
		LedgerPath:      policy.PrivateLossLedger,
		ManifestRoot:    policy.PrivateManifestRoot,
		ByLevel:         map[string]int{},
		BySourceContext: map[string]int{},
		ByTargetContext: map[string]int{},
		CheckedAt:       time.Now().UTC().Format(time.RFC3339),
	}
}
