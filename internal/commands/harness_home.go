package commands

func RunHomeHarness() HarnessReport {
	return runBuildHarness("home", []HarnessCase{
		{Name: "open_youtube empty payload success", Command: "open-youtube", Payload: `{}`, ShouldPass: true, Contains: "https://www.youtube.com"},
		{Name: "open_youtube_search lofi music success", Command: "open-youtube-search", Payload: `{"query":"lofi music"}`, ShouldPass: true, Contains: "search_query=lofi+music"},
		{Name: "open_ott netflix success", Command: "open-ott", Payload: `{"service":"netflix"}`, ShouldPass: true, Contains: "https://www.netflix.com"},
		{Name: "open_ott unknown fail", Command: "open-ott", Payload: `{"service":"unknown"}`, ShouldPass: false},
		{Name: "open_netflix shortcut success", Command: "open-netflix", Payload: `{}`, ShouldPass: true, Contains: "https://www.netflix.com"},
		{Name: "open_disney_plus shortcut success", Command: "open-disney-plus", Payload: `{}`, ShouldPass: true, Contains: "https://www.disneyplus.com"},
		{Name: "open_tving shortcut success", Command: "open-tving", Payload: `{}`, ShouldPass: true, Contains: "https://www.tving.com"},
		{Name: "open_wavve shortcut success", Command: "open-wavve", Payload: `{}`, ShouldPass: true, Contains: "https://www.wavve.com"},
		{Name: "open_coupang_play shortcut success", Command: "open-coupang-play", Payload: `{}`, ShouldPass: true, Contains: "https://www.coupangplay.com"},
		{Name: "volume_set 30 success", Command: "volume-set", Payload: `{"level":30}`, ShouldPass: true, Contains: "30"},
		{Name: "volume_set 101 fail", Command: "volume-set", Payload: `{"level":101}`, ShouldPass: false},
		{Name: "volume_up step 10 success", Command: "volume-up", Payload: `{"step":10}`, ShouldPass: true, Contains: "+ 10"},
		{Name: "volume_down step 10 success", Command: "volume-down", Payload: `{"step":10}`, ShouldPass: true, Contains: "- 10"},
		{Name: "display_sleep success", Command: "display-sleep", Payload: `{}`, ShouldPass: true, Contains: "displaysleepnow"},
		{Name: "open_url https success", Command: "open-url", Payload: `{"url":"https://example.com"}`, ShouldPass: true, Contains: "https://example.com"},
		{Name: "open_url javascript fail", Command: "open-url", Payload: `{"url":"javascript:alert(1)"}`, ShouldPass: false},
		{Name: "movie_mode dry-run success", Command: "movie-mode", Payload: `{}`, ShouldPass: true, Contains: "movie_volume"},
		{Name: "sleep_mode dry-run success", Command: "sleep-mode", Payload: `{}`, ShouldPass: true, Contains: "display_sleep"},
	})
}
