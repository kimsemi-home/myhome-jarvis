use crate::*;

#[test]
fn commerce_summary_counts_fixture_spend_and_recurring_flags() {
    let purchases = fixture_purchases().expect("commerce fixture parses");
    let summary = summarize_commerce(&purchases, "KRW").expect("commerce summarizes");

    assert_eq!(summary.records, 3);
    assert_eq!(summary.total_spend_minor_units, 26_800);
    assert_eq!(summary.recurring_candidate_count, 2);
}
