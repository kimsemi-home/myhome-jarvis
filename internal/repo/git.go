package repo

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func git(root string, args ...string) (string, error) {
	output, err := gitRaw(root, args...)
	return strings.TrimSpace(output), err
}

func gitRaw(root string, args ...string) (string, error) {
	command := append([]string{"-C", root}, args...)
	cmd := exec.Command("git", command...)
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("git %s: %s", strings.Join(args, " "), strings.TrimSpace(output.String()))
	}
	return output.String(), nil
}
