package commandcenter

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/mediareadiness"
	"github.com/kimsemi-home/myhome-jarvis/internal/repofactory"
)

type mediaReadinessStatus = mediareadiness.Status
type repoFactoryStatus = repofactory.Status

func collectMediaAndRepoInputs(root string, in *inputs) error {
	var err error
	if in.MediaReadiness, err = mediareadiness.StatusForRoot(root); err != nil {
		return err
	}
	if in.RepoFactory, err = repofactory.StatusForRoot(root); err != nil {
		return err
	}
	return nil
}
