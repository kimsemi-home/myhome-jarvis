package codexcost

import "testing"

func TestAttributeCostRejectsUnsafeSubjects(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	for _, subject := range []string{
		"https://linear.example/private",
		"/private/local",
		"repo:../secret",
	} {
		payload := `{"scope":"repo","subject_key":"` + subject + `",` +
			`"unit_kind":"codex_tokens","amount":1,"basis":"manual",` +
			`"evidence_refs":["docs/codex-cost-governor.md"]}`
		if _, err := AttributeCost(root, []byte(payload)); err == nil {
			t.Fatalf("expected unsafe subject to fail: %s", subject)
		}
	}
}

func TestAttributeCostRejectsRawPrivateFields(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	payload := `{"scope":"repo","subject_key":"repo:public",` +
		`"unit_kind":"codex_tokens","amount":1,"basis":"manual",` +
		`"raw_prompt":"private",` +
		`"evidence_refs":["docs/codex-cost-governor.md"]}`
	if _, err := AttributeCost(root, []byte(payload)); err == nil {
		t.Fatal("expected raw private field to fail")
	}
}

func TestAttributeCostRejectsUnsafeCostRefs(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	for _, costRef := range []string{"https://example.invalid/x", "/local/ref", "bad ref"} {
		payload := `{"scope":"repo","subject_key":"repo:public",` +
			`"cost_ref":"` + costRef + `",` +
			`"unit_kind":"codex_tokens","amount":1,"basis":"manual",` +
			`"evidence_refs":["docs/codex-cost-governor.md"]}`
		if _, err := AttributeCost(root, []byte(payload)); err == nil {
			t.Fatalf("expected unsafe cost ref to fail: %s", costRef)
		}
	}
}
