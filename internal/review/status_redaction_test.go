package review

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
)

func TestStatusForRootCountsMissingReviewerAndEvidenceWithoutLeaks(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/review/queue.jsonl", missingReviewerJSON+"\n"+missingEvidenceJSON+"\n")

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.MissingReviewerCount != 1 || status.MissingEvidenceCount != 1 || status.Count != 0 {
		t.Fatalf("status = %#v", status)
	}
	payload, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	body := strings.ToLower(string(payload))
	for _, forbidden := range forbiddenReviewStatusMarkers() {
		if strings.Contains(body, forbidden) {
			t.Fatalf("review status leaked %q in %s", forbidden, body)
		}
	}
}

func forbiddenReviewStatusMarkers() []string {
	return []string{
		"raw_review", "raw_review_notes", "raw_rationale",
		"reviewer_identity", "private reviewer", "private rationale",
		"data/private/review/raw-notes.json", "evidence_ref", "evidence_refs",
		"token", "secret", "credential", "linear.app",
		string(filepath.Separator) + "users" + string(filepath.Separator),
	}
}
