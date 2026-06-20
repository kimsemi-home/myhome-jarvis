package commandcenter

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/codexcost"
	"github.com/kimsemi-home/myhome-jarvis/internal/codexsustainability"
	"github.com/kimsemi-home/myhome-jarvis/internal/contextpack"
	"github.com/kimsemi-home/myhome-jarvis/internal/financeconsent"
	"github.com/kimsemi-home/myhome-jarvis/internal/monetization"
	"github.com/kimsemi-home/myhome-jarvis/internal/review"
	"github.com/kimsemi-home/myhome-jarvis/internal/storagearchive"
)

func collectOpsInputs(root string, in *inputs) error {
	var err error
	if in.Review, err = review.StatusForRoot(root); err != nil {
		return err
	}
	if in.FinanceConsent, err = financeconsent.StatusForRoot(root); err != nil {
		return err
	}
	if in.Cost, err = codexcost.StatusForRoot(root); err != nil {
		return err
	}
	if in.CostBrief, err = codexcost.BriefForRoot(root); err != nil {
		return err
	}
	if in.CodexSustainability, err = codexsustainability.StatusForRoot(root); err != nil {
		return err
	}
	if in.StorageArchive, err = storagearchive.StatusForRoot(root); err != nil {
		return err
	}
	if in.ContextPack, err = contextpack.StatusForRoot(root); err != nil {
		return err
	}
	if in.Monetization, err = monetization.StatusForRoot(root); err != nil {
		return err
	}
	return nil
}
