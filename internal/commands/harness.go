package commands

import "strings"

type HarnessCase struct {
	Name       string `json:"name"`
	Command    string `json:"command"`
	Payload    string `json:"payload"`
	ShouldPass bool   `json:"should_pass"`
	Contains   string `json:"contains,omitempty"`
}

type HarnessCaseResult struct {
	Name    string `json:"name"`
	Passed  bool   `json:"passed"`
	Message string `json:"message"`
}

type HarnessReport struct {
	Name    string              `json:"name"`
	Passed  bool                `json:"passed"`
	Results []HarnessCaseResult `json:"results"`
}

func RunHomeHarness() HarnessReport {
	cases := []HarnessCase{
		{Name: "open_youtube empty payload success", Command: "open-youtube", Payload: `{}`, ShouldPass: true, Contains: "https://www.youtube.com"},
		{Name: "open_youtube_search lofi music success", Command: "open-youtube-search", Payload: `{"query":"lofi music"}`, ShouldPass: true, Contains: "search_query=lofi+music"},
		{Name: "open_ott netflix success", Command: "open-ott", Payload: `{"service":"netflix"}`, ShouldPass: true, Contains: "https://www.netflix.com"},
		{Name: "open_ott unknown fail", Command: "open-ott", Payload: `{"service":"unknown"}`, ShouldPass: false},
		{Name: "volume_set 30 success", Command: "volume-set", Payload: `{"level":30}`, ShouldPass: true, Contains: "30"},
		{Name: "volume_set 101 fail", Command: "volume-set", Payload: `{"level":101}`, ShouldPass: false},
		{Name: "volume_up step 10 success", Command: "volume-up", Payload: `{"step":10}`, ShouldPass: true, Contains: "+ 10"},
		{Name: "volume_down step 10 success", Command: "volume-down", Payload: `{"step":10}`, ShouldPass: true, Contains: "- 10"},
		{Name: "display_sleep success", Command: "display-sleep", Payload: `{}`, ShouldPass: true, Contains: "displaysleepnow"},
		{Name: "open_url https success", Command: "open-url", Payload: `{"url":"https://example.com"}`, ShouldPass: true, Contains: "https://example.com"},
		{Name: "open_url javascript fail", Command: "open-url", Payload: `{"url":"javascript:alert(1)"}`, ShouldPass: false},
		{Name: "movie_mode dry-run success", Command: "movie-mode", Payload: `{}`, ShouldPass: true, Contains: "movie_volume"},
		{Name: "sleep_mode dry-run success", Command: "sleep-mode", Payload: `{}`, ShouldPass: true, Contains: "display_sleep"},
	}

	report := HarnessReport{Name: "home", Passed: true}
	for _, tc := range cases {
		result := HarnessCaseResult{Name: tc.Name}
		plan, err := Build(tc.Command, []byte(tc.Payload))
		if tc.ShouldPass {
			if err != nil {
				result.Message = err.Error()
			} else if tc.Contains != "" && !strings.Contains(planText(plan), tc.Contains) {
				result.Message = "expected output to contain " + tc.Contains
			} else {
				result.Passed = true
				result.Message = "ok"
			}
		} else {
			if err != nil {
				result.Passed = true
				result.Message = "failed safely: " + err.Error()
			} else {
				result.Message = "expected safe failure but command passed"
			}
		}
		if !result.Passed {
			report.Passed = false
		}
		report.Results = append(report.Results, result)
	}
	return report
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
