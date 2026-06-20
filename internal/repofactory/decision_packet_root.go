package repofactory

import "github.com/kimsemi-home/myhome-jarvis/internal/security"

func DecisionPacketForRoot(root string) (DecisionPacket, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return DecisionPacket{}, err
	}
	status, err := security.StatusForRoot(root)
	if err != nil {
		return DecisionPacket{}, err
	}
	return decisionPacketFromPolicyEvidence(
		policy,
		publicSafetyEvidenceFromStatus(status),
	), nil
}
