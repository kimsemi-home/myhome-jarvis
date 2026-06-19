use crate::{fixture_transactions, summarize_by_owner, Owner};

#[test]
fn owner_cashflow_keeps_household_and_personal_views_separate() {
    let transactions = fixture_transactions().expect("finance fixture parses");
    let summaries = summarize_by_owner(&transactions, "KRW").expect("owners summarize");
    let user = summaries
        .iter()
        .find(|summary| summary.owner == Owner::User)
        .expect("user summary exists");
    let spouse = summaries
        .iter()
        .find(|summary| summary.owner == Owner::Spouse)
        .expect("spouse summary exists");
    let household = summaries
        .iter()
        .find(|summary| summary.owner == Owner::Household)
        .expect("household summary exists");

    assert_eq!(user.records, 1);
    assert_eq!(user.net_minor_units, -87_300);
    assert_eq!(spouse.records, 0);
    assert_eq!(household.records, 2);
    assert_eq!(household.net_minor_units, 4_434_100);
}
