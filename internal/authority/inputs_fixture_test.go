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

func clearInputs() Inputs {
	return Inputs{
		Confidence: confidence.Status{
			LevelCap:       "high",
			PublicSafetyOK: true,
		},
		EvidenceQuality: evidencequality.Status{},
		Incidents:       incidents.Status{},
		ControlPlane:    controlplane.Status{},
		Translation:     translation.Status{},
		Review:          review.Status{CapacityState: "available"},
		PublicSafety:    security.Status{OK: true, CurrentOK: true, HistoryOK: true},
	}
}
