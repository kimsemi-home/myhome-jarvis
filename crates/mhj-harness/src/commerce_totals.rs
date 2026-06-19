use crate::check::check;
use crate::commerce_snapshot::CommerceSnapshot;
use crate::HarnessResult;

pub(crate) fn commerce_total_checks(snapshot: &CommerceSnapshot) -> Vec<HarnessResult> {
    let summary = &snapshot.summary;
    vec![
        check(
            "commerce fixture records",
            snapshot.purchase_count == 3 && summary.records == 3,
            format!(
                "records={} summary={}",
                snapshot.purchase_count, summary.records
            ),
        ),
        check(
            "commerce currency",
            summary.currency == "KRW",
            format!("currency={}", summary.currency),
        ),
        check(
            "commerce spend total",
            summary.total_spend_minor_units == 26_800,
            format!("spend={}", summary.total_spend_minor_units),
        ),
        check(
            "commerce recurring fixture rows",
            summary.recurring_candidate_count == 2,
            format!("recurring_rows={}", summary.recurring_candidate_count),
        ),
    ]
}
