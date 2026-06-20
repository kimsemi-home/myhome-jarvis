package commandcenter

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/codexcost"
	"github.com/kimsemi-home/myhome-jarvis/internal/codexsustainability"
	"github.com/kimsemi-home/myhome-jarvis/internal/contextpack"
	"github.com/kimsemi-home/myhome-jarvis/internal/financeconsent"
	"github.com/kimsemi-home/myhome-jarvis/internal/incidents"
	"github.com/kimsemi-home/myhome-jarvis/internal/monetization"
	"github.com/kimsemi-home/myhome-jarvis/internal/pdca"
	"github.com/kimsemi-home/myhome-jarvis/internal/review"
	"github.com/kimsemi-home/myhome-jarvis/internal/storagearchive"
)

type inputs struct {
	Vision              visionPolicy
	PDCA                pdca.Status
	Evidence            evidenceStatus
	EvidenceIntegrity   evidenceIntegrityStatus
	Incidents           incidents.Status
	Authority           authorityStatus
	AuthorityReview     authorityReviewPlanStatus
	Review              review.Status
	FinanceConsent      financeconsent.Status
	Cost                codexcost.Status
	CostBrief           codexcost.Brief
	CostScaling         codexcost.ScalingPacket
	CodexSustainability codexsustainability.Status
	StorageArchive      storagearchive.Status
	ContextPack         contextpack.Status
	MediaReadiness      mediaReadinessStatus
	MergeEvidence       mergeEvidenceStatus
	Monetization        monetization.Status
	RepoFactory         repoFactoryStatus
}

func collectInputs(root string) (inputs, error) {
	var err error
	var in inputs
	if in.Vision, err = readVisionPolicy(root); err != nil {
		return inputs{}, err
	}
	if in.PDCA, err = pdca.StatusForRoot(root); err != nil {
		return inputs{}, err
	}
	if err = collectEvidenceInputs(root, &in); err != nil {
		return inputs{}, err
	}
	if in.Incidents, err = incidents.StatusForRoot(root); err != nil {
		return inputs{}, err
	}
	if err = collectAuthorityInputs(root, &in); err != nil {
		return inputs{}, err
	}
	if err = collectOpsInputs(root, &in); err != nil {
		return inputs{}, err
	}
	if err = collectMediaAndRepoInputs(root, &in); err != nil {
		return inputs{}, err
	}
	return in, nil
}
