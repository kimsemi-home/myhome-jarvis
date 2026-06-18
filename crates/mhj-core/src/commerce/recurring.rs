use super::PurchaseIr;
use std::collections::BTreeMap;

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct RecurringPurchaseCandidate {
    pub merchant_name: String,
    pub item_name: String,
    pub currency: String,
    pub purchase_count: usize,
    pub latest_total_minor_units: i64,
}

pub fn recurring_candidates(purchases: &[PurchaseIr]) -> Vec<RecurringPurchaseCandidate> {
    let mut grouped: BTreeMap<(String, String, String), Vec<&PurchaseIr>> = BTreeMap::new();
    for purchase in purchases {
        if purchase.validate().is_err() || !purchase.recurring_candidate {
            continue;
        }
        grouped
            .entry((
                purchase.merchant_name.clone(),
                purchase.item_name.clone(),
                purchase.total_price.currency.clone(),
            ))
            .or_default()
            .push(purchase);
    }

    grouped
        .into_iter()
        .filter_map(|((merchant_name, item_name, currency), records)| {
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
            })
        })
        .collect()
}
