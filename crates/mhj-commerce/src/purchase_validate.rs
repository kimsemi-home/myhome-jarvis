use crate::{
    ensure_iso_datetime, ensure_known_owner, ensure_non_empty, ensure_tags, PurchaseIr,
    ValidationError,
};

impl PurchaseIr {
    pub fn validate(&self) -> Result<(), ValidationError> {
        ensure_non_empty("purchase_id", &self.purchase_id)?;
        ensure_non_empty("source", &self.source)?;
        ensure_known_owner("owner", self.owner)?;
        ensure_iso_datetime("purchased_at", &self.purchased_at)?;
        ensure_non_empty("merchant_name", &self.merchant_name)?;
        ensure_non_empty("item_name", &self.item_name)?;
        ensure_positive_quantity(self.quantity)?;
        self.unit_price.validate_non_negative("unit_price")?;
        self.total_price.validate_non_negative("total_price")?;
        ensure_price_currency_matches(self)?;
        ensure_total_matches_quantity(self)?;
        ensure_non_empty("raw_ref", &self.raw_ref)?;
        ensure_tags("tags", &self.tags)
    }
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

fn ensure_price_currency_matches(purchase: &PurchaseIr) -> Result<(), ValidationError> {
    if purchase.unit_price.currency != purchase.total_price.currency {
        return Err(ValidationError::new(
            "total_price",
            "total price currency must match unit price currency",
        ));
    }
    Ok(())
}

fn ensure_total_matches_quantity(purchase: &PurchaseIr) -> Result<(), ValidationError> {
    let expected_total = purchase.unit_price.minor_units * i64::from(purchase.quantity);
    if expected_total != purchase.total_price.minor_units {
        return Err(ValidationError::new(
            "total_price",
            "total price must equal unit price times quantity",
        ));
    }
    Ok(())
}
