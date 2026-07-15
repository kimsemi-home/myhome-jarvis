package shortsfactory

import (
	"errors"
	"fmt"
)

func ValidateContract(value Contract) error {
	if value.SchemaVersion != "shorts.factory-public-contract/v1" || value.ContractID == "" || !value.PublicSafe {
		return errors.New("public Shorts contract identity mismatch")
	}
	if value.LogicalChannelSlots != 20 || value.ApprovalAuthority != "evidence_gate" || value.CriteriaVersion != "evidence-gate/v1" {
		return errors.New("channel count or evidence-gate authority mismatch")
	}
	if value.MinimumIndependentSourcesPerClaim < 2 || value.MinimumPrimarySourcesPerClaim < 1 {
		return errors.New("scientific evidence thresholds are too low")
	}
	if value.MaximumAPIDataRevalidationDays < 1 || value.MaximumAPIDataRevalidationDays > 30 {
		return errors.New("API data revalidation exceeds thirty days")
	}
	if value.DefaultUploadVisibility != "private" || value.ExternalWriteDefault != "deny" || !value.YouTubeConsentReceiptRequired {
		return errors.New("external write defaults are unsafe")
	}
	for _, required := range requiredCriteria {
		if !contains(value.RequiredCriteria, required) {
			return fmt.Errorf("required criterion %q is missing", required)
		}
	}
	if len(value.ReleasedOpenLoopSteps) == 0 || len(value.LinkedPublicContracts) < 6 {
		return errors.New("open-loop or linked contract set is incomplete")
	}
	return nil
}

var requiredCriteria = []string{
	"claim_evidence", "contradiction_review", "uncertainty_disclosure",
	"originality", "asset_rights", "synthetic_content_disclosure",
	"privacy", "youtube_action_consent", "input_hash_integrity",
}

func contains(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}
