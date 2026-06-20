package main

import (
	"errors"

	"github.com/kimsemi-home/myhome-jarvis/internal/contextpack"
)

func contextPackStatus(root string) error {
	status, err := contextpack.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func contextPackVerify(root string, declarationPath string) error {
	result, err := contextpack.VerifyDeclarationForRoot(root, declarationPath)
	if err != nil {
		return err
	}
	if err := writeJSON(result); err != nil {
		return err
	}
	if !result.Valid {
		return errors.New("context pack declaration drift detected")
	}
	return nil
}
