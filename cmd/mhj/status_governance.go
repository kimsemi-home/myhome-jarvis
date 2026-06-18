package main

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/authority"
	"github.com/kimsemi-home/myhome-jarvis/internal/evidencequality"
	"github.com/kimsemi-home/myhome-jarvis/internal/incidents"
	"github.com/kimsemi-home/myhome-jarvis/internal/review"
)

func incidentsStatus(root string) error {
	status, err := incidents.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func evidenceQualityStatus(root string) error {
	status, err := evidencequality.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func reviewStatus(root string) error {
	status, err := review.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func authorityStatus(root string) error {
	status, err := authority.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}
