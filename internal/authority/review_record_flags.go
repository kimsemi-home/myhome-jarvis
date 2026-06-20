package authority

import "fmt"

func validateExplicitNonApproval(request ReviewRecordRequest) error {
	if request.ApprovalGranted == nil ||
		request.ExternalWritesAllowed == nil ||
		request.SelfApprovalAllowed == nil {
		return fmt.Errorf("authority review record requires explicit non-approval flags")
	}
	if *request.ApprovalGranted ||
		*request.ExternalWritesAllowed ||
		*request.SelfApprovalAllowed {
		return fmt.Errorf("authority review record must not grant approval or external writes")
	}
	return nil
}
