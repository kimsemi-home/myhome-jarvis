package financeconnector

import (
	"errors"
	"strings"
)

var errFinancePreflightFailed = errors.New("finance preflight failed")

func validatePreflightCredential(config RuntimeConfig, credential credentialEnvelope) error {
	if strings.TrimSpace(credential.AuthAccessToken) == "" &&
		(strings.TrimSpace(credential.ClientID) == "" || strings.TrimSpace(credential.ClientSecret) == "") {
		return errors.New("auth material missing")
	}
	_, err := (LiveClient{Config: config}).httpClientForCredential(credential)
	return err
}

func allPreflightBundlesReady(connections []PreflightConnection) bool {
	for _, connection := range connections {
		if !connection.BundleReady {
			return false
		}
	}
	return true
}
