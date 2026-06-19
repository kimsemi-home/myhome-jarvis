use crate::*;

#[test]
fn recurring_candidates_require_repeated_flagged_purchases() {
    let purchases = fixture_purchases().expect("commerce fixture parses");
    let candidates = recurring_candidates(&purchases).expect("recurring summarizes");

    assert_eq!(candidates.len(), 1);
    assert_eq!(candidates[0].merchant_name, "Coupang");
    assert_eq!(candidates[0].item_name, "Bottled water 2L x 6");
    assert_eq!(candidates[0].purchase_count, 2);
    assert_eq!(candidates[0].latest_total_minor_units, 11_800);
    assert_eq!(
        candidates[0].latest_purchased_at,
        "2026-06-11T11:15:00+09:00"
    );
}
