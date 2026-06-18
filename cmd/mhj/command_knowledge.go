package main

import (
	"errors"
	"strings"

	"github.com/kimsemi-home/myhome-jarvis/internal/knowledge"
	"github.com/kimsemi-home/myhome-jarvis/internal/learning"
)

func runDDDVerify(root string) error {
	report, err := knowledge.Verify(root)
	if err != nil {
		return err
	}
	if err := writeJSON(report); err != nil {
		return err
	}
	if !report.OK {
		return errors.New("ddd verify failed")
	}
	return nil
}

func runKnowledge(root string, args []string) error {
	if len(args) == 1 && args[0] == "verify" {
		return runDDDVerify(root)
	}
	if len(args) >= 2 && args[0] == "search" {
		report, err := knowledge.Search(root, strings.Join(args[1:], " "))
		if err != nil {
			return err
		}
		return writeJSON(report)
	}
	return errors.New("usage: mhj knowledge <verify|search query>")
}

func runLearning(root string, args []string) error {
	if len(args) == 1 && args[0] == "status" {
		status, err := learning.StatusForRoot(root)
		if err != nil {
			return err
		}
		return writeJSON(status)
	}
	if len(args) == 2 && args[0] == "record" {
		result, err := learning.Record(root, []byte(args[1]))
		if err != nil {
			return err
		}
		return writeJSON(result)
	}
	return errors.New("usage: mhj learning <status|record json-payload>")
}
