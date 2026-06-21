package externalbootstrap

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/externalevidence"
)

func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		next := filepath.Dir(dir)
		if next == dir {
			t.Fatal("could not locate repo root")
		}
		dir = next
	}
}

func splitFixture(now time.Time) externalevidence.RepoSplitDecisionPacket {
	return externalevidence.RepoSplitDecisionPacket{
		PublicSafe:              true,
		DecisionState:           "review_only",
		FutureRepoCandidate:     "kimsemi-home/myhome-external-evidence-lake",
		RepoCreationGate:        "authority_review_required",
		PrivateLakeStaysPrivate: true,
		RawPayloadPublicAllowed: false,
		ExternalWritesAllowed:   false,
		CheckedAt:               now.UTC().Format(time.RFC3339),
	}
}
