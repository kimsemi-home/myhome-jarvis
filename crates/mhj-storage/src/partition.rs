use crate::manifest::plan_dataset;
use crate::path_rules::validate_repo_relative_path;
use crate::{DatasetKind, LakeLayer, StorageError};
use std::path::{Path, PathBuf};

pub fn partitioned_dataset_path(
    lake_root: &Path,
    layer: LakeLayer,
    dataset: DatasetKind,
    partition: &str,
) -> Result<PathBuf, StorageError> {
    validate_partition(partition)?;
    let file = plan_dataset(lake_root, layer, dataset)?;
    let planned_path = PathBuf::from(&file.relative_path);
    let filename = planned_path
        .file_name()
        .ok_or_else(|| StorageError::InvalidPath {
            field: "relative_path",
            message: "dataset path must include a file name".to_string(),
        })?;
    let partitioned = planned_path.with_file_name(partition).join(filename);
    validate_repo_relative_path("partitioned_path", &partitioned)?;
    Ok(partitioned)
}

pub fn validate_partition(partition: &str) -> Result<(), StorageError> {
    let path = Path::new(partition);
    validate_repo_relative_path("partition", path)?;
    if partition.contains('/') {
        return Err(StorageError::InvalidPath {
            field: "partition",
            message: "must be one relative partition segment".to_string(),
        });
    }
    Ok(())
}
