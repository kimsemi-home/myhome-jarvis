package linear

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (fn roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return fn(request)
}

func TestAuthorizationHeaderPreservesPersonalAPIKey(t *testing.T) {
	got := authorizationHeader("linear-example-key")
	if got != "linear-example-key" {
		t.Fatalf("authorization header = %q", got)
	}
}

func TestQueryViewerUsesGraphQLEndpointAndAuthorization(t *testing.T) {
	client := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		if request.URL.String() != Endpoint {
			t.Fatalf("endpoint = %s", request.URL.String())
		}
		if request.Header.Get("Authorization") != "linear-example-key" {
			t.Fatalf("authorization = %q", request.Header.Get("Authorization"))
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"X-RateLimit-Remaining": []string{"4999"}},
			Body: io.NopCloser(strings.NewReader(`{
				"data": {
					"viewer": {"id": "user-id", "name": "Example User"},
					"teams": {"nodes": [{"id": "team-id", "name": "Home"}]}
				}
			}`)),
		}, nil
	})}

	viewer, teams, status, remaining, err := queryViewer(context.Background(), client, "linear-example-key")
	if err != nil {
		t.Fatal(err)
	}
	if status != http.StatusOK || remaining != 4999 {
		t.Fatalf("status=%d remaining=%d", status, remaining)
	}
	if viewer.Name != "Example User" || len(teams) != 1 || teams[0].Name != "Home" {
		t.Fatalf("unexpected viewer/team: %#v %#v", viewer, teams)
	}
}
