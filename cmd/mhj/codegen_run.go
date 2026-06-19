package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func runCodegen(root string) error {
	if _, err := exec.LookPath("sbcl"); err != nil {
		return errors.New("missing executable: sbcl")
	}
	cmd := exec.Command("sbcl", "--script", "lisp/scripts/codegen.lisp")
	cmd.Dir = root
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func runCodegenVerify(root string) error {
	before, err := generatedSnapshot(root)
	if err != nil {
		return err
	}
	if err := runCodegen(root); err != nil {
		return err
	}
	after, err := generatedSnapshot(root)
	if err != nil {
		return err
	}
	changed := changedGeneratedFiles(before, after)
	if len(changed) > 0 {
		return fmt.Errorf("generated artifacts are out of date: %s", strings.Join(changed, ", "))
	}
	if _, err := validateVerificationGenerated(root); err != nil {
		return err
	}
	fmt.Fprintln(os.Stdout, "Generated artifacts verified")
	return nil
}
