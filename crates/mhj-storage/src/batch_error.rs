use crate::StorageError;
use std::path::PathBuf;

pub(crate) fn record_batch_error(dataset: &'static str, error: impl ToString) -> StorageError {
    StorageError::Parquet {
        path: PathBuf::from(dataset),
        message: error.to_string(),
    }
}
