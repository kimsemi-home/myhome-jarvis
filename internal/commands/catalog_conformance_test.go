package commands

import (
	"reflect"
	"sort"
	"testing"
)

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
