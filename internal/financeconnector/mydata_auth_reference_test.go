package financeconnector

import (
	"context"
	"strings"
	"testing"
)

func TestSignVerificationFromReferenceKeepsProviderAtomic(t *testing.T) {
	client := LiveClient{
		Config: RuntimeConfig{Provider: ProviderTossMyData, AuthBaseURL: "https://example.invalid"},
		Resolve: func(_ context.Context, reference, provider string) (credentialEnvelope, error) {
			if reference != "op://vault/item/bundle" || provider != ProviderTossMyData {
				t.Fatalf("reference = %q provider = %q", reference, provider)
			}
			return credentialEnvelope{Provider: "other", AuthAccessToken: "token"}, nil
		},
	}
	_, err := client.SignVerificationFromReference(context.Background(), "op://vault/item/bundle", SignVerificationRequest{})
	if err == nil || !strings.Contains(err.Error(), "provider mismatch") {
		t.Fatalf("err = %v", err)
	}
}
