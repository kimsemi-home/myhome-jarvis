use crate::classify::is_subscription;
use crate::validation::ensure_currency;
use crate::{CashflowSummary, TransactionDirection, TransactionIr, ValidationError};

pub fn summarize_cashflow(
    transactions: &[TransactionIr],
    currency: &str,
) -> Result<CashflowSummary, ValidationError> {
    ensure_currency("currency", currency)?;
    let mut summary = CashflowSummary {
        records: 0,
        currency: currency.to_string(),
        credit_minor_units: 0,
        debit_minor_units: 0,
        net_minor_units: 0,
        subscription_minor_units: 0,
        subscription_count: 0,
    };
    for transaction in transactions {
        transaction.validate()?;
        if transaction.amount.currency != currency {
            continue;
        }
        summary.records += 1;
        apply_transaction(&mut summary, transaction);
    }
    summary.net_minor_units = summary.credit_minor_units - summary.debit_minor_units;
    Ok(summary)
}

fn apply_transaction(summary: &mut CashflowSummary, transaction: &TransactionIr) {
    match transaction.direction {
        TransactionDirection::Debit => {
            summary.debit_minor_units += transaction.amount.minor_units;
            if is_subscription(transaction) {
                summary.subscription_minor_units += transaction.amount.minor_units;
                summary.subscription_count += 1;
            }
        }
        TransactionDirection::Credit => {
            summary.credit_minor_units += transaction.amount.minor_units;
        }
    }
}
