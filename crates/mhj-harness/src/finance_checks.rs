use crate::finance_candidates::finance_candidate_checks;
use crate::finance_owner::finance_owner_checks;
use crate::finance_snapshot::FinanceSnapshot;
use crate::finance_totals::finance_total_checks;
use crate::HarnessResult;

pub(crate) fn finance_checks(snapshot: &FinanceSnapshot) -> Vec<HarnessResult> {
    let mut results = Vec::new();
    results.extend(finance_total_checks(snapshot));
    results.extend(finance_candidate_checks(snapshot));
    results.extend(finance_owner_checks(snapshot));
    results
}
