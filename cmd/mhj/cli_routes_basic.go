package main

import (
	"fmt"

	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
)

func routeBasics(root string, args []string) (bool, error) {
	switch args[0] {
	case "version":
		fmt.Println("myhome-jarvis " + version)
		return true, nil
	case "commands":
		return true, writeJSON(commands.Specs())
	case "auth":
		return true, runAuth(root, args[1:])
	case "assistant":
		return true, runAssistant(root, args[1:])
	case "security":
		return true, runSecurity(root, args[1:])
	case "command":
		return true, runCommand(root, args[1:])
	case "learning":
		return true, runLearning(root, args[1:])
	case "harness":
		return true, runHarness(root, args[1:])
	case "linear":
		return true, runLinear(root, args[1:])
	case "repo-factory":
		return true, routeRepoFactory(root, args[1:])
	case "daemon":
		return true, runDaemon(root, args[1:])
	case "knowledge":
		return true, runKnowledge(root, args[1:])
	default:
		return false, nil
	}
}
