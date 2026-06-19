package commands

import "sort"

func Specs() []Spec {
	specs := []Spec{
		{Name: "display_sleep", Summary: "Put the display to sleep", PayloadFields: []string{}},
		{Name: "mac_sleep", Summary: "Put the Mac to sleep", PayloadFields: []string{}},
		{Name: "movie_mode", Summary: "Apply dry-run movie mode actions", PayloadFields: []string{}},
		{Name: "open_coupang_play", Summary: "Open Coupang Play", PayloadFields: []string{}},
		{Name: "open_disney_plus", Summary: "Open Disney+", PayloadFields: []string{}},
		{Name: "open_netflix", Summary: "Open Netflix", PayloadFields: []string{}},
		{Name: "open_ott", Summary: "Open a supported OTT service", PayloadFields: []string{"service"}},
		{Name: "open_url", Summary: "Open a safe http or https URL", PayloadFields: []string{"url"}},
		{Name: "open_tving", Summary: "Open TVING", PayloadFields: []string{}},
		{Name: "open_wavve", Summary: "Open Wavve", PayloadFields: []string{}},
		{Name: "open_youtube", Summary: "Open YouTube", PayloadFields: []string{}},
		{Name: "open_youtube_search", Summary: "Open a YouTube search", PayloadFields: []string{"query"}},
		{Name: "sleep_mode", Summary: "Apply dry-run sleep mode actions", PayloadFields: []string{}},
		{Name: "volume_down", Summary: "Lower output volume by a step", PayloadFields: []string{"step"}},
		{Name: "volume_mute", Summary: "Mute output volume", PayloadFields: []string{}},
		{Name: "volume_set", Summary: "Set output volume to 0..100", PayloadFields: []string{"level"}},
		{Name: "volume_up", Summary: "Raise output volume by a step", PayloadFields: []string{"step"}},
	}
	sort.Slice(specs, func(i, j int) bool { return specs[i].Name < specs[j].Name })
	return specs
}
