use crate::path_rules::{path_to_repo_string, validate_repo_relative_path};
use crate::{
    Compression, DatasetFile, DatasetKind, LakeLayer, LakeManifest, StorageError, StorageFormat,
    DEFAULT_LAKE_ROOT, SCHEMA_VERSION,
};
use std::path::Path;

pub fn default_manifest() -> Result<LakeManifest, StorageError> {
    manifest_for(Path::new(DEFAULT_LAKE_ROOT), &DatasetKind::all())
}

pub fn manifest_for(
    lake_root: &Path,
    datasets: &[DatasetKind],
) -> Result<LakeManifest, StorageError> {
    validate_repo_relative_path("lake_root", lake_root)?;
    let mut files = Vec::new();
    for layer in LakeLayer::all() {
        for dataset in datasets {
            files.push(plan_dataset(lake_root, layer, *dataset)?);
        }
    }
    Ok(LakeManifest {
        lake_root: path_to_repo_string(lake_root),
        schema_version: SCHEMA_VERSION.to_string(),
        files,
    })
}

pub fn plan_dataset(
    lake_root: &Path,
    layer: LakeLayer,
    dataset: DatasetKind,
) -> Result<DatasetFile, StorageError> {
    validate_repo_relative_path("lake_root", lake_root)?;
    let format = format_for_layer(layer);
    let compression = compression_for_format(format);
    let filename = format!("{}.{}", dataset.slug(), format.extension());
    let relative_path = lake_root
        .join(layer.as_str())
        .join(dataset.slug())
        .join(filename);
    validate_repo_relative_path("relative_path", &relative_path)?;
    Ok(DatasetFile {
        layer,
        dataset,
        format,
        compression,
        schema_version: SCHEMA_VERSION.to_string(),
        relative_path: path_to_repo_string(&relative_path),
    })
}

fn format_for_layer(layer: LakeLayer) -> StorageFormat {
    match layer {
        LakeLayer::Raw => StorageFormat::Jsonl,
        LakeLayer::Bronze | LakeLayer::Silver | LakeLayer::Gold => StorageFormat::Parquet,
    }
}

fn compression_for_format(format: StorageFormat) -> Compression {
    match format {
        StorageFormat::Jsonl => Compression::None,
        StorageFormat::Parquet => Compression::Zstd,
    }
}
