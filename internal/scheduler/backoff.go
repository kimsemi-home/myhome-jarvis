package scheduler

import "time"

func backoff(policy Policy, failures int) time.Duration {
	if failures <= 0 {
		return policy.Interval
	}
	delay := policy.MinBackoff
	for index := 1; index < failures; index++ {
		delay *= 2
		if delay >= policy.MaxBackoff {
			return policy.MaxBackoff
		}
	}
	return delay
}
