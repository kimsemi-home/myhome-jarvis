use crate::records::JsonMoney;

pub(crate) fn ensure_non_empty(
    field: &'static str,
    value: &str,
) -> Result<(), (&'static str, String)> {
    if value.trim().is_empty() {
        return Err((field, "must not be empty".to_string()));
    }
    Ok(())
}

pub(crate) fn ensure_owner(owner: &str) -> Result<(), (&'static str, String)> {
    match owner {
        "user" | "spouse" | "household" => Ok(()),
        _ => Err(("owner", "must be user, spouse, or household".to_string())),
    }
}

pub(crate) fn ensure_positive_amount(
    field: &'static str,
    amount: &JsonMoney,
) -> Result<(), (&'static str, String)> {
    if amount.currency.trim().is_empty() {
        return Err((field, "currency must not be empty".to_string()));
    }
    if amount.minor_units <= 0 {
        return Err((field, "minor units must be greater than zero".to_string()));
    }
    Ok(())
}
