package commands

import "strings"

func runBuildHarness(name string, cases []HarnessCase) HarnessReport {
	report := HarnessReport{Name: name, Passed: true}
	for _, tc := range cases {
		result := runHarnessCase(tc)
		if !result.Passed {
			report.Passed = false
		}
		report.Results = append(report.Results, result)
	}
	return report
}

func runHarnessCase(tc HarnessCase) HarnessCaseResult {
	result := HarnessCaseResult{Name: tc.Name}
	plan, err := Build(tc.Command, []byte(tc.Payload))
	if !tc.ShouldPass {
		if err != nil {
			result.Passed = true
			result.Message = "failed safely: " + err.Error()
		} else {
			result.Message = "expected safe failure but command passed"
		}
		return result
	}
	if err != nil {
		result.Message = err.Error()
	} else if tc.Contains != "" && !strings.Contains(planText(plan), tc.Contains) {
		result.Message = "expected output to contain " + tc.Contains
	} else {
		result.Passed = true
		result.Message = "ok"
	}
	return result
}

func planText(plan Plan) string {
	var b strings.Builder
	b.WriteString(plan.Name)
	for _, invocation := range plan.Invocations {
		b.WriteString(" ")
		b.WriteString(invocation.Label)
		b.WriteString(" ")
		b.WriteString(invocation.URL)
		for _, arg := range invocation.Argv {
			b.WriteString(" ")
			b.WriteString(arg)
		}
	}
	return b.String()
}
