use crate::home_cases;
use crate::home_eval::run_home_case;
use crate::{HarnessReport, HarnessResult};

pub fn run_home_harness_report() -> HarnessReport {
    HarnessReport::new("home", run_home_harness())
}

pub fn run_home_harness() -> Vec<HarnessResult> {
    home_cases().into_iter().map(run_home_case).collect()
}
