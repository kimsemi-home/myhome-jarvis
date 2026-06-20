package main

func routeStorageArchive(root string, args []string) error {
	if len(args) != 1 {
		return usage()
	}
	switch args[0] {
	case "run":
		return storageArchiveRun(root)
	case "status":
		return storageArchiveStatus(root)
	default:
		return usage()
	}
}
