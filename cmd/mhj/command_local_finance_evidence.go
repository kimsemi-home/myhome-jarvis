package main

import (
	"path/filepath"

	"github.com/kimsemi-home/myhome-jarvis/internal/localfinanceevidence"
	"github.com/kimsemi-home/myhome-jarvis/internal/localfinancereadiness"
)

func routeLocalFinance(root string, args []string) error {
	if len(args) == 1 && args[0] == "evidence" {
		return localFinanceEvidence(root, "fixtures/local_finance/manifest.json")
	}
	if len(args) == 2 && args[0] == "evidence" {
		return localFinanceEvidence(root, args[1])
	}
	if len(args) == 1 && args[0] == "readiness" {
		return localFinanceReadiness(root, "fixtures/local_finance_readiness/manifest.json")
	}
	if len(args) == 2 && args[0] == "readiness" {
		return localFinanceReadiness(root, args[1])
	}
	return usage()
}

func localFinanceReadiness(root, path string) error {
	if !filepath.IsAbs(path) {
		path = filepath.Join(root, filepath.FromSlash(path))
	}
	manifest, err := localfinancereadiness.Read(path)
	if err != nil {
		return err
	}
	return writeJSON(manifest)
}

func localFinanceEvidence(root, path string) error {
	if !filepath.IsAbs(path) {
		path = filepath.Join(root, filepath.FromSlash(path))
	}
	manifest, err := localfinanceevidence.Read(path)
	if err != nil {
		return err
	}
	return writeJSON(manifest)
}
