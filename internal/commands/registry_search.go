package commands

import (
	"fmt"
	"net/url"
	"strings"
)

func youtubeSearchPlan(command string, payload []byte) (Plan, error) {
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
}
