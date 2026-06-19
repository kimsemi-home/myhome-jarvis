use crate::commerce_candidates::commerce_candidate_checks;
use crate::commerce_owner::commerce_owner_checks;
use crate::commerce_snapshot::CommerceSnapshot;
use crate::commerce_totals::commerce_total_checks;
use crate::HarnessResult;

pub(crate) fn commerce_checks(snapshot: &CommerceSnapshot) -> Vec<HarnessResult> {
    let mut results = Vec::new();
    results.extend(commerce_total_checks(snapshot));
    results.extend(commerce_candidate_checks(snapshot));
    results.extend(commerce_owner_checks(snapshot));
    results
}
