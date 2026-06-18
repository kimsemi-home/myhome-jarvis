use crate::curated_meta::{checked_row_count, ensure_zstd_row_groups};
use crate::manifest::plan_dataset;
use crate::{Compression, CuratedReadReport, DatasetKind, LakeLayer, StorageError};
use parquet::file::reader::{FileReader, SerializedFileReader};
use std::fs::File;
use std::path::Path;

pub fn inspect_curated_parquet(
    repository_root: &Path,
    lake_root: &Path,
    layer: LakeLayer,
    dataset: DatasetKind,
) -> Result<CuratedReadReport, StorageError> {
    reject_raw_layer(layer)?;
    let file = plan_dataset(lake_root, layer, dataset)?;
    let destination = repository_root.join(&file.relative_path);
    let reader_file = File::open(&destination).map_err(|error| StorageError::Io {
        path: destination.clone(),
        message: error.to_string(),
    })?;
    let reader = SerializedFileReader::new(reader_file).map_err(|error| StorageError::Parquet {
        path: destination.clone(),
        message: error.to_string(),
    })?;
    let metadata = reader.metadata();
    let row_count = checked_row_count(metadata.file_metadata().num_rows(), &destination)?;
    ensure_zstd_row_groups(metadata, &destination)?;
    Ok(CuratedReadReport {
        relative_path: file.relative_path,
        layer,
        dataset,
        format: file.format,
        compression: Compression::Zstd,
        schema_version: file.schema_version,
        row_count,
        row_group_count: metadata.num_row_groups(),
        column_count: metadata.file_metadata().schema_descr().num_columns(),
    })
}

fn reject_raw_layer(layer: LakeLayer) -> Result<(), StorageError> {
    if matches!(layer, LakeLayer::Raw) {
        return Err(StorageError::InvalidPath {
            field: "layer",
            message: "curated parquet reads require bronze, silver, or gold".to_string(),
        });
    }
    Ok(())
}
