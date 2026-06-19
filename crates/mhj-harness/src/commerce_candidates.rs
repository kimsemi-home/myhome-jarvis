use crate::check::check;
use crate::commerce_snapshot::CommerceSnapshot;
use crate::HarnessResult;

pub(crate) fn commerce_candidate_checks(snapshot: &CommerceSnapshot) -> Vec<HarnessResult> {
    vec![
        check(
            "commerce top merchant",
            top_merchant_ok(snapshot),
            top_merchant_message(snapshot),
        ),
        check(
            "recurring purchase candidate detail",
            recurring_candidate_ok(snapshot),
            recurring_candidate_message(snapshot),
        ),
    ]
}

fn top_merchant_ok(snapshot: &CommerceSnapshot) -> bool {
    snapshot.merchants.first().is_some_and(|merchant| {
        merchant.merchant_name == "Coupang"
            && merchant.records == 2
            && merchant.total_spend_minor_units == 23_600
    })
}

fn top_merchant_message(snapshot: &CommerceSnapshot) -> String {
    snapshot
        .merchants
        .first()
        .map(|merchant| {
            format!(
                "{} records={} spend={}",
                merchant.merchant_name, merchant.records, merchant.total_spend_minor_units
            )
        })
        .unwrap_or_else(|| "missing merchant".to_string())
}

fn recurring_candidate_ok(snapshot: &CommerceSnapshot) -> bool {
    snapshot.recurring.first().is_some_and(|candidate| {
        candidate.merchant_name == "Coupang"
            && candidate.item_name == "Bottled water 2L x 6"
            && candidate.purchase_count == 2
            && candidate.latest_total_minor_units == 11_800
    })
}

fn recurring_candidate_message(snapshot: &CommerceSnapshot) -> String {
    snapshot
        .recurring
        .first()
        .map(|candidate| {
            format!(
                "{} {} count={} total={}",
                candidate.merchant_name,
                candidate.item_name,
                candidate.purchase_count,
                candidate.latest_total_minor_units
            )
        })
        .unwrap_or_else(|| "missing recurring candidate".to_string())
}
