package main

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/confidence"
	"github.com/kimsemi-home/myhome-jarvis/internal/controlplane"
	"github.com/kimsemi-home/myhome-jarvis/internal/evidence"
	"github.com/kimsemi-home/myhome-jarvis/internal/translation"
)

func evidenceStatus(root string) error {
	status, err := evidence.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func confidenceStatus(root string) error {
	status, err := confidence.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func translationStatus(root string) error {
	status, err := translation.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func controlPlaneStatus(root string) error {
	status, err := controlplane.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}
