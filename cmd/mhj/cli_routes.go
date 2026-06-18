package main

type cliRoute func(string, []string) (bool, error)

func routeCommand(root string, args []string) error {
	for _, route := range []cliRoute{routeBasics, routeStatuses, routeOperations} {
		if ok, err := route(root, args); ok {
			return err
		}
	}
	return usage()
}
