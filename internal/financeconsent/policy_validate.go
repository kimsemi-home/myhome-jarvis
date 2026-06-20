package financeconsent

import (
	"fmt"
	"strings"
)

func validatePolicy(policy Policy) error {
	if policy.Context != "HouseholdFinanceConsentLedger" {
		return fmt.Errorf("finance consent policy context = %q", policy.Context)
	}
	if !strings.HasPrefix(policy.PrivateConsentLedger, "data/private/") ||
		!strings.HasSuffix(policy.PrivateConsentLedger, ".jsonl") {
		return fmt.Errorf("finance consent ledger must stay in data/private JSONL")
	}
	if !policy.AppendOnly || !policy.PublicStatusRedacted {
		return fmt.Errorf("finance consent ledger must be append-only and redacted")
	}
	if !policy.ReadOnly || !policy.ReviewOnly || !policy.FixtureOnlyUntilConsent {
		return fmt.Errorf("finance consent policy must stay read-only and review-only")
	}
	if forbiddenActionCount(policy) > 0 || policy.ExternalWritesAllowed {
		return fmt.Errorf("finance consent policy must not allow write actions")
	}
	if err := validatePolicyRequirements(policy); err != nil {
		return err
	}
	if !contains(policy.Commands, "mhj finance-consent status") {
		return fmt.Errorf("finance consent status command is missing")
	}
	if !contains(policy.Commands, "mhj finance-consent record <json-payload>") {
		return fmt.Errorf("finance consent record command is missing")
	}
	return nil
}

func validatePolicyRequirements(policy Policy) error {
	requirements := []struct {
		label    string
		values   []string
		required []string
	}{
		{"consent kind", policy.ConsentKinds, requiredConsentKinds},
		{"consent status", policy.ConsentStatuses, requiredConsentStatuses},
		{"review status", policy.ReviewStatuses, requiredReviewStatuses},
		{"required field", policy.RequiredFields, requiredFields},
		{"public summary field", policy.PublicSummaryFields, requiredSummaryFields},
		{"authority profile", policy.AuthorityProfiles, []string{"finance_review_only"}},
	}
	for _, item := range requirements {
		if missing := missingValues(item.values, item.required); len(missing) > 0 {
			return fmt.Errorf("finance consent %s missing %q", item.label, missing[0])
		}
	}
	return nil
}

func forbiddenActionCount(policy Policy) int {
	count := 0
	for _, enabled := range []bool{
		policy.TransferActionsAllowed,
		policy.PaymentActionsAllowed,
		policy.TradeActionsAllowed,
		policy.CardActionsAllowed,
	} {
		if enabled {
			count++
		}
	}
	if policy.ExternalWritesAllowed {
		count++
	}
	return count
}
