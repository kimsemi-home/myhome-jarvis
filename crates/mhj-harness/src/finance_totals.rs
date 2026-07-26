use crate::check::check;
use crate::finance_snapshot::FinanceSnapshot;
use crate::HarnessResult;

pub(crate) fn finance_total_checks(snapshot: &FinanceSnapshot) -> Vec<HarnessResult> {
    let summary = &snapshot.summary;
    vec![
        check(
            "finance fixture records",
            snapshot.transaction_count == 4 && summary.records == 4,
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
            summary.debit_minor_units == 195_200,
            format!("debit={}", summary.debit_minor_units),
        ),
        check(
            "finance net total",
            summary.net_minor_units == 4_304_800,
            format!("net={}", summary.net_minor_units),
        ),
    ]
}
