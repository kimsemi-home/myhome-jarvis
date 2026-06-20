package main

func routeRepoFactory(root string, args []string) error {
	if len(args) != 1 {
		return usage()
	}
	switch args[0] {
	case "status":
		return repoFactoryStatus(root)
	case "decision-packet":
		return repoFactoryDecisionPacket(root)
	default:
		return usage()
	}
}
