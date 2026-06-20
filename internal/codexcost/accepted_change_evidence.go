package codexcost

import "github.com/kimsemi-home/myhome-jarvis/internal/codexsustainability"

type acceptedChangeEvidence struct {
	AcceptedCount int64
	LedgerCount   int64
	MergeCount    int64
	Source        string
	LogLimit      int
}

func acceptedChangeEvidenceFor(
	sustainability codexsustainability.Status,
	merge mergeAcceptance,
) acceptedChangeEvidence {
	evidence := acceptedChangeEvidence{
		LedgerCount: sustainability.AcceptedChangeCount,
		MergeCount:  merge.Count,
		LogLimit:    merge.Limit,
		Source:      acceptedChangeSource(sustainability.AcceptedChangeCount, merge),
	}
	evidence.AcceptedCount = maxInt64(evidence.LedgerCount, evidence.MergeCount)
	return evidence
}

func acceptedChangeSource(ledgerCount int64, merge mergeAcceptance) string {
	if ledgerCount > 0 && merge.Count > 0 {
		return "codex_sustainability_ledger_and_git_merge_commits"
	}
	if merge.Count > 0 {
		return merge.Source
	}
	if ledgerCount > 0 {
		return "codex_sustainability_ledger"
	}
	if merge.Source == "unavailable" {
		return "unavailable"
	}
	return "missing"
}

func maxInt64(left int64, right int64) int64 {
	if left > right {
		return left
	}
	return right
}
