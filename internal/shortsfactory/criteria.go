package shortsfactory

import "regexp"

var slotPattern = regexp.MustCompile(`^channel-(0[1-9]|1[0-9]|20)$`)
var hashPattern = regexp.MustCompile(`^[a-f0-9]{64}$`)

func evaluateCriteria(value GateRequest, contract Contract) []Criterion {
	evidence := value.ClaimEvidence
	originality := value.Originality
	privacy := value.Privacy
	return []Criterion{
		{ID: "logical-channel", Passed: slotPattern.MatchString(value.LogicalChannelSlot)},
		{ID: "claim-evidence", Passed: evidence.IndependentSources >= contract.MinimumIndependentSourcesPerClaim &&
			evidence.PrimarySources >= contract.MinimumPrimarySourcesPerClaim && evidence.FreshSources >= contract.MinimumIndependentSourcesPerClaim},
		{ID: "contradiction-review", Passed: evidence.ContradictionReviewed && evidence.UnresolvedContradictions == 0},
		{ID: "uncertainty-disclosure", Passed: evidence.UncertaintyDisclosed},
		{ID: "originality", Passed: originality.OriginalScript && originality.OriginalAnalysis && !originality.CrossChannelDuplicate && !originality.TemplateOnly},
		{ID: "asset-rights", Passed: value.Rights.AllAssetsCleared && value.Rights.EvidenceCount > 0},
		{ID: "synthetic-disclosure", Passed: !value.SyntheticContent.RealisticOrMeaningful ||
			(value.SyntheticContent.DisclosureRequired && value.SyntheticContent.DisclosurePlanned)},
		{ID: "privacy", Passed: !privacy.ContainsCredentials && !privacy.ContainsAccountIdentifiers &&
			!privacy.ContainsRawPrivateEvidence && !privacy.ContainsRevenueDetails},
		{ID: "youtube-consent", Passed: value.YouTubeConsent.Required && value.YouTubeConsent.ReceiptRef != "" &&
			value.YouTubeConsent.Visibility == contract.DefaultUploadVisibility},
		{ID: "input-hash-integrity", Passed: validHashes(value.InputHashes)},
	}
}

func validHashes(value InputHashes) bool {
	items := []string{value.Content, value.Evidence, value.Assets, value.ExecutionPlan, value.ChannelBinding, value.ConsentReceipt}
	for _, item := range items {
		if !hashPattern.MatchString(item) {
			return false
		}
	}
	return true
}
