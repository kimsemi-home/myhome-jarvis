package externalevidence

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestRepoSplitDecisionPacketDoesNotExposePrivateLakePaths(t *testing.T) {
	packet, err := RepoSplitDecisionPacketForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(packet)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{
		"data/private",
		"raw_layer_path",
		"bronze_layer_path",
		"silver_layer_path",
		"gold_layer_path",
		"kim" + "jooyoon",
		"kim" + "-joo-yoon",
		"/" + "Users" + "/",
		"token",
		"private_key",
	} {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("repo split decision leaked %q in %s", forbidden, body)
		}
	}
}
