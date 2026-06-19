use crate::run_finance_harness_report;

#[test]
fn finance_harness_passes() {
    let report = run_finance_harness_report();
    assert_eq!(report.name, "finance");
    assert!(report.passed, "{:#?}", report.results);
    assert!(report.results.len() >= 9);
}
