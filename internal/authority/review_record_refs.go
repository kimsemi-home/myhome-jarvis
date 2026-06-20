package authority

import (
	"fmt"
	"strings"
)

func validateReviewRecordRefs(
	requestID string,
	evidenceRef string,
	queueItemRef string,
	packet ReviewRequestPacket,
	evidence ReviewRequestEvidenceStatus,
	queue ReviewQueueStatus,
) error {
	if !validReviewRequestID(requestID) || requestID != packet.RequestID {
		return fmt.Errorf("authority review request id does not match current packet")
	}
	if evidenceRef != evidence.EvidenceRef ||
		evidenceRef != "authority_review_request:"+requestID {
		return fmt.Errorf("authority review evidence ref does not match current packet")
	}
	if queueItemRef != queue.QueueItemRef ||
		queueItemRef != "authority_review_queue:"+requestID {
		return fmt.Errorf("authority review queue item ref does not match current packet")
	}
	return nil
}

func validReviewRequestID(value string) bool {
	const prefix = "authority-review-"
	if !strings.HasPrefix(value, prefix) || len(value) != len(prefix)+12 {
		return false
	}
	for _, char := range value[len(prefix):] {
		if (char < '0' || char > '9') && (char < 'a' || char > 'f') {
			return false
		}
	}
	return true
}
