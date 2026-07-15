package main

import (
	"errors"
	"fmt"

	"github.com/kimsemi-home/myhome-jarvis/internal/repogovernance"
)

func repogovernanceChecks(root string) error {
	if err := repogovernance.CheckDocuments(root); err != nil {
		return err
	}
	_, err := repogovernance.LoadManifest(root)
	return err
}

func traceCommand(root, action, base string) error {
	manifest, err := repogovernance.LoadManifest(root)
	if err != nil {
		return err
	}
	switch action {
	case "verify":
		fmt.Println("Shorts factory traceability verified")
		return nil
	case "staged":
		files, err := repogovernance.StagedFiles(root)
		return checkTrace(manifest, files, err)
	case "range":
		files, err := repogovernance.RangeFiles(root, base)
		return checkTrace(manifest, files, err)
	case "ci":
		if ciBase() == "" {
			fmt.Println("document-code range gate skipped outside pull request")
			return nil
		}
		files, err := repogovernance.RangeFiles(root, ciBase())
		return checkTrace(manifest, files, err)
	default:
		return errors.New("unknown trace action")
	}
}

func checkTrace(manifest repogovernance.Manifest, files []string, err error) error {
	if err != nil {
		return err
	}
	if err := repogovernance.CheckChanges(manifest, files); err != nil {
		return err
	}
	fmt.Println("Shorts factory document-code changes verified")
	return nil
}
