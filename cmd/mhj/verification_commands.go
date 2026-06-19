package main

import "fmt"

func verifyGraphCommands(graph verificationGraphFile) error {
	commands := map[string]bool{}
	for _, unit := range graph.Units {
		for _, command := range unit.Commands {
			commands[command] = true
		}
	}
	for _, command := range requiredVerificationCommands() {
		if !commands[command] {
			return fmt.Errorf("verification graph missing command %q", command)
		}
	}
	return nil
}
