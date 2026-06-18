use crate::StorageError;
use std::path::{Component, Path};

pub(crate) fn validate_repo_relative_path(
    field: &'static str,
    path: &Path,
) -> Result<(), StorageError> {
    let rendered = path.to_string_lossy();
    if rendered.trim().is_empty() {
        return invalid_path(field, "must not be empty");
    }
    if rendered.contains('\\') {
        return invalid_path(field, "must use forward slashes");
    }
    if path.is_absolute() {
        return invalid_path(field, "must be repo-relative");
    }
    for component in path.components() {
        match component {
            Component::Normal(segment) if !segment.to_string_lossy().trim().is_empty() => {}
            _ => return invalid_path(field, "must not contain traversal or empty segments"),
        }
    }
    Ok(())
}

pub(crate) fn path_to_repo_string(path: &Path) -> String {
    path.to_string_lossy().replace('\\', "/")
}

fn invalid_path<T>(field: &'static str, message: &str) -> Result<T, StorageError> {
    Err(StorageError::InvalidPath {
        field,
        message: message.to_string(),
    })
}
