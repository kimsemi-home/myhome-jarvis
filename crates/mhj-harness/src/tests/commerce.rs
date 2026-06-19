use crate::run_commerce_harness_report;

#[test]
fn commerce_harness_passes() {
    let report = run_commerce_harness_report();
    assert_eq!(report.name, "commerce");
    assert!(report.passed, "{:#?}", report.results);
    assert!(report.results.len() >= 8);
}
