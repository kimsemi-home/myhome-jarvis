use crate::ValidationError;

use super::{TransactionDirection, TransactionIr};

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct CashflowSummary {
    pub currency: String,
    pub credit_minor_units: i64,
    pub debit_minor_units: i64,
    pub net_minor_units: i64,
}

pub fn summarize_cashflow(
    transactions: &[TransactionIr],
    currency: &str,
) -> Result<CashflowSummary, ValidationError> {
    let mut credit_minor_units = 0;
    let mut debit_minor_units = 0;
    for transaction in transactions {
        transaction.validate()?;
        if transaction.amount.currency != currency {
            continue;
        }
        match transaction.direction {
            TransactionDirection::Debit => debit_minor_units += transaction.amount.minor_units,
            TransactionDirection::Credit => credit_minor_units += transaction.amount.minor_units,
        }
    }
    Ok(CashflowSummary {
        currency: currency.to_string(),
        credit_minor_units,
        debit_minor_units,
        net_minor_units: credit_minor_units - debit_minor_units,
    })
}
