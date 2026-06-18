use crate::StorageError;
use parquet::basic::Compression as ParquetCompression;
use parquet::file::metadata::ParquetMetaData;
use std::path::Path;

pub(crate) fn checked_row_count(row_count: i64, path: &Path) -> Result<usize, StorageError> {
    if row_count < 0 {
        return parquet_error(path, "row count must not be negative");
    }
    Ok(row_count as usize)
}

pub(crate) fn ensure_zstd_row_groups(
    metadata: &ParquetMetaData,
    path: &Path,
) -> Result<(), StorageError> {
    if metadata.num_row_groups() == 0 {
        return parquet_error(path, "curated parquet file must contain a row group");
    }
    for row_group_index in 0..metadata.num_row_groups() {
        ensure_row_group_zstd(metadata, row_group_index, path)?;
    }
    Ok(())
}

fn ensure_row_group_zstd(
    metadata: &ParquetMetaData,
    row_group_index: usize,
    path: &Path,
) -> Result<(), StorageError> {
    let row_group = metadata.row_group(row_group_index);
    for column_index in 0..row_group.num_columns() {
        if !matches!(
            row_group.column(column_index).compression(),
            ParquetCompression::ZSTD(_)
        ) {
            return parquet_error(path, "curated parquet columns must use zstd compression");
        }
    }
    Ok(())
}

fn parquet_error<T>(path: &Path, message: &str) -> Result<T, StorageError> {
    Err(StorageError::Parquet {
        path: path.to_path_buf(),
        message: message.to_string(),
    })
}
