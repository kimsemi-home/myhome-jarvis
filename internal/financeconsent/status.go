package financeconsent

import (
	"errors"
	"os"
	"path/filepath"
	"time"
)

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	status := newStatus(policy)
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(policy.PrivateConsentLedger)))
	if errors.Is(err, os.ErrNotExist) {
		finalizeStatus(&status, map[string]bool{})
		return status, nil
	}
	if err != nil {
		return Status{}, err
	}
	defer file.Close()
	status.Exists = true
	activeKinds, err := scanLedger(file, policy, &status)
	if err != nil {
		return Status{}, err
	}
	finalizeStatus(&status, activeKinds)
	return status, nil
}

func newStatus(policy Policy) Status {
	return Status{
		FinanceMode:                 policy.FinanceMode,
		ForbiddenActionEnabledCount: forbiddenActionCount(policy),
		CheckedAt:                   time.Now().UTC().Format(time.RFC3339),
	}
}

func finalizeStatus(status *Status, activeKinds map[string]bool) {
	status.MissingRequiredConsentCount = len(missingActiveKinds(activeKinds))
	status.ConsentDebtCount = status.InvalidRecordCount +
		status.MissingEvidenceCount +
		status.ReviewRequiredCount +
		status.MissingRequiredConsentCount
	status.ReadinessState = readinessState(*status)
}
