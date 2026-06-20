package main

func requireSubcommand(args []string, name string, run func() error) error {
	if len(args) == 2 && args[1] == name {
		return run()
	}
	return usage()
}
