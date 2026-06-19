use crate::datetime::ensure_iso_datetime;
use crate::validation::{ensure_known_owner, ensure_non_empty, ensure_tags};
use crate::{TransactionDirection, TransactionIr, ValidationError};

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
    if !blank_optional(&transaction.account_id) || !blank_optional(&transaction.card_id) {
        return Ok(());
    }
    Err(ValidationError::new(
        "account_id",
        "account_id or card_id is required",
    ))
}

fn ensure_debit_has_merchant(transaction: &TransactionIr) -> Result<(), ValidationError> {
    if !matches!(transaction.direction, TransactionDirection::Debit) {
        return Ok(());
    }
    if !blank_optional(&transaction.merchant_name) {
        return Ok(());
    }
    Err(ValidationError::new(
        "merchant_name",
        "merchant_name is required for debit transactions",
    ))
}

fn blank_optional(value: &Option<String>) -> bool {
    value.as_deref().unwrap_or_default().trim().is_empty()
}
