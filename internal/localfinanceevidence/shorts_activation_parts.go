package localfinanceevidence

import "errors"

func validateShortsActivationCallback(value ShortsActivationCallback) error {
	if value.Network != "tcp4" || value.Host != "127.0.0.1" || value.PortStrategy != "random_available" ||
		value.Path != "/oauth/callback" || value.InvalidAttemptsRejected != 6 || !value.ValidCallbackConsumed ||
		!value.DuplicateCallbackDenied || !value.CanceledReceiverClosed || value.RawStateReported ||
		value.RawCodeReported || value.ExternalNetworkRequested {
		return errors.New("Shorts activation callback proof is invalid")
	}
	return nil
}

func validateShortsActivationKeychain(value ShortsActivationKeychain) error {
	if value.Executable != "/usr/bin/security" || !value.FakeRunner || value.ActualKeychainExecuted ||
		value.CommandsValidated != 3 || !value.InteractiveProvisioning || value.SecretValueInArguments ||
		!value.DefaultPermitDenied || !value.ExpiredPermitDenied || !value.UnlistedReferenceDenied ||
		!value.ExternalNetworkPermitDenied || !value.ReadinessPlanBound || value.RawCredentialReported ||
		!value.ReturnedMaterialZeroed {
		return errors.New("Shorts activation Keychain proof is invalid")
	}
	return nil
}

func slicesEqual(left, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for index := range left {
		if left[index] != right[index] {
			return false
		}
	}
	return true
}
