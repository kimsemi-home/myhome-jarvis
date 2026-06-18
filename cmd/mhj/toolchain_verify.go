package main

import "fmt"

func runToolchainVerify(root string) error {
	if err := validateToolchainPins(root); err != nil {
		return err
	}
	return writeJSON(map[string]any{"ok": true})
}

func validateToolchainPins(root string) error {
	goVersion, err := readTrimmedFile(root, ".go-version")
	if err != nil {
		return err
	}
	goModVersion, err := parseFirstMatchFile(root, "go.mod", `(?m)^go\s+([0-9]+\.[0-9]+\.[0-9]+)\s*$`)
	if err := requireEqual("go.mod go directive", goVersion, goModVersion, err); err != nil {
		return err
	}
	generatedGo, err := generatedProjectGoVersion(root)
	if err != nil {
		return err
	}
	if err := requireEqualValue("generated project go_version", goVersion, generatedGo); err != nil {
		return err
	}
	workflowGoVersion, err := parseWorkflowEnv(root, "GO_VERSION")
	if err := requireEqual("workflow GO_VERSION", goVersion, workflowGoVersion, err); err != nil {
		return err
	}
	rustVersion, err := parseFirstMatchFile(root, "rust-toolchain.toml", `(?m)^channel\s*=\s*"([^"]+)"\s*$`)
	if err != nil {
		return err
	}
	workflowRustVersion, err := parseWorkflowEnv(root, "RUST_TOOLCHAIN")
	return requireEqual("workflow RUST_TOOLCHAIN", rustVersion, workflowRustVersion, err)
}

func requireEqual(label string, expected string, actual string, err error) error {
	if err != nil {
		return err
	}
	return requireEqualValue(label, expected, actual)
}

func requireEqualValue(label string, expected string, actual string) error {
	if actual != expected {
		return fmt.Errorf("%s = %q, expected %q", label, actual, expected)
	}
	return nil
}
