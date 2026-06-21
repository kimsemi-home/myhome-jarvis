package externalevidence

import "testing"

func TestRepoSplitDecisionPacketRequiresReviewBeforeCreation(t *testing.T) {
	packet, err := RepoSplitDecisionPacketForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	if !packet.PublicSafe || packet.DecisionState != "review_only" {
		t.Fatalf("packet boundary = %#v", packet)
	}
	if packet.CanCreateRepo || packet.ExternalWritesAllowed ||
		packet.RawPayloadPublicAllowed || !packet.PrivateLakeStaysPrivate {
		t.Fatalf("packet grants unsafe authority = %#v", packet)
	}
	if packet.RepoCreationGate != "authority_review_required" ||
		!packet.AuthorityDecisionRecordRequired {
		t.Fatalf("review gate = %#v", packet)
	}
	if packet.FutureRepoCandidate !=
		"kimsemi-home/myhome-external-evidence-lake" {
		t.Fatalf("future repo candidate = %q", packet.FutureRepoCandidate)
	}
}

func TestRepoSplitDecisionPacketComparesOptions(t *testing.T) {
	packet, err := RepoSplitDecisionPacketForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	keep := repoSplitOption(packet, "keep_contract_in_myhome_jarvis")
	split := repoSplitOption(packet, "split_new_external_evidence_lake_repo")
	if keep.Key == "" || split.Key == "" {
		t.Fatalf("options = %#v", packet.Options)
	}
	if keep.RepoCreationAllowed || split.RepoCreationAllowed ||
		!split.HumanApprovalRequired {
		t.Fatalf("option authority = %#v %#v", keep, split)
	}
	if split.ContextPackHandoff == "" || split.GitHubActionsCost == "" ||
		split.ArchiveCacheBehavior == "" ||
		split.OntologyVersionDiscovery == "" {
		t.Fatalf("split comparison = %#v", split)
	}
}

func repoSplitOption(
	packet RepoSplitDecisionPacket,
	key string,
) RepoSplitDecisionOption {
	for _, option := range packet.Options {
		if option.Key == key {
			return option
		}
	}
	return RepoSplitDecisionOption{}
}
