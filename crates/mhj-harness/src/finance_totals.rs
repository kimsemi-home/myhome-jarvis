use crate::check::check;
use crate::finance_snapshot::FinanceSnapshot;
use crate::HarnessResult;

pub(crate) fn finance_total_checks(snapshot: &FinanceSnapshot) -> Vec<HarnessResult> {
    let summary = &snapshot.summary;
    vec![
        check(
            "finance fixture records",
            snapshot.transaction_count == 3 && summary.records == 3,
            format!(
                "records={} summary={}",
                snapshot.transaction_count, summary.records
            ),
        ),
        check(
            "finance currency",
            summary.currency == "KRW",
            format!("currency={}", summary.currency),
        ),
        check(
            "finance credit total",
            summary.credit_minor_units == 4_500_000,
            format!("credit={}", summary.credit_minor_units),
        ),
        check(
            "finance debit total",
            summary.debit_minor_units == 153_200,
            format!("debit={}", summary.debit_minor_units),
        ),
        check(
            "finance net total",
            summary.net_minor_units == 4_346_800,
            format!("net={}", summary.net_minor_units),
        ),
    ]
}
