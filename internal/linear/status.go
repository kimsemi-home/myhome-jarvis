package linear

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const Endpoint = "https://api.linear.app/graphql"

type Status struct {
	Mode               string        `json:"mode"`
	TokenConfigured    bool          `json:"token_configured"`
	TokenSource        string        `json:"token_source,omitempty"`
	Synced             bool          `json:"synced"`
	QueuePath          string        `json:"queue_path"`
	Endpoint           string        `json:"endpoint,omitempty"`
	HTTPStatus         int           `json:"http_status,omitempty"`
	RateLimitRemaining int           `json:"rate_limit_remaining,omitempty"`
	Viewer             *ViewerStatus `json:"viewer,omitempty"`
	Teams              []TeamStatus  `json:"teams,omitempty"`
	Message            string        `json:"message"`
}

type ViewerStatus struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

type TeamStatus struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type OfflineEvent struct {
	At      string `json:"at"`
	Kind    string `json:"kind"`
	Message string `json:"message"`
	Synced  bool   `json:"synced"`
}

type tokenConfig struct {
	Value  string
	Source string
}

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

type graphQLResponse struct {
	Data struct {
		Viewer *ViewerStatus `json:"viewer"`
		Teams  struct {
			Nodes []TeamStatus `json:"nodes"`
		} `json:"teams"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

func CurrentStatus(root string) Status {
	return StatusForRequest(context.Background(), root, http.DefaultClient)
}

func StatusForRequest(parent context.Context, root string, client *http.Client) Status {
	queuePath := filepath.Join(root, "data", "private", "linear-offline-queue.jsonl")
	status := Status{
		Mode:      "offline",
		Synced:    false,
		QueuePath: queuePath,
		Endpoint:  Endpoint,
	}
	token, err := loadToken(root)
	if err != nil {
		status.Message = "No Linear token found. Continuing in offline mode."
		return status
	}
	status.Mode = "configured"
	status.TokenConfigured = true
	status.TokenSource = token.Source

	ctx, cancel := context.WithTimeout(parent, 15*time.Second)
	defer cancel()
	viewer, teams, httpStatus, remaining, err := queryViewer(ctx, client, token.Value)
	status.HTTPStatus = httpStatus
	status.RateLimitRemaining = remaining
	if err != nil {
		status.Mode = "offline"
		status.Message = "Linear GraphQL status check failed; continuing with offline fallback: " + err.Error()
		return status
	}
	status.Mode = "online"
	status.Synced = true
	status.Viewer = viewer
	status.Teams = teams
	status.Message = "Linear GraphQL status check succeeded."
	return status
}

func AppendOfflineEvent(root string, kind string, message string) error {
	return AppendOfflineAction(root, kind, message, nil)
}

func AppendOfflineAction(root string, kind string, message string, payload any) error {
	queuePath := filepath.Join(root, "data", "private", "linear-offline-queue.jsonl")
	if err := os.MkdirAll(filepath.Dir(queuePath), 0o700); err != nil {
		return err
	}
	file, err := os.OpenFile(queuePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		return err
	}
	defer file.Close()

	event := OfflineEvent{
		At:      time.Now().UTC().Format(time.RFC3339),
		Kind:    kind,
		Message: message,
		Synced:  false,
	}
	if payload != nil {
		return json.NewEncoder(file).Encode(struct {
			OfflineEvent
			Payload any `json:"payload"`
		}{OfflineEvent: event, Payload: payload})
	}
	encoder := json.NewEncoder(file)
	return encoder.Encode(event)
}

func loadToken(root string) (tokenConfig, error) {
	if value := strings.TrimSpace(os.Getenv("LINEAR_API_KEY")); value != "" {
		return tokenConfig{Value: value, Source: "env:LINEAR_API_KEY"}, nil
	}
	tokenPath := filepath.Join(root, "data", "private", "linear-token.txt")
	data, err := os.ReadFile(tokenPath)
	if err != nil {
		return tokenConfig{}, err
	}
	value := strings.TrimSpace(string(data))
	if value == "" {
		return tokenConfig{}, errors.New("Linear token file is empty")
	}
	return tokenConfig{Value: value, Source: "file:data/private/linear-token.txt"}, nil
}

func queryViewer(ctx context.Context, client *http.Client, token string) (*ViewerStatus, []TeamStatus, int, int, error) {
	var decoded viewerResponse
	httpStatus, remaining, err := doGraphQL(ctx, client, token, `query Me { viewer { id name email } teams(first: 1) { nodes { id name } } }`, nil, &decoded)
	if err != nil {
		return nil, nil, httpStatus, remaining, err
	}
	if decoded.Viewer == nil {
		return nil, nil, httpStatus, remaining, errors.New("Linear response did not include viewer")
	}
	return decoded.Viewer, decoded.Teams.Nodes, httpStatus, remaining, nil
}

func doGraphQL(ctx context.Context, client *http.Client, token string, query string, variables any, target any) (int, int, error) {
	if client == nil {
		client = http.DefaultClient
	}
	body, err := json.Marshal(graphQLRequest{Query: query, Variables: variables})
	if err != nil {
		return 0, 0, err
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, Endpoint, bytes.NewReader(body))
	if err != nil {
		return 0, 0, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", authorizationHeader(token))

	response, err := client.Do(request)
	if err != nil {
		return 0, 0, err
	}
	defer response.Body.Close()
	remaining := rateLimitRemaining(response.Header)
	payload, err := io.ReadAll(io.LimitReader(response.Body, 1<<20))
	if err != nil {
		return response.StatusCode, remaining, err
	}
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return response.StatusCode, remaining, fmt.Errorf("Linear HTTP status %d", response.StatusCode)
	}

	var envelope graphQLEnvelope
	if err := json.Unmarshal(payload, &envelope); err != nil {
		return response.StatusCode, remaining, err
	}
	if len(envelope.Errors) > 0 {
		return response.StatusCode, remaining, fmt.Errorf("Linear GraphQL error: %s", envelope.Errors[0].Message)
	}
	if len(envelope.Data) == 0 || string(envelope.Data) == "null" {
		return response.StatusCode, remaining, errors.New("Linear response did not include data")
	}
	if target != nil {
		if err := json.Unmarshal(envelope.Data, target); err != nil {
			return response.StatusCode, remaining, err
		}
	}
	return response.StatusCode, remaining, nil
}

func authorizationHeader(token string) string {
	trimmed := strings.TrimSpace(token)
	if strings.HasPrefix(strings.ToLower(trimmed), "bearer ") {
		return trimmed
	}
	return trimmed
}

func rateLimitRemaining(header http.Header) int {
	for _, name := range []string{
		"X-RateLimit-Remaining",
		"X-RateLimit-Requests-Remaining",
		"Linear-RateLimit-Remaining",
	} {
		value := strings.TrimSpace(headerValue(header, name))
		if value == "" {
			continue
		}
		parsed, err := strconv.Atoi(value)
		if err == nil {
			return parsed
		}
	}
	return 0
}

func headerValue(header http.Header, name string) string {
	if value := header.Get(name); value != "" {
		return value
	}
	for key, values := range header {
		if !strings.EqualFold(key, name) || len(values) == 0 {
			continue
		}
		return values[0]
	}
	return ""
}
