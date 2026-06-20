package contextpack

import "time"

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	return statusFromPolicy(policy, time.Now().UTC()), nil
}
