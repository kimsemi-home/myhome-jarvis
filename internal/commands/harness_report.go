package commands

import "strings"

func (report *HarnessReport) addCheck(name string, passed bool, message string) {
	if strings.TrimSpace(message) == "" {
		message = "ok"
	}
	if passed {
		message = "ok"
	}
	if !passed {
		report.Passed = false
	}
	report.Results = append(report.Results, HarnessCaseResult{
		Name:    name,
		Passed:  passed,
		Message: message,
	})
}
