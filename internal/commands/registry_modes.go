package commands

import "fmt"

func buildModeCommand(command string) (Plan, error) {
	switch command {
	case "movie_mode":
		return movieModePlan(), nil
	case "sleep_mode":
		return sleepModePlan(), nil
	default:
		return Plan{}, fmt.Errorf("unknown mode command %q", command)
	}
}

func movieModePlan() Plan {
	return Plan{
		Name:   "movie_mode",
		DryRun: true,
		Invocations: []Invocation{
			{Label: "movie_volume", Argv: []string{"osascript", "-e", "set volume output volume 35"}},
			{Label: "open_youtube", Argv: []string{"open", "https://www.youtube.com"}, URL: "https://www.youtube.com"},
		},
	}
}

func sleepModePlan() Plan {
	return Plan{
		Name:   "sleep_mode",
		DryRun: true,
		Invocations: []Invocation{
			{Label: "mute", Argv: []string{"osascript", "-e", "set volume output muted true"}},
			{Label: "display_sleep", Argv: []string{"pmset", "displaysleepnow"}},
		},
	}
}
