use super::*;

#[test]
fn commerce_fixture_ir_is_valid_jsonl() {
    let purchases = fixture_purchases().expect("commerce fixture parses");

    assert_eq!(purchases.len(), 3);
    assert!(purchases.iter().all(|purchase| purchase.validate().is_ok()));
}

#[test]
fn recurring_candidates_require_repeated_flagged_purchases() {
    let purchases = fixture_purchases().expect("commerce fixture parses");
    let candidates = recurring_candidates(&purchases);

    assert_eq!(candidates.len(), 1);
    assert_eq!(candidates[0].merchant_name, "Coupang");
    assert_eq!(candidates[0].item_name, "Bottled water 2L x 6");
    assert_eq!(candidates[0].purchase_count, 2);
}

#[test]
fn purchase_totals_must_match_quantity() {
    let mut purchase = fixture_purchases()
        .expect("commerce fixture parses")
        .into_iter()
        .next()
        .expect("fixture has purchase");

    purchase.total_price.minor_units += 1;

    let error = purchase.validate().expect_err("bad total fails");
    assert_eq!(error.field, "total_price");
}
