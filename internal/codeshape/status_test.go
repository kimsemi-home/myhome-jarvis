package codeshape

import "testing"

func TestStatusForRootReportsLegacyDebtWithoutRegression(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy([]LegacyDebtFile{{Path: "src/a.go", MaxLines: 76}}))
	writeFile(t, root, "src/a.go", numberedLines(76))

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if !status.OK || status.LegacyDebtCount != 1 || status.BudgetRegressionCount != 0 {
		t.Fatalf("status = %#v", status)
	}
}

func TestStatusForRootRejectsNewOverBudgetFile(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy(nil))
	writeFile(t, root, "src/a.go", numberedLines(76))

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.OK || status.BudgetRegressionCount != 1 || len(status.Regressions) != 1 {
		t.Fatalf("status = %#v", status)
	}
}
