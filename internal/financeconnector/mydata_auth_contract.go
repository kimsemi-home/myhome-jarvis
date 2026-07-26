package financeconnector

const myDataSignVerificationPath = "/v1/ca/sign_verification"

// SignVerificationRequest is the provider-issued consent proof for the
// MyData authentication plane. Signed consent stays server-side and is never
// returned through daemon status or the Flutter client.
type SignVerificationRequest struct {
	CertTxID         string `json:"cert_tx_id"`
	TxID             string `json:"tx_id"`
	SignedConsentLen string `json:"signed_consent_len"`
	SignedConsent    string `json:"signed_consent"`
	ConsentType      string `json:"consent_type"`
	ConsentLen       string `json:"consent_len"`
	Consent          string `json:"consent"`
}

type signVerificationResponse struct {
	TxID    string `json:"tx_id"`
	RspCode string `json:"rsp_code"`
	RspMsg  string `json:"rsp_msg"`
	Result  string `json:"result"`
	UserCI  string `json:"user_ci"`
}
