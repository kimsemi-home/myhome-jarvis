use crate::{Owner, ValidationError};

pub(crate) fn ensure_non_empty(field: &'static str, value: &str) -> Result<(), ValidationError> {
    if value.trim().is_empty() {
        return Err(ValidationError::new(field, "must not be empty"));
    }
    Ok(())
}

pub(crate) fn ensure_currency(field: &'static str, value: &str) -> Result<(), ValidationError> {
    if value.len() != 3
        || !value
            .chars()
            .all(|character| character.is_ascii_uppercase())
    {
        return Err(ValidationError::new(
            field,
            "currency must be a three-letter ISO code",
        ));
    }
    Ok(())
}

pub(crate) fn ensure_iso_datetime(field: &'static str, value: &str) -> Result<(), ValidationError> {
    ensure_non_empty(field, value)?;
    let bytes = value.as_bytes();
    let valid_date_prefix = bytes.len() >= 10
        && bytes[0..4].iter().all(u8::is_ascii_digit)
        && bytes[4] == b'-'
        && bytes[5..7].iter().all(u8::is_ascii_digit)
        && bytes[7] == b'-'
        && bytes[8..10].iter().all(u8::is_ascii_digit);
    if !valid_date_prefix {
        return Err(ValidationError::new(
            field,
            "must start with an ISO-8601 date",
        ));
    }
    if bytes.len() > 10 && !value.contains('T') {
        return Err(ValidationError::new(
            field,
            "timestamp values must include 'T'",
        ));
    }
    Ok(())
}

pub(crate) fn ensure_known_owner(field: &'static str, owner: Owner) -> Result<(), ValidationError> {
    if !owner.is_known() {
        return Err(ValidationError::new(field, "owner must be known"));
    }
    Ok(())
}

pub(crate) fn ensure_tags(field: &'static str, tags: &[String]) -> Result<(), ValidationError> {
    for tag in tags {
        let trimmed = tag.trim();
        if trimmed.is_empty() {
            return Err(ValidationError::new(field, "tags must not be empty"));
        }
        if trimmed != tag {
            return Err(ValidationError::new(
                field,
                "tags must not contain outer whitespace",
            ));
        }
    }
    Ok(())
}
