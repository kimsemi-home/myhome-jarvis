use crate::{home_cases, run_home_harness, run_home_harness_report};

#[test]
fn home_harness_passes() {
    let results = run_home_harness();
    assert!(results.iter().all(|result| result.passed), "{results:#?}");
}

#[test]
fn home_harness_report_passes() {
    let report = run_home_harness_report();
    assert_eq!(report.name, "home");
    assert!(report.passed, "{:#?}", report.results);
    assert_eq!(report.results.len(), home_cases().len());
}
