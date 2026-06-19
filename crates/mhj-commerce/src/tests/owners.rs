use crate::*;

#[test]
fn owner_spend_keeps_household_and_personal_purchases_separate() {
    let purchases = fixture_purchases().expect("commerce fixture parses");
    let summaries = summarize_by_owner(&purchases, "KRW").expect("owners summarize");
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
    assert_eq!(user.total_spend_minor_units, 3_200);
    assert_eq!(spouse.records, 0);
    assert_eq!(household.records, 2);
    assert_eq!(household.total_spend_minor_units, 23_600);
}
