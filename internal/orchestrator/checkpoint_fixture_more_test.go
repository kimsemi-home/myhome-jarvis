package orchestrator

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/knowledge"
	"github.com/kimsemi-home/myhome-jarvis/internal/security"
)

func redactedKnowledgeEvidence() *knowledge.Evidence {
	return &knowledge.Evidence{
		Query:        "planner KnowledgeIndex Linear closed loop",
		ConceptCount: 3,
		HitCount:     9,
		LinearIssues: []string{"KIM-14"},
		MustRead:     []string{"generated/concepts.generated.json", "docs/knowledge-index.md"},
		CheckedAt:    "2026-06-15T00:00:00Z",
	}
}

func redactedSecurityStatus() security.Status {
	return security.Status{
		OK:                  true,
		CurrentOK:           true,
		HistoryOK:           true,
		CurrentFindingCount: 0,
		HistoryFindingCount: 0,
		CheckedAt:           "2026-06-15T00:00:00Z",
	}
}
