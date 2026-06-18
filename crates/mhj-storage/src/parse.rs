use crate::StorageError;
use serde::de::DeserializeOwned;

pub(crate) fn parse_jsonl<T>(
    input: &str,
    validate: fn(&T) -> Result<(), (&'static str, String)>,
) -> Result<Vec<T>, StorageError>
where
    T: DeserializeOwned,
{
    let mut items = Vec::new();
    for (index, line) in input.lines().enumerate() {
        let line_number = index + 1;
        let trimmed = line.trim();
        if trimmed.is_empty() {
            continue;
        }
        let item: T = serde_json::from_str(trimmed).map_err(|error| StorageError::Json {
            line: line_number,
            message: error.to_string(),
        })?;
        validate(&item).map_err(|(field, message)| StorageError::Validation {
            line: line_number,
            field,
            message,
        })?;
        items.push(item);
    }
    Ok(items)
}
