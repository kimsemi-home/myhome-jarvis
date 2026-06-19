package knowledge

import "testing"

func TestSearchConfidenceAssessorReturnsGateEvidence(t *testing.T) {
	assertSearchEvidence(t, searchExpectation{
		query:   "confidence assessor",
		concept: "ConfidenceAssessor",
		mustRead: []string{
			"generated/confidence.generated.json",
			"docs/confidence-assessor.md",
		},
	})
}

func TestSearchTranslationManifestReturnsLossLedgerEvidence(t *testing.T) {
	assertSearchEvidence(t, searchExpectation{
		query:   "translation manifest loss ledger",
		concept: "TranslationManifest",
		mustRead: []string{
			"generated/translation.generated.json",
			"docs/translation-manifest.md",
		},
	})
}

func TestSearchControlPlaneManifestReturnsRoutingEvidence(t *testing.T) {
	assertSearchEvidence(t, searchExpectation{
		query:   "control plane manifest orchestration",
		concept: "ControlPlaneManifest",
		mustRead: []string{
			"generated/control_plane.generated.json",
			"docs/control-plane-manifest.md",
		},
	})
}

func TestSearchAuthorityGateReturnsRBACEvidence(t *testing.T) {
	assertSearchEvidence(t, searchExpectation{
		query:   "authority gate reasoning rbac domain abac",
		concept: "AuthorityGate",
		mustRead: []string{
			"generated/authority.generated.json",
			"docs/authority-gate.md",
		},
	})
}
