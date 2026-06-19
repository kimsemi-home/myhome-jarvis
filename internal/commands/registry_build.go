package commands

import "fmt"

func Build(name string, payload []byte) (Plan, error) {
	command := normalizeName(name)
	switch command {
	case "open_coupang_play", "open_disney_plus", "open_netflix",
		"open_youtube", "open_youtube_search", "open_ott",
		"open_url", "open_tving", "open_wavve":
		return buildOpenCommand(command, payload)
	case "volume_set", "volume_up", "volume_down", "volume_mute":
		return buildVolumeCommand(command, payload)
	case "display_sleep":
		return argvPlan(command, "display_sleep", []string{"pmset", "displaysleepnow"}), nil
	case "mac_sleep":
		return argvPlan(command, "mac_sleep", []string{"pmset", "sleepnow"}), nil
	case "movie_mode", "sleep_mode":
		return buildModeCommand(command)
	default:
		return Plan{}, fmt.Errorf("unknown command %q", name)
	}
}
