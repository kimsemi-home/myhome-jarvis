package financeconnector

import (
	"net/http"
	"strings"
	"testing"
)

func assertMyDataHeaders(t *testing.T, request *http.Request, token string) {
	t.Helper()
	if request.Header.Get("Authorization") != "Bearer "+token {
		t.Fatalf("authorization header = %q", request.Header.Get("Authorization"))
	}
	if request.Header.Get("x-api-type") != "2" {
		t.Fatalf("x-api-type = %q", request.Header.Get("x-api-type"))
	}
	transactionID := request.Header.Get("x-api-tran-id")
	if len(transactionID) != 25 || !strings.HasPrefix(transactionID, "MHJ") {
		t.Fatalf("x-api-tran-id = %q", transactionID)
	}
}
