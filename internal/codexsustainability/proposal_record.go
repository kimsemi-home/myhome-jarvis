package codexsustainability

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

func RecordProposal(root string, payload []byte) (ProposalRecordResult, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return ProposalRecordResult{}, err
	}
	request, err := decodeProposalRecordRequest(payload)
	if err != nil {
		return ProposalRecordResult{}, err
	}
	record, err := normalizeProposalRecord(policy, request, time.Now().UTC())
	if err != nil {
		return ProposalRecordResult{}, err
	}
	if err := appendRecords(root, policy, []Record{record}); err != nil {
		return ProposalRecordResult{}, err
	}
	status, err := StatusForRoot(root)
	if err != nil {
		return ProposalRecordResult{}, err
	}
	return resultForProposal(record, status), nil
}

func decodeProposalRecordRequest(payload []byte) (ProposalRecordRequest, error) {
	var request ProposalRecordRequest
	decoder := json.NewDecoder(bytes.NewReader(payload))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		return ProposalRecordRequest{}, fmt.Errorf("invalid proposal payload: %w", err)
	}
	return request, nil
}
