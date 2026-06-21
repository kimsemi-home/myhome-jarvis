package externalevidence

import "github.com/kimsemi-home/myhome-jarvis/internal/contextpack"

func repoSplitDecisionEvidenceChecks(
	policy Policy,
	contextStatus contextpack.Status,
) []RepoSplitDecisionEvidenceCheck {
	assessment := policy.RepoSplitAssessment
	return []RepoSplitDecisionEvidenceCheck{
		decisionCheck("external_evidence_policy", policy.SchemaVersion, true,
			policy.PublicSafe),
		decisionCheck("authority_gate", assessment.CreationGate, true,
			assessment.CreationGate == "authority_review_required"),
		decisionCheck("public_repo_rules", repoRuleState(assessment), true,
			repoRulesPublicSafe(assessment)),
		decisionCheck("context_pack_handoff", contextStatus.Version, true,
			contextStatus.PublicSafe && contextStatus.OntologyVersion != ""),
		decisionCheck("archive_cache_contract", policy.ArchiveSourceKey, true,
			policy.ArchiveSourceKey == "external_evidence"),
		decisionCheck("private_payload_boundary", privatePayloadState(policy), true,
			!policy.RawPayloadPublicAllowed),
	}
}

func decisionCheck(
	key string,
	state string,
	required bool,
	publicSafe bool,
) RepoSplitDecisionEvidenceCheck {
	return RepoSplitDecisionEvidenceCheck{Key: key, State: state,
		Required: required, PublicSafe: publicSafe}
}
