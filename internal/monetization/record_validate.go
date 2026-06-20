package monetization

import "fmt"

var errMissingEvidenceRef = fmt.Errorf("monetization evidence ref is required")
var errMissingCostEstimate = fmt.Errorf("monetization cost estimate is required")

func normalizeRecord(policy Policy, record Record) (Record, error) {
	record.ExperimentID = normalizeToken(record.ExperimentID)
	record.HypothesisKey = normalizeToken(record.HypothesisKey)
	record.State = normalizeToken(record.State)
	record.DecisionKind = normalizeToken(record.DecisionKind)
	record.ReviewStatus = normalizeToken(record.ReviewStatus)
	record.ExpectedValueBand = normalizeToken(record.ExpectedValueBand)
	record.CostUnitKind = normalizeToken(record.CostUnitKind)
	if record.At == "" || record.ExperimentID == "" || record.HypothesisKey == "" {
		return Record{}, fmt.Errorf("monetization time, experiment, and hypothesis are required")
	}
	if err := validateRecordEnums(policy, record); err != nil {
		return Record{}, err
	}
	if len(record.EvidenceRefs) == 0 {
		return Record{}, errMissingEvidenceRef
	}
	if record.CostEstimateUnits <= 0 {
		return Record{}, errMissingCostEstimate
	}
	for _, ref := range record.EvidenceRefs {
		if err := validateRef(policy, ref); err != nil {
			return Record{}, err
		}
	}
	return record, nil
}

func validateRecordEnums(policy Policy, record Record) error {
	if !contains(policy.ExperimentStates, record.State) {
		return fmt.Errorf("monetization state %q is not allowed", record.State)
	}
	if !contains(policy.DecisionKinds, record.DecisionKind) {
		return fmt.Errorf("monetization decision %q is not allowed", record.DecisionKind)
	}
	if !contains(policy.ReviewStatuses, record.ReviewStatus) {
		return fmt.Errorf("monetization review status %q is not allowed", record.ReviewStatus)
	}
	if !contains(policy.ExpectedValueBands, record.ExpectedValueBand) {
		return fmt.Errorf("monetization expected value %q is not allowed", record.ExpectedValueBand)
	}
	if !contains(policy.CostUnitKinds, record.CostUnitKind) {
		return fmt.Errorf("monetization cost unit %q is not allowed", record.CostUnitKind)
	}
	return nil
}
