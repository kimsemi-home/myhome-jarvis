use crate::{Owner, ValidationError};

pub(crate) fn ensure_non_empty(field: &'static str, value: &str) -> Result<(), ValidationError> {
    if value.trim().is_empty() {
        return Err(ValidationError::new(field, "must not be empty"));
    }
    Ok(())
}

pub(crate) fn ensure_currency(field: &'static str, value: &str) -> Result<(), ValidationError> {
    let valid = value.len() == 3
        && value
            .chars()
            .all(|character| character.is_ascii_uppercase());
    if !valid {
        return Err(ValidationError::new(
            field,
            "currency must be a three-letter ISO code",
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
