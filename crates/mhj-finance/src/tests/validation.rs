use crate::{fixture_transactions, Owner, TransactionDirection};

#[test]
fn debit_transactions_require_merchants() {
    let mut transaction = fixture_transactions()
        .expect("finance fixture parses")
        .into_iter()
        .find(|transaction| matches!(transaction.direction, TransactionDirection::Debit))
        .expect("fixture includes debit");
    transaction.merchant_name = None;
    let error = transaction.validate().expect_err("missing merchant fails");
    assert_eq!(error.field, "merchant_name");
}

#[test]
fn unknown_owners_are_rejected() {
    let mut transaction = fixture_transactions()
        .expect("finance fixture parses")
        .into_iter()
        .next()
        .expect("fixture includes transactions");
    transaction.owner = Owner::Unknown;
    let error = transaction.validate().expect_err("unknown owner fails");
    assert_eq!(error.field, "owner");
}
