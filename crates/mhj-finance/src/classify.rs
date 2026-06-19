use crate::{TransactionDirection, TransactionIr};

pub(crate) fn is_subscription(transaction: &TransactionIr) -> bool {
    matches!(transaction.direction, TransactionDirection::Debit)
        && transaction
            .category
            .as_deref()
            .unwrap_or_default()
            .eq_ignore_ascii_case("subscription")
}
