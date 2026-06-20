package repofactory

import "github.com/kimsemi-home/myhome-jarvis/internal/security"

func publicSafetyEvidenceUnknown() PublicSafetyEvidence {
	return PublicSafetyEvidence{
		EvidenceState:           "required_before_creation",
		ValidationCommands:      publicSafetyValidationCommands(),
		RawDetailsPublicAllowed: false,
	}
}

func publicSafetyEvidenceFromStatus(
	status security.Status,
) PublicSafetyEvidence {
	state := "blocked_public_safety_findings"
	if status.OK {
		state = "ready"
	}
	return PublicSafetyEvidence{
		OK:                      status.OK,
		CurrentOK:               status.CurrentOK,
		HistoryOK:               status.HistoryOK,
		CurrentFindingCount:     status.CurrentFindingCount,
		HistoryFindingCount:     status.HistoryFindingCount,
		EvidenceState:           state,
		ValidationCommands:      publicSafetyValidationCommands(),
		RawDetailsPublicAllowed: false,
	}
}

func publicSafetyValidationCommands() []string {
	return []string{"mhj security check", "mhj security history"}
}
