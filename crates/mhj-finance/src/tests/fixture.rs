use crate::fixture_transactions;

#[test]
fn finance_fixture_ir_is_valid_jsonl() {
    let transactions = fixture_transactions().expect("finance fixture parses");
    assert_eq!(transactions.len(), 3);
    assert!(transactions
        .iter()
        .all(|transaction| transaction.validate().is_ok()));
}
