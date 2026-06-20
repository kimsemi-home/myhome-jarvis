package main

import (
	"errors"

	"github.com/kimsemi-home/myhome-jarvis/internal/security"
)

func runSecurity(root string, args []string) error {
	if len(args) != 1 {
		return usage()
	}
	switch args[0] {
	case "status":
		return securityStatus(root)
	case "check":
		report, err := security.Check(root)
		if err != nil {
			return err
		}
		if err := writeJSON(report); err != nil {
			return err
		}
		if !report.OK {
			return errors.New("security check failed")
		}
		return nil
	case "history":
		report, err := security.CheckHistory(root)
		if err != nil {
			return err
		}
		if err := writeJSON(report); err != nil {
			return err
		}
		if !report.OK {
			return errors.New("security history check failed")
		}
		return nil
	default:
		return usage()
	}
}
