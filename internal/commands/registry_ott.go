package commands

import "fmt"

func ottShortcutPlan(name string, service string) (Plan, error) {
	target, ok := ottURLs()[service]
	if !ok {
		return Plan{}, fmt.Errorf("missing OTT shortcut target for %q", service)
	}
	return openURLPlan(name, target), nil
}

func ottURLs() map[string]string {
	return map[string]string{
		"coupangplay": "https://www.coupangplay.com",
		"disney":      "https://www.disneyplus.com",
		"netflix":     "https://www.netflix.com",
		"tving":       "https://www.tving.com",
		"wavve":       "https://www.wavve.com",
		"youtube":     "https://www.youtube.com",
	}
}
