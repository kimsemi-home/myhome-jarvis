use std::error::Error;
use std::fmt;
use std::path::PathBuf;

#[derive(Debug, Clone, PartialEq, Eq)]
pub enum StorageError {
    InvalidPath {
        field: &'static str,
        message: String,
    },
    Io {
        path: PathBuf,
        message: String,
    },
    Json {
        line: usize,
        message: String,
    },
    Parquet {
        path: PathBuf,
        message: String,
    },
    Validation {
        line: usize,
        field: &'static str,
        message: String,
    },
}

impl fmt::Display for StorageError {
    fn fmt(&self, formatter: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            StorageError::InvalidPath { field, message } => {
                write!(formatter, "{field}: {message}")
            }
            StorageError::Io { path, message } => {
                write!(formatter, "{}: {}", path.display(), message)
            }
            StorageError::Json { line, message } => {
                write!(formatter, "line {line}: invalid json: {message}")
            }
            StorageError::Parquet { path, message } => {
                write!(formatter, "{}: parquet: {}", path.display(), message)
            }
            StorageError::Validation {
                line,
                field,
                message,
            } => write!(formatter, "line {line}: {field}: {message}"),
        }
    }
}

impl Error for StorageError {}
