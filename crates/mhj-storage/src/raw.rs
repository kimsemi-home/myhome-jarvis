use crate::manifest::plan_dataset;
use crate::{DatasetKind, LakeLayer, StorageError};
use std::fs;
use std::path::{Path, PathBuf};

pub fn write_raw_jsonl(
    repository_root: &Path,
    lake_root: &Path,
    dataset: DatasetKind,
    jsonl: &str,
) -> Result<PathBuf, StorageError> {
    let file = plan_dataset(lake_root, LakeLayer::Raw, dataset)?;
    let relative_path = PathBuf::from(file.relative_path);
    let destination = repository_root.join(&relative_path);
    let parent = destination
        .parent()
        .ok_or_else(|| StorageError::InvalidPath {
            field: "relative_path",
            message: "raw dataset path must have a parent directory".to_string(),
        })?;
    fs::create_dir_all(parent).map_err(|error| StorageError::Io {
        path: parent.to_path_buf(),
        message: error.to_string(),
    })?;
    let mut data = jsonl.as_bytes().to_vec();
    if !data.ends_with(b"\n") {
        data.push(b'\n');
    }
    fs::write(&destination, data).map_err(|error| StorageError::Io {
        path: destination,
        message: error.to_string(),
    })?;
    Ok(relative_path)
}
