use crate::{ensure_non_empty, ValidationError};

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
