package codexcost

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/codexsustainability"
	"github.com/kimsemi-home/myhome-jarvis/internal/storagearchive"
)

const valueProxyMethod = "accepted_changes_plus_cache_savings_by_cost_share"

func ROISummaryForRoot(root string) (ROISummary, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return ROISummary{}, err
	}
	cost, err := StatusForRoot(root)
	if err != nil {
		return ROISummary{}, err
	}
	attribution, err := AttributionStatusForRoot(root)
	if err != nil {
		return ROISummary{}, err
	}
	sustainability, err := codexsustainability.StatusForRoot(root)
	if err != nil {
		return ROISummary{}, err
	}
	storage, err := storagearchive.StatusForRoot(root)
	if err != nil {
		return ROISummary{}, err
	}
	merge := mergeAcceptanceForRoot(root, policy)
	return buildROISummary(policy, cost, attribution, sustainability, storage, merge), nil
}
