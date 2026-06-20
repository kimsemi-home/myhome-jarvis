package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/evidence"

type evidenceStatus = evidence.Status
type evidenceIntegrityStatus = evidence.IntegrityStatus

func collectEvidenceInputs(root string, in *inputs) error {
	var err error
	if in.Evidence, err = evidence.StatusForRoot(root); err != nil {
		return err
	}
	if in.EvidenceIntegrity, err = evidence.IntegrityForRoot(root); err != nil {
		return err
	}
	return nil
}
