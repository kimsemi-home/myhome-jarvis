package translation

import (
	"encoding/json"
	"os"
)

func inspectManifestFile(path string, policy Policy, status *Status) error {
	body, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var current manifest
	if err := json.Unmarshal(body, &current); err != nil {
		status.InvalidManifestCount++
		return nil
	}
	if err := validateManifest(policy, current); err != nil {
		status.InvalidManifestCount++
		return nil
	}
	recordContext(status, current.SourceContext, current.TargetContext)
	for _, loss := range current.KnownLosses {
		if err := recordManifestLoss(policy, status, current, loss); err != nil {
			status.InvalidManifestCount++
			return nil
		}
	}
	return nil
}

func recordManifestLoss(policy Policy, status *Status, current manifest, loss knownLoss) error {
	status.LossCount++
	lossStatus := "closed"
	if normalizeToken(loss.Level) != "l0_none" {
		lossStatus = "open"
	}
	return recordLoss(
		policy,
		status,
		current.SourceContext,
		current.TargetContext,
		loss.Level,
		loss.Category,
		lossStatus,
	)
}
