package authority

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

func RecordReviewRequest(root string, payload []byte) (ReviewRecordResult, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return ReviewRecordResult{}, err
	}
	packet, err := ReviewRequestPacketForRoot(root)
	if err != nil {
		return ReviewRecordResult{}, err
	}
	evidence := ReviewRequestEvidenceFromPacket(packet)
	queue := ReviewQueueStatusFromPacket(packet, evidence)
	request, err := decodeReviewRecordRequest(payload)
	if err != nil {
		return ReviewRecordResult{}, err
	}
	record, err := normalizeReviewRecordRequest(policy, request, packet, evidence, queue, time.Now().UTC())
	if err != nil {
		return ReviewRecordResult{}, err
	}
	if err := appendReviewRecord(root, policy, record); err != nil {
		return ReviewRecordResult{}, err
	}
	return resultForReviewRecord(record), nil
}

func decodeReviewRecordRequest(payload []byte) (ReviewRecordRequest, error) {
	var request ReviewRecordRequest
	decoder := json.NewDecoder(bytes.NewReader(payload))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		return ReviewRecordRequest{}, fmt.Errorf("invalid authority review record payload: %w", err)
	}
	return request, nil
}
