package authority

import (
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/externalevidence"
)

func approvalFixturePacket(now time.Time) externalevidence.RepoSplitDecisionPacket {
	return externalevidence.RepoSplitDecisionPacket{
		Context:             "ExternalEvidenceRepoSplitDecisionPacket",
		PublicSafe:          true,
		DecisionState:       "review_only",
		FutureRepoCandidate: "kimsemi-home/myhome-external-evidence-lake",
		RepoCreationGate:    "authority_review_required",
		CanCreateRepo:       false,
		CheckedAt:           now.UTC().Format(time.RFC3339),
	}
}

func approvalFixtureRequest(now time.Time) ApprovalDecisionRequest {
	trueValue := true
	falseValue := false
	return ApprovalDecisionRequest{
		At:                      now.UTC().Format(time.RFC3339),
		DecisionPacketRef:       "external_evidence_repo_split_decision",
		DecisionPacketContext:   "ExternalEvidenceRepoSplitDecisionPacket",
		DecisionPacketCheckedAt: now.UTC().Format(time.RFC3339),
		Scope:                   "repo_creation",
		Target:                  "kimsemi-home/myhome-external-evidence-lake",
		ReviewerBoundary:        "human_governance_steward",
		ReviewerIsRequester:     &falseValue,
		ExpiresAt:               now.Add(time.Hour).UTC().Format(time.RFC3339),
		ApprovalGranted:         &trueValue,
		RepoCreationAllowed:     &trueValue,
		WorkflowChangesAllowed:  &falseValue,
		ExternalWritesAllowed:   &falseValue,
		SelfApprovalAllowed:     &falseValue,
	}
}
