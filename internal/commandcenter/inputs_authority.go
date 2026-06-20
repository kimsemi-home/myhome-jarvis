package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/authority"

type authorityStatus = authority.Status
type authorityReviewPlanStatus = authority.ReviewPlanStatus

func collectAuthorityInputs(root string, in *inputs) error {
	var err error
	if in.Authority, err = authority.StatusForRoot(root); err != nil {
		return err
	}
	if in.AuthorityReview, err = authority.ReviewPlanForRoot(root); err != nil {
		return err
	}
	return nil
}
