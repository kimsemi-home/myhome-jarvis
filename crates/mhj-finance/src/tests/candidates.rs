use crate::{card_usage_candidates, fixture_transactions, subscription_candidates, Owner};

#[test]
fn subscription_candidates_are_review_only() {
    let transactions = fixture_transactions().expect("finance fixture parses");
    let candidates = subscription_candidates(&transactions).expect("subscriptions summarize");
    assert_eq!(candidates.len(), 1);
    assert_eq!(candidates[0].merchant_name, "Streaming Bundle");
    assert_eq!(candidates[0].monthly_minor_units, 65_900);
    assert_eq!(candidates[0].owner, Owner::Household);
}

#[test]
fn card_usage_candidates_are_review_only() {
    let transactions = fixture_transactions().expect("finance fixture parses");
    let candidates = card_usage_candidates(&transactions, "KRW").expect("cards summarize");
    assert_eq!(candidates.len(), 1);
    assert_eq!(candidates[0].currency, "KRW");
    assert_eq!(candidates[0].transaction_count, 2);
    assert_eq!(candidates[0].debit_minor_units, 153_200);
    assert_eq!(candidates[0].subscription_count, 1);
    assert_eq!(candidates[0].subscription_minor_units, 65_900);
}
