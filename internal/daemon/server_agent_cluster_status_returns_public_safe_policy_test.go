package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAgentClusterStatusReturnsPublicSafePolicy(t *testing.T) {
	server, err := New(DefaultConfig(repoRoot(t), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/agent-cluster/status", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"context": "AgentCluster"`,
		`"public_safe": true`,
		`"external_agent_execution_allowed": false`,
		`"self_approval_allowed": false`,
		`"authority_gate_required": true`,
		`"generated_path": "generated/agent_cluster.generated.json"`,
		`"key": "evidence_first"`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{
		`"token":`,
		`"secret":`,
		`"credential":`,
		`"cookie":`,
		`"generated_path": "/"`,
		`"generated_path": "\\"`,
	} {
		if bytes.Contains([]byte(body), []byte(forbidden)) {
			t.Fatalf("agent cluster status leaked %s in %s", forbidden, body)
		}
	}
}
