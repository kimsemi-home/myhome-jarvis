package commands

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

type generatedCommandCatalog struct {
	Commands []generatedCommand `json:"commands"`
}

type generatedCommand struct {
	Name            string   `json:"name"`
	Summary         string   `json:"summary"`
	PayloadFields   []string `json:"payload_fields"`
	DryRunDefault   bool     `json:"dry_run_default"`
	ShortcutFor     string   `json:"shortcut_for"`
	Target          string   `json:"target"`
	AllowedServices []string `json:"allowed_services"`
}

func loadGeneratedCommandCatalog(t *testing.T) generatedCommandCatalog {
	t.Helper()
	body, err := os.ReadFile(filepath.Join(repoRoot(t), "generated", "commands.generated.json"))
	if err != nil {
		t.Fatal(err)
	}
	var catalog generatedCommandCatalog
	if err := json.Unmarshal(body, &catalog); err != nil {
		t.Fatal(err)
	}
	if len(catalog.Commands) == 0 {
		t.Fatal("generated command catalog is empty")
	}
	return catalog
}

func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		next := filepath.Dir(dir)
		if next == dir {
			t.Fatal("could not find repo root")
		}
		dir = next
	}
}
