use crate::commerce_batch::commerce_record_batch;
use crate::finance_batch::finance_record_batch;
use crate::manifest::plan_dataset;
use crate::parquet_writer::write_zstd_batch;
use crate::{CuratedWriteReport, DatasetKind, LakeLayer, StorageError};
use arrow_array::RecordBatch;
use std::fs;
use std::path::{Path, PathBuf};

pub fn write_curated_parquet_from_jsonl(
    repository_root: &Path,
    lake_root: &Path,
    layer: LakeLayer,
    dataset: DatasetKind,
    jsonl: &str,
) -> Result<CuratedWriteReport, StorageError> {
    reject_raw_layer(layer)?;
    let file = plan_dataset(lake_root, layer, dataset)?;
    let relative_path = PathBuf::from(&file.relative_path);
    let destination = repository_root.join(&relative_path);
    create_parent(&destination)?;
    let batch = curated_batch(dataset, jsonl)?;
    let row_count = batch.num_rows();
    write_zstd_batch(&destination, batch)?;
    Ok(CuratedWriteReport {
        relative_path: file.relative_path,
        layer,
        dataset,
        format: file.format,
        compression: file.compression,
        schema_version: file.schema_version,
        row_count,
    })
}

fn reject_raw_layer(layer: LakeLayer) -> Result<(), StorageError> {
    if matches!(layer, LakeLayer::Raw) {
        return Err(StorageError::InvalidPath {
            field: "layer",
            message: "curated parquet writes require bronze, silver, or gold".to_string(),
        });
    }
    Ok(())
}

fn create_parent(destination: &Path) -> Result<(), StorageError> {
    let parent = destination
        .parent()
        .ok_or_else(|| StorageError::InvalidPath {
            field: "relative_path",
            message: "curated dataset path must have a parent directory".to_string(),
        })?;
    fs::create_dir_all(parent).map_err(|error| StorageError::Io {
        path: parent.to_path_buf(),
        message: error.to_string(),
    })
}

fn curated_batch(dataset: DatasetKind, jsonl: &str) -> Result<RecordBatch, StorageError> {
    match dataset {
        DatasetKind::FinanceTransactions => finance_record_batch(jsonl),
        DatasetKind::CommercePurchases => commerce_record_batch(jsonl),
    }
}
