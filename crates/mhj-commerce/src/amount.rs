use crate::{ensure_currency, ValidationError};
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct MoneyAmount {
    pub minor_units: i64,
    pub currency: String,
}

impl MoneyAmount {
    pub fn validate_non_negative(&self, field: &'static str) -> Result<(), ValidationError> {
        ensure_currency(field, &self.currency)?;
        if self.minor_units < 0 {
            return Err(ValidationError::new(field, "amount must not be negative"));
        }
        Ok(())
    }
}
