package commandcenter

import "testing"

func TestAuthorityReviewDecisionContractChecksExternalEvidence(t *testing.T) {
	packet, err := AuthorityReviewDecisionPacketForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	check := contractEvidenceCheckByKey(
		packet.DecisionContract.RequiredEvidenceChecks,
		"external_evidence",
	)
	if !check.Required || !check.PublicSafe || check.State == "" {
		t.Fatalf("external evidence contract check = %#v", check)
	}
}

func readyExternalEvidenceSummary() ExternalEvidenceSummary {
	return ExternalEvidenceSummary{
		PublicSafe:                  true,
		EvidenceReady:               true,
		ExternalCollectionAllowed:   true,
		ForbiddenCollectionDisabled: true,
		RawPayloadPublicAllowed:     false,
		SourceCount:                 4,
		SourceClassCount:            5,
		ManifestPresent:             true,
		ManifestRecordCount:         2,
		ArchiveSourceKey:            "external_evidence",
		RepoSplitRecommendation:     "keep_contract_in_myhome_jarvis_defer_repo_creation",
		RepoCreationGate:            "authority_review_required",
		SplitTriggerCount:           4,
	}
}

func assertReadyExternalEvidenceSummary(
	t *testing.T,
	summary ExternalEvidenceSummary,
) {
	t.Helper()
	if !summary.PublicSafe ||
		summary.RawPayloadPublicAllowed ||
		summary.ArchiveSourceKey != "external_evidence" ||
		summary.RepoCreationGate != "authority_review_required" {
		t.Fatalf("external evidence summary = %#v", summary)
	}
}

func contractEvidenceCheckByKey(
	checks []AuthorityReviewContractEvidenceCheck,
	key string,
) AuthorityReviewContractEvidenceCheck {
	for _, check := range checks {
		if check.Key == key {
			return check
		}
	}
	return AuthorityReviewContractEvidenceCheck{}
}
