package domain

import "testing"

func assertHouseholdSummary(t *testing.T, summary Summary) {
	t.Helper()
	if len(summary.Household.Scopes) != 3 {
		t.Fatalf("household scopes = %#v", summary.Household.Scopes)
	}
	user := summary.Household.Scopes[0]
	if user.Scope != "user" || user.FinanceNetMinorUnits != -87_300 || user.PurchaseSpendMinorUnits != 3_200 {
		t.Fatalf("user scope = %#v", user)
	}
	spouse := summary.Household.Scopes[1]
	if spouse.Scope != "spouse" || spouse.FinanceRecords != 0 || spouse.PurchaseRecords != 0 {
		t.Fatalf("spouse scope = %#v", spouse)
	}
	household := summary.Household.Scopes[2]
	if household.Scope != "household" || household.FinanceNetMinorUnits != 4_346_800 || household.PurchaseSpendMinorUnits != 26_800 {
		t.Fatalf("household scope = %#v", household)
	}
}
