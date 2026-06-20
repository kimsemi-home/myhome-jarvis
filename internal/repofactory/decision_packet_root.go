package repofactory

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/contextpack"
	"github.com/kimsemi-home/myhome-jarvis/internal/security"
)

func DecisionPacketForRoot(root string) (DecisionPacket, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return DecisionPacket{}, err
	}
	status, err := security.StatusForRoot(root)
	if err != nil {
		return DecisionPacket{}, err
	}
	verify, err := contextpack.VerifyDeclarationForRoot(root, "")
	if err != nil {
		return DecisionPacket{}, err
	}
	contextStatus, err := contextpack.StatusForRoot(root)
	if err != nil {
		return DecisionPacket{}, err
	}
	return decisionPacketFromEvidence(
		policy,
		publicSafetyEvidenceFromStatus(status),
		contextPackEvidenceFromStatus(contextStatus, verify),
	), nil
}
