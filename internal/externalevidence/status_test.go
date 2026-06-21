package externalevidence

import "testing"

func TestStatusSummarizesExternalEvidencePolicy(t *testing.T) {
	status, err := StatusForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	if !status.PublicSafe || status.RawPayloadPublicAllowed ||
		status.CredentialsAllowed || status.CookiesAllowed {
		t.Fatalf("external evidence safety = %#v", status)
	}
	if status.SourceCount < 4 || status.ArchiveSourceKey != "external_evidence" {
		t.Fatalf("external evidence source summary = %#v", status)
	}
	if status.RepoCreationGate != "authority_review_required" ||
		status.FutureRepoCandidate == "" || status.SplitTriggerCount == 0 {
		t.Fatalf("repo split assessment = %#v", status)
	}
	if !contains(status.PublicRepoRules, "no_raw_payloads") ||
		!contains(status.Commands, "mhj external-evidence repo-split-decision") ||
		!contains(status.Commands, "mhj external-evidence collect") {
		t.Fatalf("public rules or commands missing = %#v", status)
	}
}
