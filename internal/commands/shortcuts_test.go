package commands

import "testing"

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
