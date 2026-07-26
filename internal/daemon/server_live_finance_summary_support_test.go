package daemon

import "net/http"

const daemonActiveLedger = `{"at":"2026-06-19T00:00:00Z","consent_kind":"finance_connector","subject_scope":"user","status":"granted","review_status":"approved","authority_profile":"finance_review_only","evidence_refs":["docs/finance-consent.md"]}
{"at":"2026-06-19T00:01:00Z","consent_kind":"spouse_scope","subject_scope":"spouse","status":"granted","review_status":"approved","authority_profile":"finance_review_only","evidence_refs":["docs/finance-consent.md"]}
{"at":"2026-06-19T00:02:00Z","consent_kind":"household_scope","subject_scope":"household","status":"granted","review_status":"approved","authority_profile":"finance_review_only","evidence_refs":["docs/finance-consent.md"]}
`

func daemonLiveProvider(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/v2/bank/accounts" {
		_, _ = writer.Write([]byte(`{"rsp_code":"00000","account_list":[{"account_num":"ACCOUNT","is_consent":true}]}`))
		return
	}
	if request.URL.Path == "/v2/bank/accounts/deposit/transactions" {
		_, _ = writer.Write([]byte(`{"rsp_code":"00000","trans_list":[{"trans_dtime":"20260726090000","trans_no":"income-1","trans_type":"03","currency_code":"KRW","trans_amt":"4500000","trans_memo":"Salary"},{"trans_dtime":"20260726100000","trans_no":"spend-1","trans_type":"02","currency_code":"KRW","trans_amt":42000,"trans_memo":"Cafe"}]}`))
		return
	}
	if request.URL.Path == "/v2/card/cards" {
		_, _ = writer.Write([]byte(`{"rsp_code":"00000","card_list":[{"card_id":"CARD-1","is_consent":true},{"card_id":"CARD-2","is_consent":false}]}`))
		return
	}
	if request.URL.Path == "/v2/card/cards/CARD-1/approval-domestic" {
		_, _ = writer.Write([]byte(`{"rsp_code":"00000","approved_list":[{"approved_num":"card-spend-1","approved_dtime":"20260726110000","status":"01","merchant_name":"Card Cafe","approved_amt":"42000","currency_code":"KRW"}]}`))
		return
	}
	if request.URL.Path == "/v2/card/cards/CARD-1/approval-overseas" {
		_, _ = writer.Write([]byte(`{"rsp_code":"00000","approved_list":[{"approved_num":"card-income-1","approved_dtime":"20260726120000","status":"02","merchant_name":"Card Refund","approved_amt":"10","currency_code":"USD","krw_amt":"13000"}]}`))
		return
	}
	http.NotFound(writer, request)
}
