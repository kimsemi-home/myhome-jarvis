use crate::*;

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

#[test]
fn unknown_owners_are_rejected() {
    let mut purchase = fixture_purchases()
        .expect("commerce fixture parses")
        .into_iter()
        .next()
        .expect("fixture has purchase");
    purchase.owner = Owner::Unknown;
    let error = purchase.validate().expect_err("unknown owner fails");
    assert_eq!(error.field, "owner");
}
