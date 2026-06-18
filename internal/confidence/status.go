package confidence

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/evidence"
	"github.com/kimsemi-home/myhome-jarvis/internal/qualitylog"
	"github.com/kimsemi-home/myhome-jarvis/internal/security"
)

const PolicyRelativePath = "generated/confidence.generated.json"

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	evidenceStatus, err := evidence.StatusForRoot(root)
	if err != nil {
		return Status{}, err
	}
	qualityStatus, err := qualitylog.StatusForRoot(root)
	if err != nil {
		return Status{}, err
	}
	securityStatus, err := security.StatusForRoot(root)
	if err != nil {
		return Status{}, err
	}
	return Assess(policy, Inputs{
		Evidence:     evidenceStatus,
		Quality:      qualityStatus,
		PublicSafety: securityStatus,
	}), nil
}
