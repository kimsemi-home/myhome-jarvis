package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

type Spec struct {
	Name          string   `json:"name"`
	Summary       string   `json:"summary"`
	PayloadFields []string `json:"payload_fields"`
}

type Invocation struct {
	Label string   `json:"label"`
	Argv  []string `json:"argv"`
	URL   string   `json:"url,omitempty"`
}

type Plan struct {
	Name           string       `json:"name"`
	DryRun         bool         `json:"dry_run"`
	ExecuteAllowed bool         `json:"execute_allowed"`
	Invocations    []Invocation `json:"invocations"`
	Executions     []Execution  `json:"executions,omitempty"`
	Warnings       []string     `json:"warnings,omitempty"`
}

var errInvalidPayload = errors.New("invalid payload")

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

func Build(name string, payload []byte) (Plan, error) {
	command := normalizeName(name)
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
		var body struct {
			Query string `json:"query"`
		}
		if err := decodePayload(payload, &body); err != nil {
			return Plan{}, err
		}
		query := strings.TrimSpace(body.Query)
		if query == "" {
			return Plan{}, fmt.Errorf("%w: query is required", errInvalidPayload)
		}
		values := url.Values{}
		values.Set("search_query", query)
		return openURLPlan(command, "https://www.youtube.com/results?"+values.Encode()), nil
	case "open_ott":
		var body struct {
			Service string `json:"service"`
		}
		if err := decodePayload(payload, &body); err != nil {
			return Plan{}, err
		}
		service := strings.ToLower(strings.TrimSpace(body.Service))
		target, ok := ottURLs()[service]
		if !ok {
			return Plan{}, fmt.Errorf("%w: unsupported ott service %q", errInvalidPayload, body.Service)
		}
		return openURLPlan(command, target), nil
	case "open_url":
		var body struct {
			URL string `json:"url"`
		}
		if err := decodePayload(payload, &body); err != nil {
			return Plan{}, err
		}
		target, err := validateHTTPURL(body.URL)
		if err != nil {
			return Plan{}, err
		}
		return openURLPlan(command, target), nil
	case "open_tving":
		return ottShortcutPlan(command, "tving")
	case "open_wavve":
		return ottShortcutPlan(command, "wavve")
	case "volume_set":
		var body struct {
			Level *int `json:"level"`
		}
		if err := decodePayload(payload, &body); err != nil {
			return Plan{}, err
		}
		if body.Level == nil {
			return Plan{}, fmt.Errorf("%w: level is required", errInvalidPayload)
		}
		return volumeSetPlan(command, *body.Level)
	case "volume_up":
		step, err := volumeStep(payload)
		if err != nil {
			return Plan{}, err
		}
		return appleScriptPlan(command, fmt.Sprintf("set volume output volume ((output volume of (get volume settings)) + %d)", step)), nil
	case "volume_down":
		step, err := volumeStep(payload)
		if err != nil {
			return Plan{}, err
		}
		return appleScriptPlan(command, fmt.Sprintf("set volume output volume ((output volume of (get volume settings)) - %d)", step)), nil
	case "volume_mute":
		return appleScriptPlan(command, "set volume output muted true"), nil
	case "display_sleep":
		return argvPlan(command, "display_sleep", []string{"pmset", "displaysleepnow"}), nil
	case "mac_sleep":
		return argvPlan(command, "mac_sleep", []string{"pmset", "sleepnow"}), nil
	case "movie_mode":
		return Plan{
			Name:   command,
			DryRun: true,
			Invocations: []Invocation{
				{Label: "movie_volume", Argv: []string{"osascript", "-e", "set volume output volume 35"}},
				{Label: "open_youtube", Argv: []string{"open", "https://www.youtube.com"}, URL: "https://www.youtube.com"},
			},
		}, nil
	case "sleep_mode":
		return Plan{
			Name:   command,
			DryRun: true,
			Invocations: []Invocation{
				{Label: "mute", Argv: []string{"osascript", "-e", "set volume output muted true"}},
				{Label: "display_sleep", Argv: []string{"pmset", "displaysleepnow"}},
			},
		}, nil
	default:
		return Plan{}, fmt.Errorf("unknown command %q", name)
	}
}

func ottShortcutPlan(name string, service string) (Plan, error) {
	target, ok := ottURLs()[service]
	if !ok {
		return Plan{}, fmt.Errorf("missing OTT shortcut target for %q", service)
	}
	return openURLPlan(name, target), nil
}

func WithExecuteAllowed(plan Plan, executeAllowed bool) Plan {
	plan.ExecuteAllowed = executeAllowed
	return plan
}

func normalizeName(name string) string {
	return strings.ReplaceAll(strings.TrimSpace(strings.ToLower(name)), "-", "_")
}

func decodePayload(payload []byte, target any) error {
	if len(strings.TrimSpace(string(payload))) == 0 {
		payload = []byte("{}")
	}
	if err := json.Unmarshal(payload, target); err != nil {
		return fmt.Errorf("%w: %v", errInvalidPayload, err)
	}
	return nil
}

func openURLPlan(name string, target string) Plan {
	return Plan{
		Name:   name,
		DryRun: true,
		Invocations: []Invocation{
			{Label: name, Argv: []string{"open", target}, URL: target},
		},
	}
}

func argvPlan(name string, label string, argv []string) Plan {
	return Plan{
		Name:        name,
		DryRun:      true,
		Invocations: []Invocation{{Label: label, Argv: argv}},
	}
}

func appleScriptPlan(name string, script string) Plan {
	return argvPlan(name, name, []string{"osascript", "-e", script})
}

func volumeSetPlan(name string, level int) (Plan, error) {
	if level < 0 || level > 100 {
		return Plan{}, fmt.Errorf("%w: level must be between 0 and 100", errInvalidPayload)
	}
	return appleScriptPlan(name, fmt.Sprintf("set volume output volume %d", level)), nil
}

func volumeStep(payload []byte) (int, error) {
	var body struct {
		Step *int `json:"step"`
	}
	if err := decodePayload(payload, &body); err != nil {
		return 0, err
	}
	step := 10
	if body.Step != nil {
		step = *body.Step
	}
	if step < 1 || step > 100 {
		return 0, fmt.Errorf("%w: step must be between 1 and 100", errInvalidPayload)
	}
	return step, nil
}

func validateHTTPURL(raw string) (string, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return "", fmt.Errorf("%w: url is required", errInvalidPayload)
	}
	parsed, err := url.Parse(trimmed)
	if err != nil {
		return "", fmt.Errorf("%w: %v", errInvalidPayload, err)
	}
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return "", fmt.Errorf("%w: only http and https URLs are allowed", errInvalidPayload)
	}
	if parsed.Host == "" {
		return "", fmt.Errorf("%w: URL host is required", errInvalidPayload)
	}
	return parsed.String(), nil
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
