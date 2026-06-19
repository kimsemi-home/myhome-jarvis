package linear

import (
	"net/http"
	"testing"
)

func commentReplayResponse(t *testing.T, variables map[string]string) *http.Response {
	t.Helper()
	if variables["issueId"] != "KIM-16" || variables["body"] != "Started work" {
		t.Fatalf("unexpected comment variables: %#v", variables)
	}
	return linearResponse(250, `{"data":{"commentCreate":{"success":true,"comment":{"id":"comment-id","createdAt":"2026-06-14T00:00:00.000Z"}}}}`)
}

func transitionReplayResponse(t *testing.T, variables map[string]string) *http.Response {
	t.Helper()
	if variables["issueId"] != "KIM-16" || variables["stateId"] != "done-state" {
		t.Fatalf("unexpected transition variables: %#v", variables)
	}
	return linearResponse(248, `{"data":{"issueUpdate":{"success":true,"issue":{"id":"issue-id","identifier":"KIM-16","title":"Replay","state":{"id":"done-state","name":"Done","type":"completed"}}}}}`)
}
