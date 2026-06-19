use crate::{fixture_transactions, summarize_cashflow};

#[test]
fn cashflow_summary_uses_direction_not_signed_amounts() {
    let transactions = fixture_transactions().expect("finance fixture parses");
    let summary = summarize_cashflow(&transactions, "KRW").expect("cashflow summarizes");
    assert_eq!(summary.records, 3);
    assert_eq!(summary.credit_minor_units, 4_500_000);
    assert_eq!(summary.debit_minor_units, 153_200);
    assert_eq!(summary.net_minor_units, 4_346_800);
    assert_eq!(summary.subscription_minor_units, 65_900);
    assert_eq!(summary.subscription_count, 1);
}
