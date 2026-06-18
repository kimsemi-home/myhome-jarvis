package main

import "os"

func run(args []string) error {
	if len(args) == 0 {
		return usage()
	}
	root, err := os.Getwd()
	if err != nil {
		return err
	}
	return routeCommand(root, args)
}
