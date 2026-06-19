use crate::{TransactionDirection, TransactionIr};

impl TransactionIr {
    pub fn signed_minor_units(&self) -> i64 {
        match self.direction {
            TransactionDirection::Debit => -self.amount.minor_units,
            TransactionDirection::Credit => self.amount.minor_units,
        }
    }
}
