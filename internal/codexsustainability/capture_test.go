package codexsustainability

import "testing"

func TestCaptureQualityRunSeedsFreshTrendEvidence(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeQualityRun(t, root, true)

	capture, err := captureQualityRunAt(root, mustTime(t, "2026-06-20T01:00:00Z"))
	if err != nil {
		t.Fatal(err)
	}
	if capture.CaptureState != "recorded" || capture.RecordedRecordCount != 2 {
		t.Fatalf("capture = %#v", capture)
	}
	status, err := statusForRootAt(root, mustTime(t, "2026-06-20T01:00:00Z"))
	if err != nil {
		t.Fatal(err)
	}
	if status.TrendPosture != "on_trend" ||
		status.EvidenceFreshness != "fresh" ||
		status.SustainabilityPosture != "sustainable" {
		t.Fatalf("status = %#v", status)
	}
}

func TestCaptureQualityRunIsIdempotent(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeQualityRun(t, root, true)
	if _, err := captureQualityRunAt(root, mustTime(t, "2026-06-20T01:00:00Z")); err != nil {
		t.Fatal(err)
	}
	capture, err := captureQualityRunAt(root, mustTime(t, "2026-06-20T01:05:00Z"))
	if err != nil {
		t.Fatal(err)
	}
	if capture.CaptureState != "already_recorded" || capture.RecordedRecordCount != 0 {
		t.Fatalf("capture = %#v", capture)
	}
}

func TestCaptureQualityRunHandlesMissingOrFailedQuality(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	missing, err := captureQualityRunAt(root, mustTime(t, "2026-06-20T01:00:00Z"))
	if err != nil {
		t.Fatal(err)
	}
	if missing.CaptureState != "missing_quality" {
		t.Fatalf("missing = %#v", missing)
	}
	writeQualityRun(t, root, false)
	failed, err := captureQualityRunAt(root, mustTime(t, "2026-06-20T01:00:00Z"))
	if err != nil {
		t.Fatal(err)
	}
	if failed.CaptureState != "last_run_not_successful" {
		t.Fatalf("failed = %#v", failed)
	}
}
