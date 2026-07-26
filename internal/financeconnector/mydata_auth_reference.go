package financeconnector

import (
	"context"
	"fmt"
)

// SignVerificationFromReference resolves one atomic bundle and performs the
// private authentication-plane call without exposing the bundle to callers.
func (client LiveClient) SignVerificationFromReference(
	ctx context.Context, reference string, request SignVerificationRequest,
) (string, error) {
	if client.Resolve == nil {
		client.Resolve = ResolveOnePasswordCLI
	}
	credential, err := client.Resolve(ctx, reference, client.Config.Provider)
	if err != nil {
		return "", err
	}
	if credential.Provider != client.Config.Provider {
		return "", fmt.Errorf("credential bundle provider mismatch")
	}
	credential.AuthAccessToken, err = client.authAccessToken(ctx, credential)
	if err != nil {
		return "", err
	}
	return client.SignVerification(ctx, credential, request)
}
