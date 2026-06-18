use crate::commerce::{fixture_purchases, recurring_candidates};
use crate::finance::{fixture_transactions, summarize_cashflow};
use crate::storage::{plan_dataset, DatasetKind, LakeLayer};
use crate::FixtureError;
use std::path::Path;
use std::time::{Duration, Instant};

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct BenchmarkSmokeReport {
    pub iterations: usize,
    pub finance_records: usize,
    pub commerce_records: usize,
    pub recurring_candidates: usize,
    pub storage_plans: usize,
    pub elapsed: Duration,
}

impl BenchmarkSmokeReport {
    pub fn assert_reasonable(&self, ceiling: Duration) {
        assert!(
            self.elapsed <= ceiling,
            "benchmark smoke exceeded {ceiling:?}: {self:?}"
        );
    }
}

pub fn run_benchmark_smoke(iterations: usize) -> Result<BenchmarkSmokeReport, FixtureError> {
    let started = Instant::now();
    let mut finance_records = 0;
    let mut commerce_records = 0;
    let mut recurring_candidate_count = 0;
    let mut storage_plans = 0;

    for _ in 0..iterations {
        let transactions = fixture_transactions()?;
        let _summary =
            summarize_cashflow(&transactions, "KRW").map_err(|error| FixtureError::Validation {
                line: 0,
                field: error.field,
                message: error.message,
            })?;
        finance_records += transactions.len();

        let purchases = fixture_purchases()?;
        let candidates = recurring_candidates(&purchases);
        commerce_records += purchases.len();
        recurring_candidate_count += candidates.len();

        let lake_root = Path::new("data/lake");
        for layer in [
            LakeLayer::Raw,
            LakeLayer::Bronze,
            LakeLayer::Silver,
            LakeLayer::Gold,
        ] {
            let _finance_plan = plan_dataset(lake_root, layer, DatasetKind::FinanceTransactions);
            let _commerce_plan = plan_dataset(lake_root, layer, DatasetKind::CommercePurchases);
            storage_plans += 2;
        }
    }

    Ok(BenchmarkSmokeReport {
        iterations,
        finance_records,
        commerce_records,
        recurring_candidates: recurring_candidate_count,
        storage_plans,
        elapsed: started.elapsed(),
    })
}
