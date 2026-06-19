use crate::*;

#[test]
fn commerce_fixture_ir_is_valid_jsonl() {
    let purchases = fixture_purchases().expect("commerce fixture parses");
    assert_eq!(purchases.len(), 3);
    assert!(purchases.iter().all(|purchase| purchase.validate().is_ok()));
}
