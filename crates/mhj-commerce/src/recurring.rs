use crate::{PurchaseIr, RecurringPurchaseCandidate, ValidationError};
use std::collections::BTreeMap;

type GroupKey = (String, String, String);
type RecurringGroups<'a> = BTreeMap<GroupKey, Vec<&'a PurchaseIr>>;

pub fn recurring_candidates(
    purchases: &[PurchaseIr],
) -> Result<Vec<RecurringPurchaseCandidate>, ValidationError> {
    let mut candidates: Vec<RecurringPurchaseCandidate> = grouped_recurring(purchases)?
        .into_iter()
        .filter_map(|(key, records)| recurring_candidate(key, records))
        .collect();
    sort_recurring_candidates(&mut candidates);
    Ok(candidates)
}

fn grouped_recurring(purchases: &[PurchaseIr]) -> Result<RecurringGroups<'_>, ValidationError> {
    let mut grouped = BTreeMap::new();
    for purchase in purchases {
        purchase.validate()?;
        if !purchase.recurring_candidate {
            continue;
        }
        grouped
            .entry((
                purchase.merchant_name.clone(),
                purchase.item_name.clone(),
                purchase.total_price.currency.clone(),
            ))
            .or_insert_with(Vec::new)
            .push(purchase);
    }
    Ok(grouped)
}

fn recurring_candidate(
    (merchant_name, item_name, currency): GroupKey,
    records: Vec<&PurchaseIr>,
) -> Option<RecurringPurchaseCandidate> {
    let latest = records
        .iter()
        .max_by_key(|purchase| purchase.purchased_at.as_str())?;
    if records.len() < 2 {
        return None;
    }
    Some(RecurringPurchaseCandidate {
        merchant_name,
        item_name,
        currency,
        purchase_count: records.len(),
        latest_total_minor_units: latest.total_price.minor_units,
        latest_purchased_at: latest.purchased_at.clone(),
    })
}

fn sort_recurring_candidates(candidates: &mut [RecurringPurchaseCandidate]) {
    candidates.sort_by(|left, right| {
        right
            .purchase_count
            .cmp(&left.purchase_count)
            .then_with(|| left.merchant_name.cmp(&right.merchant_name))
            .then_with(|| left.item_name.cmp(&right.item_name))
    });
}
