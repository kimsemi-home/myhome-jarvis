use crate::records::FinanceJsonRecord;
use crate::validate_common::{ensure_non_empty, ensure_owner, ensure_positive_amount};

pub(crate) fn validate_finance_record(
    record: &FinanceJsonRecord,
) -> Result<(), (&'static str, String)> {
    ensure_non_empty("transaction_id", &record.transaction_id)?;
    ensure_non_empty("source", &record.source)?;
    ensure_owner(&record.owner)?;
    ensure_non_empty("occurred_at", &record.occurred_at)?;
    ensure_positive_amount("amount", &record.amount)?;
    ensure_direction(&record.direction)?;
    ensure_non_empty("raw_ref", &record.raw_ref)?;
    validate_debit_merchant(record)?;
    validate_account_or_card(record)
}

fn ensure_direction(direction: &str) -> Result<(), (&'static str, String)> {
    match direction {
        "debit" | "credit" => Ok(()),
        _ => Err(("direction", "must be debit or credit".to_string())),
    }
}

fn validate_debit_merchant(record: &FinanceJsonRecord) -> Result<(), (&'static str, String)> {
    if record.direction == "debit"
        && record
            .merchant_name
            .as_deref()
            .unwrap_or_default()
            .trim()
            .is_empty()
    {
        return Err(("merchant_name", "required for debit records".to_string()));
    }
    Ok(())
}

fn validate_account_or_card(record: &FinanceJsonRecord) -> Result<(), (&'static str, String)> {
    let account = record.account_id.as_deref().unwrap_or_default().trim();
    let card = record.card_id.as_deref().unwrap_or_default().trim();
    if account.is_empty() && card.is_empty() {
        return Err((
            "account_id",
            "account_id or card_id is required".to_string(),
        ));
    }
    Ok(())
}
