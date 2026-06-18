use mhj_core::benchmark::run_benchmark_smoke;
use std::time::Duration;

const SMOKE_ITERATIONS: usize = 1_000;

#[test]
fn benchmark_smoke_runs_core_fixture_pipeline() {
    let report = run_benchmark_smoke(SMOKE_ITERATIONS).expect("benchmark smoke runs");
    assert_eq!(report.finance_records, SMOKE_ITERATIONS * 3);
    assert_eq!(report.commerce_records, SMOKE_ITERATIONS * 3);
    assert_eq!(report.recurring_candidates, SMOKE_ITERATIONS);
    assert_eq!(report.storage_plans, SMOKE_ITERATIONS * 8);
    report.assert_reasonable(Duration::from_secs(2));
}
