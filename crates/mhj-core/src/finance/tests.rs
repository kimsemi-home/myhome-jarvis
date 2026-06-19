use super::*;

#[test]
fn finance_fixture_ir_is_valid_jsonl() {
    let transactions = fixture_transactions().expect("finance fixture parses");
    assert_eq!(transactions.len(), 3);
    assert!(transactions
        .iter()
        .all(|transaction| transaction.validate().is_ok()));
}

#[test]
fn cashflow_summary_uses_direction_not_signed_amounts() {
    let transactions = fixture_transactions().expect("finance fixture parses");
    let summary = summarize_cashflow(&transactions, "KRW").expect("cashflow summarizes");
    assert_eq!(summary.credit_minor_units, 4_500_000);
    assert_eq!(summary.debit_minor_units, 153_200);
    assert_eq!(summary.net_minor_units, 4_346_800);
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

#[test]
fn debit_transactions_require_merchants() {
    let mut transaction = fixture_transactions()
        .expect("finance fixture parses")
        .into_iter()
        .find(|transaction| matches!(transaction.direction, TransactionDirection::Debit))
        .expect("fixture includes debit");
    transaction.merchant_name = None;
    let error = transaction.validate().expect_err("missing merchant fails");
    assert_eq!(error.field, "merchant_name");
}
