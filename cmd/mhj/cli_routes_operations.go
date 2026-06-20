package main

func routeOperations(root string, args []string) (bool, error) {
	switch args[0] {
	case "ci":
		return true, requireSubcommand(args, "verify", func() error { return runCIVerify(root) })
	case "verification":
		return true, routeVerification(root, args[1:])
	case "control-plane":
		return true, requireSubcommand(args, "verify", func() error { return controlPlaneVerify(root) })
	case "context-pack":
		return true, routeContextPack(root, args[1:])
	case "codex-sustainability":
		return true, routeCodexSustainability(root, args[1:])
	case "toolchain":
		return true, requireSubcommand(args, "verify", func() error { return runToolchainVerify(root) })
	case "ddd":
		return true, requireSubcommand(args, "verify", func() error { return runDDDVerify(root) })
	case "loop":
		return true, routeLoop(root, args[1:])
	case "benchmark":
		return true, requireSubcommand(args, "smoke", func() error { return runBenchmarkSmoke(root) })
	case "quality":
		return true, runQuality(root)
	case "codegen":
		return true, routeCodegen(root, args[1:])
	default:
		return false, nil
	}
}

func routeContextPack(root string, args []string) error {
	if len(args) == 1 && args[0] == "verify" {
		return contextPackVerify(root, "")
	}
	if len(args) == 2 && args[0] == "verify" {
		return contextPackVerify(root, args[1])
	}
	return usage()
}

func routeLoop(root string, args []string) error {
	if len(args) == 1 && args[0] == "once" {
		return loopOnce(root)
	}
	if len(args) == 1 && args[0] == "status" {
		return loopStatus(root)
	}
	if len(args) >= 1 && args[0] == "worker" {
		return loopWorker(root, args[1:])
	}
	return usage()
}

func routeVerification(root string, args []string) error {
	if len(args) == 1 && args[0] == "verify" {
		return runVerificationVerify(root)
	}
	if len(args) == 1 && args[0] == "evidence" {
		return runVerificationEvidence(root)
	}
	return usage()
}

func routeCodegen(root string, args []string) error {
	if len(args) == 1 && args[0] == "verify" {
		return runCodegenVerify(root)
	}
	return runCodegen(root)
}
