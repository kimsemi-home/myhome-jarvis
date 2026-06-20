package codexcost

import "testing"

func TestScalingPacketSummarizesHeadroomAndEvidence(t *testing.T) {
	packet := buildScalingPacket(Brief{
		PolicyPath:                      PolicyRelativePath,
		PublicSafe:                      true,
		Decision:                        "allow",
		Recommendation:                  "cache_value_supports_scaling",
		NextSafeAction:                  "continue_local_first_loop",
		BudgetState:                     "ok",
		TotalUnits:                      25000,
		WarningUnitThreshold:            100000,
		ReviewUnitThreshold:             500000,
		AttributionCoveragePercent:      100,
		AcceptedChangeCount:             12,
		CacheSavingsUnits:               400000,
		ValueProxyUnits:                 400012,
		CostPerAcceptedChange:           2000,
		SustainabilityPosture:           "sustainable",
		TrendPosture:                    "on_trend",
		StorageArchivePattern:           "compress_then_archive",
		StorageArchiveReady:             true,
		NoiseBudgetReady:                true,
		MaxNoiseRatioPercent:            20,
		ArchiveManifestCompressionRatio: 3,
		ConfigIsEvidence:                true,
	})
	if !packet.PublicSafe || packet.Context != "CodexCostScalingPacket" {
		t.Fatalf("packet = %#v", packet)
	}
	if packet.BudgetHeadroom.RemainingToWarning != 75000 ||
		packet.BudgetHeadroom.RemainingToReview != 475000 ||
		packet.BudgetHeadroom.WarningUsedPercent != 25 ||
		packet.BudgetHeadroom.ReviewUsedPercent != 5 {
		t.Fatalf("headroom = %#v", packet.BudgetHeadroom)
	}
	if packet.EvidencePosture.CacheSavingsUnits != 400000 ||
		!packet.StorageEvidence.ConfigIsEvidence {
		t.Fatalf("evidence = %#v %#v", packet.EvidencePosture, packet.StorageEvidence)
	}
}

func TestScalingPacketOptionsNeverGrantExpansion(t *testing.T) {
	packet := buildScalingPacket(Brief{Decision: "review_required"})
	if packet.CanApplyExpansion {
		t.Fatalf("packet can apply expansion = %#v", packet)
	}
	for _, option := range packet.ScalingOptions {
		if option.ThisPacketGrantsSpend || option.AllowsPaidExpansion ||
			option.AllowsExternalTooling || option.AllowsWorkflowChanges ||
			option.AllowsSelfApproval {
			t.Fatalf("granting option = %#v", option)
		}
	}
}
