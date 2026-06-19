use crate::{
    ensure_iso_datetime, ensure_known_owner, ensure_non_empty, ensure_tags, ValidationError,
};

use super::{TransactionDirection, TransactionIr};

impl TransactionIr {
    pub fn validate(&self) -> Result<(), ValidationError> {
        ensure_non_empty("transaction_id", &self.transaction_id)?;
        ensure_non_empty("source", &self.source)?;
        ensure_known_owner("owner", self.owner)?;
        ensure_iso_datetime("occurred_at", &self.occurred_at)?;
        if let Some(posted_at) = &self.posted_at {
            ensure_iso_datetime("posted_at", posted_at)?;
        }
        self.amount.validate_positive("amount")?;
        ensure_has_payment_anchor(self)?;
        ensure_debit_has_merchant(self)?;
        ensure_non_empty("raw_ref", &self.raw_ref)?;
        ensure_tags("tags", &self.tags)?;
        Ok(())
    }
}

fn ensure_has_payment_anchor(transaction: &TransactionIr) -> Result<(), ValidationError> {
    let account_missing = transaction
        .account_id
        .as_deref()
        .unwrap_or_default()
        .trim()
        .is_empty();
    let card_missing = transaction
        .card_id
        .as_deref()
        .unwrap_or_default()
        .trim()
        .is_empty();
    if account_missing && card_missing {
        return Err(ValidationError::new(
            "account_id",
            "account_id or card_id is required",
        ));
    }
    Ok(())
}

fn ensure_debit_has_merchant(transaction: &TransactionIr) -> Result<(), ValidationError> {
    if !matches!(transaction.direction, TransactionDirection::Debit) {
        return Ok(());
    }
    let merchant_missing = transaction
        .merchant_name
        .as_deref()
        .unwrap_or_default()
        .trim()
        .is_empty();
    if merchant_missing {
        return Err(ValidationError::new(
            "merchant_name",
            "merchant_name is required for debit transactions",
        ));
    }
    Ok(())
}
