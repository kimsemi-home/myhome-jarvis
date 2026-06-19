use std::path::{Path, PathBuf};

use super::{Compression, DatasetKind, LakeLayer, StorageFormat};

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct DatasetPlan {
    pub layer: LakeLayer,
    pub dataset: DatasetKind,
    pub format: StorageFormat,
    pub compression: Compression,
    pub path: PathBuf,
}

pub fn plan_dataset(root: &Path, layer: LakeLayer, dataset: DatasetKind) -> DatasetPlan {
    let format = format_for_layer(layer);
    let compression = compression_for_format(format);
    let filename = format!("{}.{}", dataset.slug(), format.extension());
    DatasetPlan {
        layer,
        dataset,
        format,
        compression,
        path: root
            .join(layer.as_str())
            .join(dataset.slug())
            .join(filename),
    }
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
