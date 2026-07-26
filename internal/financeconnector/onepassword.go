package financeconnector

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

type SecretResolver func(context.Context, string) (string, error)

func ResolveOnePasswordCLI(ctx context.Context, reference string) (string, error) {
	if !strings.HasPrefix(strings.TrimSpace(reference), "op://") {
		return "", fmt.Errorf("secret reference must use op://")
	}
	output, err := exec.CommandContext(
		ctx, "op", "read", "--no-newline", reference,
	).Output()
	if err != nil {
		return "", fmt.Errorf("1Password read failed")
	}
	value := strings.TrimSpace(string(output))
	if value == "" {
		return "", fmt.Errorf("1Password reference resolved to an empty value")
	}
	return value, nil
}
