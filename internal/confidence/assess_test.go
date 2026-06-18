package confidence

import (
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/evidence"
	"github.com/kimsemi-home/myhome-jarvis/internal/qualitylog"
	"github.com/kimsemi-home/myhome-jarvis/internal/security"
)

func TestAssessReturnsHighWhenEvidenceAndVerificationAreClear(t *testing.T) {
	status := Assess(testPolicy(), clearInputs(evidence.Status{EdgeCount: 2}))

	assertConfidence(t, status, "high", "evidence_backed", false)
}

func TestAssessCapsLowWithoutEvidenceLinks(t *testing.T) {
	status := Assess(testPolicy(), clearInputs(evidence.Status{}))

	assertConfidence(t, status, "low", "missing_evidence_links", false)
}

func TestAssessBlocksOnPublicSafetyFindings(t *testing.T) {
	inputs := clearInputs(evidence.Status{EdgeCount: 2})
	inputs.PublicSafety = security.Status{OK: false}

	status := Assess(testPolicy(), inputs)

	assertConfidence(t, status, "blocked", "public_safety_findings", true)
}

func TestAssessCapsMediumForOpenLearningDebt(t *testing.T) {
	status := Assess(testPolicy(), clearInputs(evidence.Status{
		EdgeCount:         2,
		OpenLearningCount: 1,
	}))

	assertConfidence(t, status, "medium", "open_learning_debt", false)
}

func clearInputs(ev evidence.Status) Inputs {
	return Inputs{
		Evidence: ev,
		Quality:  qualitylog.Status{Exists: true, Last: &qualitylog.Run{OK: true}},
		PublicSafety: security.Status{
			OK: true,
		},
	}
}

func assertConfidence(t *testing.T, status Status, level string, rule string, blocked bool) {
	t.Helper()
	if status.LevelCap != level || status.ActiveRule != rule || status.Blocked != blocked {
		t.Fatalf("status = %#v", status)
	}
}
