use crate::summary::owner_summary;
use crate::validation::ensure_currency;
use crate::{Owner, OwnerCashflowSummary, TransactionDirection, TransactionIr, ValidationError};

pub fn summarize_by_owner(
    transactions: &[TransactionIr],
    currency: &str,
) -> Result<Vec<OwnerCashflowSummary>, ValidationError> {
    ensure_currency("currency", currency)?;
    let mut summaries = [
        owner_summary(Owner::User, currency),
        owner_summary(Owner::Spouse, currency),
        owner_summary(Owner::Household, currency),
    ];
    for transaction in transactions {
        transaction.validate()?;
        if transaction.amount.currency != currency {
            continue;
        }
        apply_owner_transaction(&mut summaries, transaction);
    }
    Ok(summaries.into_iter().collect())
}

fn apply_owner_transaction(summaries: &mut [OwnerCashflowSummary], transaction: &TransactionIr) {
    let Some(summary) = summaries
        .iter_mut()
        .find(|summary| summary.owner == transaction.owner)
    else {
        return;
    };
    summary.records += 1;
    match transaction.direction {
        TransactionDirection::Debit => summary.debit_minor_units += transaction.amount.minor_units,
        TransactionDirection::Credit => {
            summary.credit_minor_units += transaction.amount.minor_units
        }
    }
    summary.net_minor_units = summary.credit_minor_units - summary.debit_minor_units;
}
