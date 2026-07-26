package financeconnector

import (
	"fmt"

	"github.com/kimsemi-home/myhome-jarvis/internal/financeconsent"
)

func requireActiveConsent(root string) error {
	status, err := financeconsent.StatusForRoot(root)
	if err != nil {
		return err
	}
	if status.ReadinessState != "ready_read_only" {
		return fmt.Errorf("live MyData requires ready_read_only consent status")
	}
	return nil
}
