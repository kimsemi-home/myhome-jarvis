package knowledge

import "time"

func Verify(root string) (VerifyReport, error) {
	report := VerifyReport{OK: true, CheckedAt: time.Now().UTC().Format(time.RFC3339)}
	registry, err := readRegistryUnchecked(root)
	if err != nil {
		report.OK = false
		report.Checks = append(report.Checks, Check{
			Name: "concept artifact", Status: "fail", Message: err.Error(),
		})
		return report, nil
	}
	report.ContextCount = len(registry.BoundedContexts)
	report.ConceptCount = len(registry.Concepts)
	report.EventCount = len(registry.DomainEvents)
	report.HarnessCount = len(registry.HarnessCaseContracts)
	return verifyRegistry(root, registry, report), nil
}

func verifyRegistry(root string, registry Registry, report VerifyReport) VerifyReport {
	failures := registryFailures(root, registry)
	if len(failures) == 0 {
		report.Checks = append(report.Checks,
			Check{Name: "bounded contexts", Status: "pass"},
			Check{Name: "ddd kinds", Status: "pass"},
			Check{Name: "domain events", Status: "pass"},
			Check{Name: "harness case contracts", Status: "pass"},
			Check{Name: "duplicate concepts", Status: "pass"},
			Check{Name: "registered domain terms", Status: "pass"},
			Check{Name: "alias drift", Status: "pass"},
			Check{Name: "generated artifact contracts", Status: "pass"},
			Check{Name: "knowledge index schema", Status: "pass"},
		)
		return report
	}
	report.OK = false
	for _, failure := range failures {
		report.Checks = append(report.Checks, Check{Name: "ddd verify", Status: "fail", Message: failure})
	}
	return report
}
