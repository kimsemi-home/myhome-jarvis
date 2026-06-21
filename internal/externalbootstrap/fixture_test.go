package externalbootstrap

import (
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/authority"
	"github.com/kimsemi-home/myhome-jarvis/internal/repofactory"
)

func approvalFixture(now time.Time) authority.ApprovalDecisionStatus {
	candidate := "kimsemi-home/myhome-external-evidence-lake"
	return authority.ApprovalDecisionStatus{
		PublicSafe:          true,
		LedgerState:         "present",
		LatestScope:         "repo_creation",
		LatestTarget:        candidate,
		LatestLeaseState:    "active",
		CanCreateRepo:       true,
		ScopeSummaries:      []authority.ApprovalScopeSummary{approvalScope(now)},
		NextSafeAction:      "use_matching_scoped_approval_only",
		CheckedAt:           now.Format(time.RFC3339),
		ActiveApprovalCount: 1,
	}
}

func approvalScope(now time.Time) authority.ApprovalScopeSummary {
	return authority.ApprovalScopeSummary{
		Scope:      "repo_creation",
		Target:     "kimsemi-home/myhome-external-evidence-lake",
		LeaseState: "active",
		ExpiresAt:  now.Add(time.Hour).UTC().Format(time.RFC3339),
		GrantFlags: authority.ApprovalGrantFlags{
			ApprovalGranted:     true,
			RepoCreationAllowed: true,
		},
		CanUnlockScopeOnly: true,
	}
}

func factoryFixture() repofactory.DecisionPacket {
	return repofactory.DecisionPacket{
		PublicSafe:         true,
		CreationDecision:   "blocked_pending_review_evidence",
		TemplateFileCount:  2,
		TemplateReadyCount: 2,
		GateReadyCount:     4,
		CreationGateCount:  5,
		PublicSafetyEvidence: repofactory.PublicSafetyEvidence{
			OK: true, EvidenceState: "ready",
		},
		ContextPackEvidence: contextPackEvidenceFixture(),
		TemplateEvidence:    templateEvidenceFixture(),
	}
}
