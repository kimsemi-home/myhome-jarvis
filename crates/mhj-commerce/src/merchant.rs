use crate::{ensure_currency, MerchantSpendSummary, PurchaseIr, ValidationError};
use std::collections::BTreeMap;

pub fn summarize_by_merchant(
    purchases: &[PurchaseIr],
    currency: &str,
) -> Result<Vec<MerchantSpendSummary>, ValidationError> {
    ensure_currency("currency", currency)?;
    let mut merchants: BTreeMap<String, MerchantSpendSummary> = BTreeMap::new();
    for purchase in purchases {
        purchase.validate()?;
        if purchase.total_price.currency != currency {
            continue;
        }
        let summary = merchants
            .entry(purchase.merchant_name.clone())
            .or_insert_with(|| MerchantSpendSummary {
                merchant_name: purchase.merchant_name.clone(),
                records: 0,
                currency: currency.to_string(),
                total_spend_minor_units: 0,
            });
        summary.records += 1;
        summary.total_spend_minor_units += purchase.total_price.minor_units;
    }
    let mut summaries: Vec<MerchantSpendSummary> = merchants.into_values().collect();
    sort_merchant_summaries(&mut summaries);
    Ok(summaries)
}

fn sort_merchant_summaries(summaries: &mut [MerchantSpendSummary]) {
    summaries.sort_by(|left, right| {
        right
            .total_spend_minor_units
            .cmp(&left.total_spend_minor_units)
            .then_with(|| left.merchant_name.cmp(&right.merchant_name))
    });
}
