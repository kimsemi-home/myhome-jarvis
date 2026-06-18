package controlplane

import "fmt"

func errForbiddenPublicMarker() error {
	return fmt.Errorf("control-plane manifest contains forbidden public marker")
}
