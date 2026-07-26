package financeconnector

import "testing"

func TestMyDataDirectionMatchesBankTransactionCodes(t *testing.T) {
	expected := map[string]string{
		"02": "debit", "03": "credit", "04": "credit", "05": "debit",
		"06": "credit", "07": "debit", "09": "credit", "10": "debit",
		"98": "credit", "99": "debit",
	}
	for code, direction := range expected {
		got, err := myDataDirection(code)
		if err != nil || got != direction {
			t.Fatalf("code %s = %q, %v; want %s", code, got, err, direction)
		}
	}
	if _, err := myDataDirection("01"); err == nil {
		t.Fatal("ambiguous 신규 transaction must not be silently classified")
	}
}
