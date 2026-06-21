package externalevidence

import "testing"

func TestRepoSplitDecisionPacketHasRequiredEvidence(t *testing.T) {
	packet, err := RepoSplitDecisionPacketForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{
		"external_evidence_policy",
		"authority_gate",
		"public_repo_rules",
		"context_pack_handoff",
		"archive_cache_contract",
		"private_payload_boundary",
	} {
		check := repoSplitEvidenceCheck(packet, key)
		if !check.Required || !check.PublicSafe || check.State == "" {
			t.Fatalf("evidence check %s = %#v", key, check)
		}
	}
	if packet.ContextPackVersion == "" || packet.OntologyVersion == "" {
		t.Fatalf("context versions = %#v", packet)
	}
}

func repoSplitEvidenceCheck(
	packet RepoSplitDecisionPacket,
	key string,
) RepoSplitDecisionEvidenceCheck {
	for _, check := range packet.EvidenceChecks {
		if check.Key == key {
			return check
		}
	}
	return RepoSplitDecisionEvidenceCheck{}
}
