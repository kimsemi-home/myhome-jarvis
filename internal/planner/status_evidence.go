package planner

import (
	"strings"

	"github.com/kimsemi-home/myhome-jarvis/internal/knowledge"
	"github.com/kimsemi-home/myhome-jarvis/internal/linear"
)

func attachWriteEvidence(root string, policy Policy, status *Status) error {
	writeEvidence, err := linear.WriteEvidenceStatusForPath(
		root,
		policy.ExternalWriteGate.EvidencePath,
	)
	if err != nil {
		return err
	}
	status.LinearWriteEvidence = writeEvidence
	status.ExternalWriteGate = ExternalWriteGateStatus{
		StandingBoundary:        policy.ExternalWriteGate.StandingBoundary,
		ApprovalRequired:        policy.ExternalWriteGate.ApprovalRequired,
		MutationSuccessRequired: policy.ExternalWriteGate.MutationSuccessRequired,
		BoundaryTaskID:          strings.TrimSpace(policy.ExternalWriteGate.BoundaryTaskID),
		EvidencePath:            writeEvidence.EvidencePath,
	}
	return nil
}

func attachKnowledgeEvidence(root string, policy Policy, status *Status) error {
	if !policy.KnowledgeIndexRequiredBeforePlanning {
		return nil
	}
	query := strings.TrimSpace(policy.KnowledgeIndexDefaultQuery)
	if query == "" {
		query = "planner"
	}
	evidence, err := knowledge.Search(root, query)
	if err != nil {
		return err
	}
	summary := knowledge.SummarizeSearch(evidence)
	status.KnowledgeEvidence = &summary
	return nil
}
