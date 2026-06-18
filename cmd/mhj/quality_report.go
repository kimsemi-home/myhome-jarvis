package main

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
	"github.com/kimsemi-home/myhome-jarvis/internal/qualitylog"
)

type qualityStep struct {
	Name    string   `json:"name"`
	Status  string   `json:"status"`
	Command []string `json:"-"`
	Output  string   `json:"-"`
}

type qualityReport struct {
	OK    bool          `json:"ok"`
	Steps []qualityStep `json:"steps"`
}

func qualityEvidenceSteps(steps []qualityStep) []qualitylog.Step {
	evidence := make([]qualitylog.Step, 0, len(steps))
	for _, step := range steps {
		evidence = append(evidence, qualitylog.Step{Name: step.Name, Status: step.Status})
	}
	return evidence
}

func (report *qualityReport) addHarness(name string, harness commands.HarnessReport) {
	if harness.Passed {
		report.Steps = append(report.Steps, qualityStep{Name: name, Status: "pass"})
		return
	}
	report.OK = false
	report.Steps = append(report.Steps, qualityStep{Name: name, Status: "fail"})
}

func (report *qualityReport) addCheck(name string, err error) {
	if err == nil {
		report.Steps = append(report.Steps, qualityStep{Name: name, Status: "pass"})
		return
	}
	report.OK = false
	report.Steps = append(report.Steps, qualityStep{Name: name, Status: "fail", Output: err.Error()})
}
