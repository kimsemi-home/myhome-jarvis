package externalevidence

import (
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/contextpack"
)

func RepoSplitDecisionPacketForRoot(root string) (RepoSplitDecisionPacket, error) {
	policy, err := readPolicy(root)
	if err != nil {
		return RepoSplitDecisionPacket{}, err
	}
	if err := ValidatePolicy(policy); err != nil {
		return RepoSplitDecisionPacket{}, err
	}
	contextStatus, err := contextpack.StatusForRoot(root)
	if err != nil {
		return RepoSplitDecisionPacket{}, err
	}
	return repoSplitDecisionPacketFromPolicy(
		policy,
		contextStatus,
		time.Now().UTC(),
	), nil
}

func repoSplitDecisionPacketFromPolicy(
	policy Policy,
	contextStatus contextpack.Status,
	at time.Time,
) RepoSplitDecisionPacket {
	assessment := policy.RepoSplitAssessment
	return RepoSplitDecisionPacket{
		Context:                         "ExternalEvidenceRepoSplitDecisionPacket",
		SchemaVersion:                   policy.SchemaVersion,
		PublicSafe:                      repoSplitDecisionPublicSafe(policy, contextStatus),
		DecisionState:                   "review_only",
		RecommendedOption:               assessment.Recommendation,
		FutureRepoCandidate:             assessment.FutureRepoCandidate,
		RepoCreationGate:                assessment.CreationGate,
		AuthorityDecisionRecordRequired: true,
		CanCreateRepo:                   false,
		PrivateLakeStaysPrivate:         true,
		RawPayloadPublicAllowed:         false,
		ExternalWritesAllowed:           false,
		ContextPackVersion:              contextStatus.Version,
		OntologyVersion:                 contextStatus.OntologyVersion,
		Options:                         repoSplitDecisionOptions(assessment),
		EvidenceChecks:                  repoSplitDecisionEvidenceChecks(policy, contextStatus),
		ForbiddenGrantFlags:             RepoSplitDecisionForbiddenGrants{},
		NextSafeAction:                  "record_authority_review_before_repo_creation",
		CheckedAt:                       at.Format(time.RFC3339),
	}
}

func repoSplitDecisionPublicSafe(policy Policy, contextStatus contextpack.Status) bool {
	return policy.PublicSafe && !policy.RawPayloadPublicAllowed &&
		!policy.CredentialsAllowed && !policy.CookiesAllowed &&
		contextStatus.PublicSafe &&
		policy.RepoSplitAssessment.CreationGate == "authority_review_required"
}
