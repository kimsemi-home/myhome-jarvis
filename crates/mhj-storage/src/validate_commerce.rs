use crate::records::CommerceJsonRecord;
use crate::validate_common::{ensure_non_empty, ensure_owner, ensure_positive_amount};

pub(crate) fn validate_commerce_record(
    record: &CommerceJsonRecord,
) -> Result<(), (&'static str, String)> {
    ensure_non_empty("purchase_id", &record.purchase_id)?;
    ensure_non_empty("source", &record.source)?;
    ensure_owner(&record.owner)?;
    ensure_non_empty("purchased_at", &record.purchased_at)?;
    ensure_non_empty("merchant_name", &record.merchant_name)?;
    ensure_non_empty("item_name", &record.item_name)?;
    ensure_positive_amount("unit_price", &record.unit_price)?;
    ensure_positive_amount("total_price", &record.total_price)?;
    ensure_non_empty("raw_ref", &record.raw_ref)?;
    validate_quantity(record)?;
    validate_total(record)
}

fn validate_quantity(record: &CommerceJsonRecord) -> Result<(), (&'static str, String)> {
    if record.quantity == 0 {
        return Err(("quantity", "must be greater than zero".to_string()));
    }
    Ok(())
}

fn validate_total(record: &CommerceJsonRecord) -> Result<(), (&'static str, String)> {
    if record.unit_price.currency != record.total_price.currency {
        return Err(("total_price", "currency must match unit price".to_string()));
    }
    let expected = record.unit_price.minor_units * record.quantity as i64;
    if expected != record.total_price.minor_units {
        return Err((
            "total_price",
            "must equal unit price times quantity".to_string(),
        ));
    }
    Ok(())
}
