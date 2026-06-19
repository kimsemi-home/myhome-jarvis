package commands

import "fmt"

func buildOpenCommand(command string, payload []byte) (Plan, error) {
	switch command {
	case "open_coupang_play":
		return ottShortcutPlan(command, "coupangplay")
	case "open_disney_plus":
		return ottShortcutPlan(command, "disney")
	case "open_netflix":
		return ottShortcutPlan(command, "netflix")
	case "open_youtube":
		return openURLPlan(command, "https://www.youtube.com"), nil
	case "open_youtube_search":
		return youtubeSearchPlan(command, payload)
	case "open_ott":
		return openOTTPlan(command, payload)
	case "open_url":
		return openURLPayloadPlan(command, payload)
	case "open_tving":
		return ottShortcutPlan(command, "tving")
	case "open_wavve":
		return ottShortcutPlan(command, "wavve")
	default:
		return Plan{}, fmt.Errorf("unknown open command %q", command)
	}
}
