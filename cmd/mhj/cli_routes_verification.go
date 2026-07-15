package main

func routeVerification(root string, args []string) error {
	if len(args) == 1 && args[0] == "verify" {
		return runVerificationVerify(root)
	}
	if len(args) == 1 && args[0] == "evidence" {
		return runVerificationEvidence(root)
	}
	return usage()
}
