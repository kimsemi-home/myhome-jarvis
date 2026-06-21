package authority

import (
	"fmt"
	"strings"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/externalevidence"
)

func fillApprovalDecisionFields(
	record *ApprovalDecisionRecord,
	request ApprovalDecisionRequest,
	packet externalevidence.RepoSplitDecisionPacket,
	now time.Time,
) error {
	packetRef := normalizeToken(request.DecisionPacketRef)
	context := strings.TrimSpace(request.DecisionPacketContext)
	scope := normalizeToken(request.Scope)
	target, err := normalizeApprovalTarget(request.Target)
	if err != nil {
		return err
	}
	reviewer, err := normalizeReviewerBoundary(request.ReviewerBoundary)
	if err != nil {
		return err
	}
	if request.ReviewerIsRequester == nil || *request.ReviewerIsRequester {
		return fmt.Errorf("approval requires independent human reviewer")
	}
	checkedAt, err := normalizePacketCheckedAt(request.DecisionPacketCheckedAt, now)
	if err != nil {
		return err
	}
	record.DecisionPacketRef = packetRef
	record.DecisionPacketContext = context
	record.DecisionPacketCheckedAt = checkedAt
	record.Scope = scope
	record.Target = target
	record.ReviewerBoundary = reviewer
	record.ReviewerIsRequester = false
	return validateRepoSplitPacketRef(*record, packet)
}
