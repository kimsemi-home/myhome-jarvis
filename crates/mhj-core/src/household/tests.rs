use super::*;

#[test]
fn fixture_household_scopes_include_user_spouse_and_household() {
    let scopes = fixture_household_scopes().expect("household scopes build");
    assert_eq!(scopes.len(), 3);

    let user = scopes
        .iter()
        .find(|scope| scope.scope == "user")
        .expect("user scope exists");
    assert_eq!(user.finance_records, 1);
    assert_eq!(user.finance_net_minor_units, -87_300);
    assert_eq!(user.purchase_records, 1);
    assert_eq!(user.purchase_spend_minor_units, 3_200);

    let spouse = scopes
        .iter()
        .find(|scope| scope.scope == "spouse")
        .expect("spouse scope exists");
    assert_eq!(spouse.finance_records, 0);
    assert_eq!(spouse.purchase_records, 0);

    let household = scopes
        .iter()
        .find(|scope| scope.scope == "household")
        .expect("household scope exists");
    assert_eq!(household.finance_records, 3);
    assert_eq!(household.finance_net_minor_units, 4_346_800);
    assert_eq!(household.purchase_records, 3);
    assert_eq!(household.purchase_spend_minor_units, 26_800);
}
