use crate::{ensure_currency, CommerceSummary, PurchaseIr, ValidationError};

pub fn summarize_commerce(
    purchases: &[PurchaseIr],
    currency: &str,
) -> Result<CommerceSummary, ValidationError> {
    ensure_currency("currency", currency)?;
    let mut summary = CommerceSummary {
        records: 0,
        currency: currency.to_string(),
        total_spend_minor_units: 0,
        recurring_candidate_count: 0,
    };
    for purchase in purchases {
        purchase.validate()?;
        if purchase.total_price.currency != currency {
            continue;
        }
        summary.records += 1;
        summary.total_spend_minor_units += purchase.total_price.minor_units;
        if purchase.recurring_candidate {
            summary.recurring_candidate_count += 1;
        }
    }
    Ok(summary)
}
