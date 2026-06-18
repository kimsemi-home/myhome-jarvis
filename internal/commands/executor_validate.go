package commands

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

func validateInvocation(invocation Invocation) error {
	if len(invocation.Argv) == 0 {
		return errors.New("execution requires an argv plan")
	}
	executable := filepath.Base(invocation.Argv[0])
	switch executable {
	case "open", "osascript", "pmset":
	default:
		return fmt.Errorf("executable %q is not allowed", executable)
	}
	for _, arg := range invocation.Argv {
		if strings.ContainsRune(arg, 0) {
			return errors.New("argv contains NUL byte")
		}
	}
	return nil
}
