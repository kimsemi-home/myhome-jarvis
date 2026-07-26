package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/kimsemi-home/myhome-jarvis/internal/shortsfactory"
)

func verify(root string) error {
	if _, err := shortsfactory.Verify(root); err != nil {
		return err
	}
	if err := repogovernanceChecks(root); err != nil {
		return err
	}
	fmt.Println("public-safe Shorts factory contract verified")
	return nil
}

func evaluate(root, path string) error {
	contract, err := shortsfactory.LoadContract(filepath.Join(root, "contracts", "shorts-factory.json"))
	if err != nil {
		return err
	}
	if !filepath.IsAbs(path) {
		path = filepath.Join(root, filepath.FromSlash(path))
	}
	request, err := shortsfactory.LoadRequest(path)
	if err != nil {
		return err
	}
	result, err := shortsfactory.Evaluate(request, contract)
	if err != nil {
		return err
	}
	body, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(body))
	if result.Decision != "approved" {
		return errors.New("evidence gate rejected the request")
	}
	return nil
}
