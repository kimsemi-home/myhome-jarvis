package knowledge

import "testing"

func TestSearchConnectorReadinessReturnsCatalogEvidence(t *testing.T) {
	assertSearchEvidence(t, searchExpectation{
		query:   "connector readiness",
		concept: "ConnectorCatalog",
		mustRead: []string{
			"generated/connectors.generated.json",
			"docs/connectors.md",
		},
	})
}

func TestSearchAgentClusterReturnsPolicyEvidence(t *testing.T) {
	assertSearchEvidence(t, searchExpectation{
		query:   "agent cluster learning loop",
		concept: "AgentClusterPolicy",
		mustRead: []string{
			"generated/agent_cluster.generated.json",
			"docs/agent-cluster.md",
		},
	})
}

func TestSearchLearningLedgerReturnsObservationEvidence(t *testing.T) {
	assertSearchEvidence(t, searchExpectation{
		query:   "learning ledger",
		concept: "LearningLedger",
		mustRead: []string{
			"generated/learning.generated.json",
			"docs/learning-ledger.md",
		},
	})
}
