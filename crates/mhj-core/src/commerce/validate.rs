use super::PurchaseIr;
use crate::{
    ensure_iso_datetime, ensure_known_owner, ensure_non_empty, ensure_tags, MoneyAmount,
    ValidationError,
};

pub fn validate_purchase(purchase: &PurchaseIr) -> Result<(), ValidationError> {
    ensure_non_empty("purchase_id", &purchase.purchase_id)?;
    ensure_non_empty("source", &purchase.source)?;
    ensure_known_owner("owner", purchase.owner)?;
    ensure_iso_datetime("purchased_at", &purchase.purchased_at)?;
    ensure_non_empty("merchant_name", &purchase.merchant_name)?;
    ensure_non_empty("item_name", &purchase.item_name)?;
    ensure_positive_quantity(purchase.quantity)?;
    purchase.unit_price.validate_non_negative("unit_price")?;
    purchase.total_price.validate_non_negative("total_price")?;
    ensure_matching_currency(&purchase.unit_price, &purchase.total_price)?;
    ensure_expected_total(
        purchase.quantity,
        &purchase.unit_price,
        &purchase.total_price,
    )?;
    ensure_non_empty("raw_ref", &purchase.raw_ref)?;
    ensure_tags("tags", &purchase.tags)?;
    Ok(())
}

fn ensure_positive_quantity(quantity: u32) -> Result<(), ValidationError> {
    if quantity == 0 {
        return Err(ValidationError::new(
            "quantity",
            "quantity must be greater than zero",
        ));
    }
    Ok(())
}

fn ensure_matching_currency(
    unit: &MoneyAmount,
    total: &MoneyAmount,
) -> Result<(), ValidationError> {
    if unit.currency != total.currency {
        return Err(ValidationError::new(
            "total_price",
            "total price currency must match unit price currency",
        ));
    }
    Ok(())
}

fn ensure_expected_total(
    quantity: u32,
    unit: &MoneyAmount,
    total: &MoneyAmount,
) -> Result<(), ValidationError> {
    let expected_total = unit.minor_units * i64::from(quantity);
    if expected_total != total.minor_units {
        return Err(ValidationError::new(
            "total_price",
            "total price must equal unit price times quantity",
        ));
    }
    Ok(())
}
