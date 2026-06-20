package commandcenter

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/evidence"
	"github.com/kimsemi-home/myhome-jarvis/internal/incidents"
	"github.com/kimsemi-home/myhome-jarvis/internal/pdca"
)

func summarizeVision(policy visionPolicy) VisionSummary {
	return VisionSummary{
		PolicyPath:      visionPolicyPath,
		Mission:         policy.Mission,
		OperatingMode:   policy.OperatingMode,
		CapabilityCount: len(policy.CapabilityPillars),
		GuardrailCount:  len(policy.Guardrails),
	}
}

func summarizePDCA(status pdca.Status) PDCASummary {
	return PDCASummary{
		Ready:                status.Ready,
		CycleCount:           status.CycleCount,
		ReadyStepCount:       status.ReadyStepCount,
		MissingArtifactCount: status.MissingArtifactCount,
	}
}

func summarizeEvidence(status evidence.Status) EvidenceSummary {
	return EvidenceSummary{
		SourceCount:              status.SourceCount,
		PresentSourceCount:       status.PresentSourceCount,
		NodeCount:                status.NodeCount,
		EdgeCount:                status.EdgeCount,
		DanglingEvidenceRefCount: status.DanglingEvidenceRefCount,
		OpenLearningCount:        status.OpenLearningCount,
	}
}

func summarizeIncidents(status incidents.Status) IncidentSummary {
	return IncidentSummary{
		OpenCount:               status.OpenCount,
		IncidentDebtCount:       status.IncidentDebtCount,
		StaleQuarantineCount:    status.StaleQuarantineCount,
		MissingEvidenceRefCount: status.MissingEvidenceRefCount,
	}
}
