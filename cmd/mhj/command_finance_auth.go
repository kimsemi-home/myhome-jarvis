package main

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/kimsemi-home/myhome-jarvis/internal/financeconnector"
)

type financeAuthReport struct {
	Provider string `json:"provider"`
	Owner    string `json:"owner"`
	Verified bool   `json:"verified"`
}

func runFinanceAuth(root string, owner string) error {
	config, err := financeconnector.RuntimeConfigFromEnv(os.Getenv)
	if err != nil {
		return err
	}
	if config.Mode != financeconnector.ModeMyData {
		return errors.New("finance auth requires MYHOME_FINANCE_MODE=mydata")
	}
	if err := financeconnector.RequireActiveConsent(root); err != nil {
		return err
	}
	connection, err := financeConnectionForOwner(config.Connections, owner)
	if err != nil {
		return err
	}
	var request financeconnector.SignVerificationRequest
	if err := json.NewDecoder(os.Stdin).Decode(&request); err != nil {
		return errors.New("finance auth requires a private signed-consent JSON document on stdin")
	}
	client := financeconnector.LiveClient{Config: config}
	if _, err := client.SignVerificationFromReference(
		context.Background(), connection.Credential.OnePasswordRef, request,
	); err != nil {
		return err
	}
	return writeJSON(financeAuthReport{Provider: config.Provider, Owner: owner, Verified: true})
}

func financeConnectionForOwner(
	connections []financeconnector.ConnectionConfig, owner string,
) (financeconnector.ConnectionConfig, error) {
	owner = strings.TrimSpace(owner)
	if owner == "" {
		return financeconnector.ConnectionConfig{}, errors.New("finance auth owner is required")
	}
	var match financeconnector.ConnectionConfig
	for _, connection := range connections {
		if connection.Owner != owner {
			continue
		}
		if match.Owner != "" {
			return financeconnector.ConnectionConfig{}, errors.New("finance auth owner is ambiguous")
		}
		match = connection
	}
	if match.Owner == "" {
		return financeconnector.ConnectionConfig{}, errors.New("finance auth owner is not configured")
	}
	return match, nil
}
