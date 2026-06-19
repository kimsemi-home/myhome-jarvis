package linear

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func queryViewer(ctx context.Context, client *http.Client, token string) (*ViewerStatus, []TeamStatus, int, int, error) {
	var decoded viewerResponse
	httpStatus, remaining, err := doGraphQL(
		ctx, client, token,
		`query Me { viewer { id name } teams(first: 1) { nodes { id name } } }`,
		nil, &decoded,
	)
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
	return doGraphQLRequest(client, request, target)
}

func doGraphQLRequest(client *http.Client, request *http.Request, target any) (int, int, error) {
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
	return decodeGraphQLPayload(response.StatusCode, remaining, payload, target)
}
