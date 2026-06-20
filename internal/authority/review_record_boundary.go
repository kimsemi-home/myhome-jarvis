package authority

import "fmt"

func validateReviewRecordPublicBoundary(
	packet ReviewRequestPacket,
	evidence ReviewRequestEvidenceStatus,
	queue ReviewQueueStatus,
) error {
	if packet.RequestState != "ready" ||
		!packet.ReviewRequestable ||
		!evidence.EvidenceReady ||
		!queue.QueueReady {
		return fmt.Errorf("authority review request is not ready to record")
	}
	if !packet.PublicSafe || !evidence.PublicSafe || !queue.PublicSafe {
		return fmt.Errorf("authority review request is not public-safe")
	}
	if packet.ApprovalGranted || evidence.ApprovalGranted || queue.ApprovalGranted ||
		packet.ExternalWritesAllowed || evidence.ExternalWritesAllowed || queue.ExternalWritesAllowed ||
		packet.SelfApprovalAllowed || evidence.SelfApprovalAllowed || queue.SelfApprovalAllowed {
		return fmt.Errorf("authority review request must not grant approval or external writes")
	}
	return nil
}
