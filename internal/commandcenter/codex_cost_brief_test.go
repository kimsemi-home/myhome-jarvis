package commandcenter

import "testing"

func TestStatusIncludesCodexCostBrief(t *testing.T) {
	status, err := StatusForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	brief := status.CodexCostBrief
	if !brief.PublicSafe ||
		brief.Decision == "" ||
		brief.NextSafeAction == "" ||
		brief.Recommendation == "" {
		t.Fatalf("codex cost brief = %#v", brief)
	}
	if brief.ValueProxyUnits < brief.AcceptedChangeCount ||
		brief.ForbiddenPublicFieldCount == 0 ||
		brief.TotalUnits < 0 {
		t.Fatalf("codex cost value evidence = %#v", brief)
	}
	if !brief.StorageArchiveReady ||
		!brief.NoiseBudgetReady ||
		brief.ArchiveManifestBudgetBreaches != 0 ||
		brief.ArchiveManifestInvalidEntries != 0 {
		t.Fatalf("codex cost archive evidence = %#v", brief)
	}
	if !hasPillar(status.Vision.ReadyPillarKeys, "codex_cost_governor") {
		if brief.Decision == "allow" {
			t.Fatalf("vision readiness did not use brief = %#v", status.Vision)
		}
	}
}
