use crate::finance_checks::finance_checks;
use crate::finance_snapshot::finance_snapshot;
use crate::{HarnessReport, HarnessResult};

pub fn run_finance_harness_report() -> HarnessReport {
    HarnessReport::new("finance", run_finance_harness())
}

pub fn run_finance_harness() -> Vec<HarnessResult> {
    match finance_snapshot() {
        Ok(snapshot) => finance_checks(&snapshot),
        Err(result) => vec![result],
    }
}
