package shortsfactory

import (
	"errors"
	"fmt"
	"sort"
)

func Evaluate(request GateRequest, contract Contract) (GateResult, error) {
	if err := ValidateContract(contract); err != nil {
		return GateResult{}, err
	}
	if request.SchemaVersion != "shorts.gate-request/v1" || request.JobID == "" || request.State != "qa_ready" {
		return GateResult{}, errors.New("gate request envelope is invalid")
	}
	criteria := evaluateCriteria(request, contract)
	inputHash, err := hashJSON(request)
	if err != nil {
		return GateResult{}, err
	}
	result := GateResult{SchemaVersion: "shorts.gate-result/v1", Decision: "approved",
		CriteriaVersion: contract.CriteriaVersion, InputHash: inputHash, Criteria: criteria}
	for _, item := range criteria {
		if !item.Passed {
			result.Decision = "rejected"
		}
	}
	if result.Decision == "approved" {
		result.ReceiptRef = "gate:sha256:" + mustHash(result)
		result.ReleasedOpenLoopSteps = append([]string{}, contract.ReleasedOpenLoopSteps...)
		sort.Strings(result.ReleasedOpenLoopSteps)
	}
	return result, nil
}

func Verify(root string) (GateResult, error) {
	contract, err := LoadContract(root + "/contracts/shorts-factory.json")
	if err != nil {
		return GateResult{}, err
	}
	request, err := LoadRequest(root + "/fixtures/shorts_gate_pass.json")
	if err != nil {
		return GateResult{}, err
	}
	result, err := Evaluate(request, contract)
	if err != nil {
		return result, err
	}
	if result.Decision != "approved" || result.ReceiptRef == "" || len(result.ReleasedOpenLoopSteps) == 0 {
		return result, fmt.Errorf("public-safe gate fixture was not approved")
	}
	return result, nil
}
