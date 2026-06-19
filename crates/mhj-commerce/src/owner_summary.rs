use crate::{ensure_currency, Owner, OwnerSpendSummary, PurchaseIr, ValidationError};

pub fn summarize_by_owner(
    purchases: &[PurchaseIr],
    currency: &str,
) -> Result<Vec<OwnerSpendSummary>, ValidationError> {
    ensure_currency("currency", currency)?;
    let mut summaries = [
        owner_summary(Owner::User, currency),
        owner_summary(Owner::Spouse, currency),
        owner_summary(Owner::Household, currency),
    ];
    for purchase in purchases {
        purchase.validate()?;
        if purchase.total_price.currency != currency {
            continue;
        }
        let Some(summary) = summaries
            .iter_mut()
            .find(|summary| summary.owner == purchase.owner)
        else {
            continue;
        };
        summary.records += 1;
        summary.total_spend_minor_units += purchase.total_price.minor_units;
    }
    Ok(summaries.into_iter().collect())
}

fn owner_summary(owner: Owner, currency: &str) -> OwnerSpendSummary {
    OwnerSpendSummary {
        owner,
        records: 0,
        currency: currency.to_string(),
        total_spend_minor_units: 0,
    }
}
