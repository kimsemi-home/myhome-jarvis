package main

import (
	"errors"
	"fmt"

	"github.com/kimsemi-home/myhome-jarvis/internal/repogovernance"
)

func docs(root, action string) error {
	switch action {
	case "generate":
		if err := repogovernance.GenerateDocuments(root); err != nil {
			return err
		}
		fmt.Println("Shorts factory documents generated")
	case "check":
		if err := repogovernance.CheckDocuments(root); err != nil {
			return err
		}
		fmt.Println("Shorts factory documents verified")
	default:
		return errors.New("usage: shortsctl docs <generate|check>")
	}
	return nil
}
