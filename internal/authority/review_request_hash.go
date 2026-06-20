package authority

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func reviewRequestID(packet ReviewRequestPacket) string {
	seed := strings.Join([]string{
		packet.PolicyPath,
		packet.RequestState,
		packet.SourceAction,
		strings.Join(packet.RequiredReviewClasses, ","),
		fmt.Sprint(packet.HighRiskBlockedDecisionCount),
		fmt.Sprint(packet.ReviewRequiredDecisionCount),
		fmt.Sprint(packet.ReviewRequiredProfileCount),
		fmt.Sprint(packet.PublicRepoReviewProfileCount),
		fmt.Sprint(packet.WorkflowReviewProfileCount),
		fmt.Sprint(packet.SelfApprovalBlockedProfileCount),
		fmt.Sprint(packet.ExternalWritesAllowedProfileCount),
		fmt.Sprint(packet.ApprovalGranted),
		fmt.Sprint(packet.ExternalWritesAllowed),
		fmt.Sprint(packet.SelfApprovalAllowed),
	}, "|")
	sum := sha256.Sum256([]byte(seed))
	return "authority-review-" + hex.EncodeToString(sum[:])[:12]
}
