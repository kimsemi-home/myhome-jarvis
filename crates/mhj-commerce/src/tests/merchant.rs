use crate::*;

#[test]
fn merchant_spend_is_sorted_by_total_spend() {
    let purchases = fixture_purchases().expect("commerce fixture parses");
    let merchants = summarize_by_merchant(&purchases, "KRW").expect("merchants summarize");

    assert_eq!(merchants.len(), 2);
    assert_eq!(merchants[0].merchant_name, "Coupang");
    assert_eq!(merchants[0].total_spend_minor_units, 23_600);
    assert_eq!(merchants[1].merchant_name, "Local Mart");
}
