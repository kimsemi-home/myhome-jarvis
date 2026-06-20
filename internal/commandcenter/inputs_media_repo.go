package commandcenter

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/mediareadiness"
	"github.com/kimsemi-home/myhome-jarvis/internal/mergeevidence"
	"github.com/kimsemi-home/myhome-jarvis/internal/repofactory"
)

type mediaReadinessStatus = mediareadiness.Status
type mergeEvidenceStatus = mergeevidence.Status
type repoFactoryStatus = repofactory.Status
type repoFactoryPreflightStatus = repofactory.DecisionPacket

func collectMediaAndRepoInputs(root string, in *inputs) error {
	var err error
	if in.MediaReadiness, err = mediareadiness.StatusForRoot(root); err != nil {
		return err
	}
	if in.MergeEvidence, err = mergeevidence.StatusForRoot(root); err != nil {
		return err
	}
	if in.RepoFactory, err = repofactory.StatusForRoot(root); err != nil {
		return err
	}
	if in.RepoFactoryPreflight, err = repofactory.DecisionPacketForRoot(root); err != nil {
		return err
	}
	return nil
}
