package commands

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"sort"
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

func TestSpecsMatchGeneratedCommandCatalog(t *testing.T) {
	catalog := loadGeneratedCommandCatalog(t)
	generatedByName := make(map[string]generatedCommand, len(catalog.Commands))
	for _, command := range catalog.Commands {
		if command.Name == "" {
			t.Fatal("generated command name is empty")
		}
		if !command.DryRunDefault {
			t.Fatalf("generated command %q is not dry-run by default", command.Name)
		}
		if _, exists := generatedByName[command.Name]; exists {
			t.Fatalf("duplicate generated command %q", command.Name)
		}
		generatedByName[command.Name] = command
	}

	specs := Specs()
	if len(specs) != len(generatedByName) {
		t.Fatalf("spec count = %d, generated command count = %d", len(specs), len(generatedByName))
	}
	for _, spec := range specs {
		command, ok := generatedByName[spec.Name]
		if !ok {
			t.Fatalf("Go command spec %q missing from generated command catalog", spec.Name)
		}
		if spec.Summary != command.Summary {
			t.Fatalf("summary mismatch for %q: Go %q generated %q", spec.Name, spec.Summary, command.Summary)
		}
		if !reflect.DeepEqual(spec.PayloadFields, command.PayloadFields) {
			t.Fatalf("payload fields mismatch for %q: Go %#v generated %#v", spec.Name, spec.PayloadFields, command.PayloadFields)
		}
		delete(generatedByName, spec.Name)
	}
	if len(generatedByName) != 0 {
		t.Fatalf("generated commands missing from Go registry: %#v", generatedByName)
	}
}

func TestGeneratedCommandTargetsMatchRegistry(t *testing.T) {
	catalog := loadGeneratedCommandCatalog(t)
	for _, command := range catalog.Commands {
		if command.Target == "" {
			continue
		}
		plan, err := Build(command.Name, []byte(`{}`))
		if err != nil {
			t.Fatalf("building generated target command %q: %v", command.Name, err)
		}
		if len(plan.Invocations) != 1 || plan.Invocations[0].URL != command.Target {
			t.Fatalf("target mismatch for %q: plan %#v generated target %q", command.Name, plan.Invocations, command.Target)
		}
	}
}

func TestGeneratedOTTServicesMatchRegistry(t *testing.T) {
	catalog := loadGeneratedCommandCatalog(t)
	var generated []string
	for _, command := range catalog.Commands {
		if command.Name == "open_ott" {
			generated = append(generated, command.AllowedServices...)
			break
		}
	}
	if len(generated) == 0 {
		t.Fatal("generated open_ott command does not list allowed services")
	}
	sort.Strings(generated)

	registry := make([]string, 0, len(ottURLs()))
	for service := range ottURLs() {
		registry = append(registry, service)
	}
	sort.Strings(registry)
	if !reflect.DeepEqual(registry, generated) {
		t.Fatalf("OTT service mismatch: registry %#v generated %#v", registry, generated)
	}
}

func TestBuildVolumeSetRejectsOutOfRange(t *testing.T) {
	if _, err := Build("volume-set", []byte(`{"level":101}`)); err == nil {
		t.Fatal("expected out-of-range level to fail")
	}
}

func TestBuildOpenURLRejectsUnsafeScheme(t *testing.T) {
	if _, err := Build("open-url", []byte(`{"url":"javascript:alert(1)"}`)); err == nil {
		t.Fatal("expected unsafe URL to fail")
	}
}

func TestBuildOTTShortcuts(t *testing.T) {
	cases := map[string]string{
		"open-netflix":      "https://www.netflix.com",
		"open-disney-plus":  "https://www.disneyplus.com",
		"open-tving":        "https://www.tving.com",
		"open-wavve":        "https://www.wavve.com",
		"open-coupang-play": "https://www.coupangplay.com",
	}
	for command, expectedURL := range cases {
		t.Run(command, func(t *testing.T) {
			plan, err := Build(command, []byte(`{}`))
			if err != nil {
				t.Fatal(err)
			}
			if len(plan.Invocations) != 1 || plan.Invocations[0].URL != expectedURL {
				t.Fatalf("unexpected plan: %#v", plan)
			}
		})
	}
}

func TestRunHomeHarness(t *testing.T) {
	report := RunHomeHarness()
	if !report.Passed {
		t.Fatalf("home harness failed: %+v", report.Results)
	}
}

func TestRunFinanceHarness(t *testing.T) {
	report := RunFinanceHarness(repoRoot(t))
	if !report.Passed {
		t.Fatalf("finance harness failed: %+v", report.Results)
	}
	if report.Name != "finance" {
		t.Fatalf("harness name = %q", report.Name)
	}
	if len(report.Results) < 8 {
		t.Fatalf("finance harness result count = %d", len(report.Results))
	}
}

func TestRunCommerceHarness(t *testing.T) {
	report := RunCommerceHarness(repoRoot(t))
	if !report.Passed {
		t.Fatalf("commerce harness failed: %+v", report.Results)
	}
	if report.Name != "commerce" {
		t.Fatalf("harness name = %q", report.Name)
	}
	if len(report.Results) < 7 {
		t.Fatalf("commerce harness result count = %d", len(report.Results))
	}
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
