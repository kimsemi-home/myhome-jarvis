package main

import (
	"flag"
	"os"

	"github.com/kimsemi-home/myhome-jarvis/internal/externalbootstrap"
	"github.com/kimsemi-home/myhome-jarvis/internal/externalevidence"
)

func routeExternalEvidence(root string, args []string) error {
	if len(args) == 1 && args[0] == "status" {
		return externalEvidenceStatus(root)
	}
	if len(args) == 1 && args[0] == "repo-split-decision" {
		return externalEvidenceRepoSplitDecision(root)
	}
	if len(args) == 1 && args[0] == "repo-bootstrap" {
		return externalEvidenceRepoBootstrap(root)
	}
	if len(args) >= 1 && args[0] == "child-repo-status" {
		return externalEvidenceChildRepoStatus(root, args[1:])
	}
	if len(args) >= 1 && args[0] == "collect" {
		return externalEvidenceCollect(root, args[1:])
	}
	return usage()
}

func externalEvidenceStatus(root string) error {
	status, err := externalevidence.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func externalEvidenceRepoSplitDecision(root string) error {
	packet, err := externalevidence.RepoSplitDecisionPacketForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(packet)
}

func externalEvidenceRepoBootstrap(root string) error {
	packet, err := externalbootstrap.PacketForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(packet)
}

func externalEvidenceCollect(root string, args []string) error {
	flags := flag.NewFlagSet("external-evidence collect", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	maxSources := flags.Int("max-sources", 0, "maximum configured sources to collect")
	if err := flags.Parse(args); err != nil {
		return err
	}
	report, err := externalevidence.CollectForRoot(root, *maxSources)
	if err != nil {
		return err
	}
	return writeJSON(report)
}
