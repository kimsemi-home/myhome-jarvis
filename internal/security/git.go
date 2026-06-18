package security

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

func gitLines(root string, args ...string) ([]string, error) {
	output, err := gitOutput(root, args...)
	if err != nil {
		return nil, err
	}
	return splitOutputLines(output), nil
}

func gitLinesAllowNoMatches(root string, args ...string) ([]string, error) {
	output, err := gitOutput(root, args...)
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) && exitErr.ExitCode() == 1 {
			return nil, nil
		}
		return nil, err
	}
	return splitOutputLines(output), nil
}

func gitOutput(root string, args ...string) (string, error) {
	allArgs := append([]string{"-C", root}, args...)
	cmd := exec.Command("git", allArgs...)
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	if err := cmd.Run(); err != nil {
		message := strings.TrimSpace(output.String())
		if message == "" {
			message = "git command failed"
		}
		return "", gitCommandError{err: err, message: message}
	}
	return output.String(), nil
}
