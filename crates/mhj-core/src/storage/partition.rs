use crate::ValidationError;
use std::path::{Path, PathBuf};

use super::{plan_dataset, DatasetKind, LakeLayer};

pub fn partitioned_dataset_path(
    root: &Path,
    layer: LakeLayer,
    dataset: DatasetKind,
    partition: &str,
) -> Result<PathBuf, ValidationError> {
    validate_partition(partition)?;
    let base = plan_dataset(root, layer, dataset);
    let filename = base
        .path
        .file_name()
        .expect("dataset plan always has file name");
    Ok(base.path.with_file_name(partition).join(filename))
}

pub fn validate_partition(partition: &str) -> Result<(), ValidationError> {
    if partition.trim().is_empty() {
        return Err(ValidationError::new("partition", "must not be empty"));
    }
    if partition.starts_with('/') || partition.contains('\\') || partition.contains("..") {
        return Err(ValidationError::new(
            "partition",
            "must be a relative path segment without traversal",
        ));
    }
    if partition
        .split('/')
        .any(|segment| segment.trim().is_empty())
    {
        return Err(ValidationError::new(
            "partition",
            "must not contain empty path segments",
        ));
    }
    Ok(())
}
