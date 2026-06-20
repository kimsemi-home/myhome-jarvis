package main

func routeCodegen(root string, args []string) error {
	if len(args) == 1 && args[0] == "verify" {
		return runCodegenVerify(root)
	}
	return runCodegen(root)
}
