package commandcenter

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/authority"
	"github.com/kimsemi-home/myhome-jarvis/internal/codexcost"
	"github.com/kimsemi-home/myhome-jarvis/internal/codexsustainability"
	"github.com/kimsemi-home/myhome-jarvis/internal/contextpack"
	"github.com/kimsemi-home/myhome-jarvis/internal/financeconsent"
	"github.com/kimsemi-home/myhome-jarvis/internal/incidents"
	"github.com/kimsemi-home/myhome-jarvis/internal/monetization"
	"github.com/kimsemi-home/myhome-jarvis/internal/pdca"
	"github.com/kimsemi-home/myhome-jarvis/internal/review"
)

type inputs struct {
	Vision              visionPolicy
	PDCA                pdca.Status
	Evidence            evidenceStatus
	EvidenceIntegrity   evidenceIntegrityStatus
	Incidents           incidents.Status
	Authority           authority.Status
	Review              review.Status
	FinanceConsent      financeconsent.Status
	Cost                codexcost.Status
	CodexSustainability codexsustainability.Status
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
	if in.Authority, err = authority.StatusForRoot(root); err != nil {
		return inputs{}, err
	}
	if in.Review, err = review.StatusForRoot(root); err != nil {
		return inputs{}, err
	}
	if in.FinanceConsent, err = financeconsent.StatusForRoot(root); err != nil {
		return inputs{}, err
	}
	if in.Cost, err = codexcost.StatusForRoot(root); err != nil {
		return inputs{}, err
	}
	if in.CodexSustainability, err = codexsustainability.StatusForRoot(root); err != nil {
		return inputs{}, err
	}
	if in.ContextPack, err = contextpack.StatusForRoot(root); err != nil {
		return inputs{}, err
	}
	if in.Monetization, err = monetization.StatusForRoot(root); err != nil {
		return inputs{}, err
	}
	if err = collectMediaAndRepoInputs(root, &in); err != nil {
		return inputs{}, err
	}
	return in, nil
}
