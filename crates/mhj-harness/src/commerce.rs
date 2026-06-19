use crate::commerce_checks::commerce_checks;
use crate::commerce_snapshot::commerce_snapshot;
use crate::{HarnessReport, HarnessResult};

pub fn run_commerce_harness_report() -> HarnessReport {
    HarnessReport::new("commerce", run_commerce_harness())
}

pub fn run_commerce_harness() -> Vec<HarnessResult> {
    match commerce_snapshot() {
        Ok(snapshot) => commerce_checks(&snapshot),
        Err(result) => vec![result],
    }
}
