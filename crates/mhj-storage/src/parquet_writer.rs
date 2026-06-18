use crate::StorageError;
use arrow_array::RecordBatch;
use parquet::arrow::ArrowWriter;
use parquet::basic::{Compression as ParquetCompression, ZstdLevel};
use parquet::file::properties::WriterProperties;
use std::fs::File;
use std::path::Path;

pub(crate) fn write_zstd_batch(destination: &Path, batch: RecordBatch) -> Result<(), StorageError> {
    let writer_file = File::create(destination).map_err(|error| StorageError::Io {
        path: destination.to_path_buf(),
        message: error.to_string(),
    })?;
    let properties = WriterProperties::builder()
        .set_compression(ParquetCompression::ZSTD(ZstdLevel::default()))
        .build();
    let mut writer =
        ArrowWriter::try_new(writer_file, batch.schema(), Some(properties)).map_err(|error| {
            StorageError::Parquet {
                path: destination.to_path_buf(),
                message: error.to_string(),
            }
        })?;
    writer
        .write(&batch)
        .map_err(|error| StorageError::Parquet {
            path: destination.to_path_buf(),
            message: error.to_string(),
        })?;
    writer.close().map_err(|error| StorageError::Parquet {
        path: destination.to_path_buf(),
        message: error.to_string(),
    })?;
    Ok(())
}
