package main

import (
	"errors"
	"os"
)

func run(args []string) error {
	root, err := repositoryRoot()
	if err != nil {
		return err
	}
	if len(args) == 1 && args[0] == "verify" {
		return verify(root)
	}
	if len(args) == 2 && args[0] == "docs" {
		return docs(root, args[1])
	}
	if len(args) == 2 && args[0] == "trace" {
		return traceCommand(root, args[1], "")
	}
	if len(args) == 3 && args[0] == "trace" && args[1] == "range" {
		return traceCommand(root, "range", args[2])
	}
	if len(args) == 3 && args[0] == "gate" && args[1] == "evaluate" {
		return evaluate(root, args[2])
	}
	return errors.New("usage: shortsctl <verify|docs generate|docs check|gate evaluate PATH|trace verify|trace staged|trace range BASE|trace ci>")
}

func ciBase() string {
	if os.Getenv("GITHUB_EVENT_NAME") != "pull_request" {
		return ""
	}
	return "origin/" + os.Getenv("GITHUB_BASE_REF")
}
