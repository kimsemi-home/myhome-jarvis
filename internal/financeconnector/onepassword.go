package financeconnector

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

type CredentialResolver func(context.Context, string, string) (credentialEnvelope, error)

func ResolveOnePasswordCLI(ctx context.Context, reference string, provider string) (credentialEnvelope, error) {
	if !strings.HasPrefix(strings.TrimSpace(reference), "op://") {
		return credentialEnvelope{}, fmt.Errorf("credential reference must use op://")
	}
	output, err := exec.CommandContext(
		ctx, "op", "read", "--no-newline", reference,
	).Output()
	if err != nil {
		return credentialEnvelope{}, fmt.Errorf("1Password read failed")
	}
	return parseCredentialEnvelope(string(output), provider)
}
