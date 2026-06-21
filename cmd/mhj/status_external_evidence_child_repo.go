package main

import (
	"errors"

	"github.com/kimsemi-home/myhome-jarvis/internal/externalbootstrap"
)

func externalEvidenceChildRepoStatus(root string, args []string) error {
	if len(args) > 1 {
		return usage()
	}
	childRoot := ""
	if len(args) == 1 {
		childRoot = args[0]
	}
	status, err := externalbootstrap.ChildRepoStatusForRoot(root, childRoot)
	if err != nil {
		return err
	}
	if err := writeJSON(status); err != nil {
		return err
	}
	if !status.Valid {
		return errors.New("external evidence child repo is not verified")
	}
	return nil
}
