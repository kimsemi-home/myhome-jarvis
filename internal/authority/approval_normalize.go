package authority

import (
	"fmt"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/externalevidence"
)

const approvalPacketFreshnessHours = 24

func normalizeApprovalDecisionRequest(
	policy Policy,
	request ApprovalDecisionRequest,
	packet externalevidence.RepoSplitDecisionPacket,
	now time.Time,
) (ApprovalDecisionRecord, error) {
	if !privateJSONL(policy.PrivateApprovalDecisionLedger) {
		return ApprovalDecisionRecord{}, fmt.Errorf("approval ledger must stay private jsonl")
	}
	at, err := normalizeReviewRecordTime(request.At, now)
	if err != nil {
		return ApprovalDecisionRecord{}, err
	}
	flags, err := normalizeApprovalGrantFlags(request)
	if err != nil {
		return ApprovalDecisionRecord{}, err
	}
	expiresAt, err := normalizeApprovalExpiry(request.ExpiresAt, now)
	if err != nil {
		return ApprovalDecisionRecord{}, err
	}
	record := ApprovalDecisionRecord{At: at, GrantFlags: flags,
		ExpiresAt: expiresAt, LeaseState: "active", PublicSafe: true}
	if err := fillApprovalDecisionFields(&record, request, packet, now); err != nil {
		return ApprovalDecisionRecord{}, err
	}
	if err := validateApprovalScope(record, packet); err != nil {
		return ApprovalDecisionRecord{}, err
	}
	return record, nil
}
