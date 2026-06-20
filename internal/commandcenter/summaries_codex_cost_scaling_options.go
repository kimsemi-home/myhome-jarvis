package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/codexcost"

func recommendedScalingOptionKeys(options []codexcost.ScalingOption) []string {
	keys := []string{}
	for _, option := range options {
		if option.Recommended {
			keys = append(keys, option.Key)
		}
	}
	return keys
}

func grantingScalingOptionCount(options []codexcost.ScalingOption) int {
	count := 0
	for _, option := range options {
		if option.ThisPacketGrantsSpend ||
			option.AllowsPaidExpansion ||
			option.AllowsExternalTooling ||
			option.AllowsWorkflowChanges ||
			option.AllowsSelfApproval {
			count++
		}
	}
	return count
}
