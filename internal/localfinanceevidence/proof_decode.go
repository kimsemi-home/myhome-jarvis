package localfinanceevidence

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

func validateProofBody(body []byte, month string, ref ProofRef) error {
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.DisallowUnknownFields()
	var validateErr error
	switch ref.ProofSchema {
	case creditProofSchema:
		var report CreditReport
		if err := decoder.Decode(&report); err != nil {
			return err
		}
		validateErr = validateCreditReport(report, month, ref)
	case portfolioProofSchema:
		var report PortfolioReport
		if err := decoder.Decode(&report); err != nil {
			return err
		}
		validateErr = validatePortfolioReport(report, month, ref)
	case revenueProofSchema:
		var report RevenueReport
		if err := decoder.Decode(&report); err != nil {
			return err
		}
		validateErr = validateRevenueReport(report, month, ref)
	case operatorProofSchema:
		var report OperatorReport
		if err := decoder.Decode(&report); err != nil {
			return err
		}
		validateErr = validateOperatorReport(report, month, ref)
	case shortsProofSchema:
		var report ShortsReport
		if err := decoder.Decode(&report); err != nil {
			return err
		}
		validateErr = validateShortsReport(report, ref)
	case shortsActivationProofSchema:
		var report ShortsActivationReport
		if err := decoder.Decode(&report); err != nil {
			return err
		}
		validateErr = validateShortsActivationReport(report, ref)
	default:
		return errors.New("local finance execution proof schema is unsupported")
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return errors.New("local finance execution proof contains extra JSON")
	}
	return validateErr
}
