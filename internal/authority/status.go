package authority

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/confidence"
	"github.com/kimsemi-home/myhome-jarvis/internal/controlplane"
	"github.com/kimsemi-home/myhome-jarvis/internal/evidencequality"
	"github.com/kimsemi-home/myhome-jarvis/internal/incidents"
	"github.com/kimsemi-home/myhome-jarvis/internal/review"
	"github.com/kimsemi-home/myhome-jarvis/internal/security"
	"github.com/kimsemi-home/myhome-jarvis/internal/translation"
)

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	inputs, err := inputsForRoot(root)
	if err != nil {
		return Status{}, err
	}
	return Assess(policy, inputs), nil
}

func inputsForRoot(root string) (Inputs, error) {
	confidenceStatus, err := confidence.StatusForRoot(root)
	if err != nil {
		return Inputs{}, err
	}
	evidenceQualityStatus, err := evidencequality.StatusForRoot(root)
	if err != nil {
		return Inputs{}, err
	}
	incidentStatus, err := incidents.StatusForRoot(root)
	if err != nil {
		return Inputs{}, err
	}
	controlPlaneStatus, err := controlplane.StatusForRoot(root)
	if err != nil {
		return Inputs{}, err
	}
	translationStatus, err := translation.StatusForRoot(root)
	if err != nil {
		return Inputs{}, err
	}
	reviewStatus, err := review.StatusForRoot(root)
	if err != nil {
		return Inputs{}, err
	}
	publicSafetyStatus, err := security.StatusForRoot(root)
	if err != nil {
		return Inputs{}, err
	}
	return Inputs{
		Confidence:      confidenceStatus,
		EvidenceQuality: evidenceQualityStatus,
		Incidents:       incidentStatus,
		ControlPlane:    controlPlaneStatus,
		Translation:     translationStatus,
		Review:          reviewStatus,
		PublicSafety:    publicSafetyStatus,
	}, nil
}
