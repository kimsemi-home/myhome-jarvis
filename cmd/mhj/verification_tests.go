package main

import "fmt"

func verifyTestManifest(tests verificationTestsFile) error {
	covered := map[string]bool{}
	for _, test := range tests.Tests {
		if test.ID == "" || test.Kind == "" || test.Evidence == "" {
			return fmt.Errorf("verification test case must declare id, kind, and evidence")
		}
		covered[test.ID] = true
	}
	for _, id := range requiredVerificationTests() {
		if !covered[id] {
			return fmt.Errorf("verification test manifest missing %q", id)
		}
	}
	return nil
}

func stringSet(values []string) map[string]bool {
	set := map[string]bool{}
	for _, value := range values {
		set[value] = true
	}
	return set
}
