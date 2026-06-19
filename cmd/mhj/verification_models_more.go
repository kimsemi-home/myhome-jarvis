package main

type verificationReleaseFile struct {
	Gates []verificationGate `json:"gates"`
}

type verificationGate struct {
	ID       string `json:"id"`
	Kind     string `json:"kind"`
	Required bool   `json:"required"`
}

type verificationTestsFile struct {
	Tests []verificationTest `json:"tests"`
}

type verificationTest struct {
	ID       string `json:"id"`
	Kind     string `json:"kind"`
	Evidence string `json:"evidence"`
}
