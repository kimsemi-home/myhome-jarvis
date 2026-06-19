package linear

import "encoding/json"

type graphQLRequest struct {
	Query     string `json:"query"`
	Variables any    `json:"variables,omitempty"`
}

type graphQLError struct {
	Message string `json:"message"`
}

type graphQLEnvelope struct {
	Data   json.RawMessage `json:"data"`
	Errors []graphQLError  `json:"errors"`
}

type viewerResponse struct {
	Viewer *ViewerStatus `json:"viewer"`
	Teams  struct {
		Nodes []TeamStatus `json:"nodes"`
	} `json:"teams"`
}
