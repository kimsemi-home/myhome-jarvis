use arrow_array::{ArrayRef, BooleanArray, Int64Array, RecordBatch, StringArray, UInt64Array};
use arrow_schema::{DataType, Field, Schema};
use parquet::arrow::ArrowWriter;
use parquet::basic::{Compression as ParquetCompression, ZstdLevel};
use parquet::file::properties::WriterProperties;
use parquet::file::reader::{FileReader, SerializedFileReader};
use serde::de::DeserializeOwned;
use serde::{Deserialize, Serialize};
use std::error::Error;
use std::fmt;
use std::fs::{self, File};
use std::path::{Component, Path, PathBuf};
use std::sync::Arc;

pub const DEFAULT_LAKE_ROOT: &str = "data/lake";
pub const SCHEMA_VERSION: &str = "1.0";

#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum LakeLayer {
    Raw,
    Bronze,
    Silver,
    Gold,
}

impl LakeLayer {
    pub fn as_str(self) -> &'static str {
        match self {
            LakeLayer::Raw => "raw",
            LakeLayer::Bronze => "bronze",
            LakeLayer::Silver => "silver",
            LakeLayer::Gold => "gold",
        }
    }

    pub fn all() -> [LakeLayer; 4] {
        [
            LakeLayer::Raw,
            LakeLayer::Bronze,
            LakeLayer::Silver,
            LakeLayer::Gold,
        ]
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum DatasetKind {
    FinanceTransactions,
    CommercePurchases,
}

impl DatasetKind {
    pub fn slug(self) -> &'static str {
        match self {
            DatasetKind::FinanceTransactions => "finance_transactions",
            DatasetKind::CommercePurchases => "commerce_purchases",
        }
    }

    pub fn all() -> [DatasetKind; 2] {
        [
            DatasetKind::FinanceTransactions,
            DatasetKind::CommercePurchases,
        ]
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum StorageFormat {
    Jsonl,
    Parquet,
}

impl StorageFormat {
    pub fn extension(self) -> &'static str {
        match self {
            StorageFormat::Jsonl => "jsonl",
            StorageFormat::Parquet => "parquet",
        }
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum Compression {
    None,
    Zstd,
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct DatasetFile {
    pub layer: LakeLayer,
    pub dataset: DatasetKind,
    pub format: StorageFormat,
    pub compression: Compression,
    pub schema_version: String,
    pub relative_path: String,
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct LakeManifest {
    pub lake_root: String,
    pub schema_version: String,
    pub files: Vec<DatasetFile>,
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct CuratedWriteReport {
    pub relative_path: String,
    pub layer: LakeLayer,
    pub dataset: DatasetKind,
    pub format: StorageFormat,
    pub compression: Compression,
    pub schema_version: String,
    pub row_count: usize,
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct CuratedReadReport {
    pub relative_path: String,
    pub layer: LakeLayer,
    pub dataset: DatasetKind,
    pub format: StorageFormat,
    pub compression: Compression,
    pub schema_version: String,
    pub row_count: usize,
    pub row_group_count: usize,
    pub column_count: usize,
}

#[derive(Debug, Clone, PartialEq, Eq)]
pub enum StorageError {
    InvalidPath {
        field: &'static str,
        message: String,
    },
    Io {
        path: PathBuf,
        message: String,
    },
    Json {
        line: usize,
        message: String,
    },
    Parquet {
        path: PathBuf,
        message: String,
    },
    Validation {
        line: usize,
        field: &'static str,
        message: String,
    },
}

impl fmt::Display for StorageError {
    fn fmt(&self, formatter: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            StorageError::InvalidPath { field, message } => {
                write!(formatter, "{field}: {message}")
            }
            StorageError::Io { path, message } => {
                write!(formatter, "{}: {}", path.display(), message)
            }
            StorageError::Json { line, message } => {
                write!(formatter, "line {line}: invalid json: {message}")
            }
            StorageError::Parquet { path, message } => {
                write!(formatter, "{}: parquet: {}", path.display(), message)
            }
            StorageError::Validation {
                line,
                field,
                message,
            } => write!(formatter, "line {line}: {field}: {message}"),
        }
    }
}

impl Error for StorageError {}

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
        path: destination.clone(),
        message: error.to_string(),
    })?;
    Ok(relative_path)
}

pub fn write_curated_parquet_from_jsonl(
    repository_root: &Path,
    lake_root: &Path,
    layer: LakeLayer,
    dataset: DatasetKind,
    jsonl: &str,
) -> Result<CuratedWriteReport, StorageError> {
    if matches!(layer, LakeLayer::Raw) {
        return Err(StorageError::InvalidPath {
            field: "layer",
            message: "curated parquet writes require bronze, silver, or gold".to_string(),
        });
    }
    let file = plan_dataset(lake_root, layer, dataset)?;
    let relative_path = PathBuf::from(&file.relative_path);
    let destination = repository_root.join(&relative_path);
    let parent = destination
        .parent()
        .ok_or_else(|| StorageError::InvalidPath {
            field: "relative_path",
            message: "curated dataset path must have a parent directory".to_string(),
        })?;
    fs::create_dir_all(parent).map_err(|error| StorageError::Io {
        path: parent.to_path_buf(),
        message: error.to_string(),
    })?;

    let batch = match dataset {
        DatasetKind::FinanceTransactions => finance_record_batch(jsonl)?,
        DatasetKind::CommercePurchases => commerce_record_batch(jsonl)?,
    };
    let row_count = batch.num_rows();
    let writer_file = File::create(&destination).map_err(|error| StorageError::Io {
        path: destination.clone(),
        message: error.to_string(),
    })?;
    let properties = WriterProperties::builder()
        .set_compression(ParquetCompression::ZSTD(ZstdLevel::default()))
        .build();
    let mut writer =
        ArrowWriter::try_new(writer_file, batch.schema(), Some(properties)).map_err(|error| {
            StorageError::Parquet {
                path: destination.clone(),
                message: error.to_string(),
            }
        })?;
    writer
        .write(&batch)
        .map_err(|error| StorageError::Parquet {
            path: destination.clone(),
            message: error.to_string(),
        })?;
    writer.close().map_err(|error| StorageError::Parquet {
        path: destination,
        message: error.to_string(),
    })?;
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

pub fn inspect_curated_parquet(
    repository_root: &Path,
    lake_root: &Path,
    layer: LakeLayer,
    dataset: DatasetKind,
) -> Result<CuratedReadReport, StorageError> {
    if matches!(layer, LakeLayer::Raw) {
        return Err(StorageError::InvalidPath {
            field: "layer",
            message: "curated parquet reads require bronze, silver, or gold".to_string(),
        });
    }
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
    let row_count = metadata.file_metadata().num_rows();
    if row_count < 0 {
        return Err(StorageError::Parquet {
            path: destination.clone(),
            message: "row count must not be negative".to_string(),
        });
    }
    let row_group_count = metadata.num_row_groups();
    if row_group_count == 0 {
        return Err(StorageError::Parquet {
            path: destination.clone(),
            message: "curated parquet file must contain a row group".to_string(),
        });
    }
    let column_count = metadata.file_metadata().schema_descr().num_columns();
    for row_group_index in 0..row_group_count {
        let row_group = metadata.row_group(row_group_index);
        for column_index in 0..row_group.num_columns() {
            if !matches!(
                row_group.column(column_index).compression(),
                ParquetCompression::ZSTD(_)
            ) {
                return Err(StorageError::Parquet {
                    path: destination.clone(),
                    message: "curated parquet columns must use zstd compression".to_string(),
                });
            }
        }
    }
    Ok(CuratedReadReport {
        relative_path: file.relative_path,
        layer,
        dataset,
        format: file.format,
        compression: Compression::Zstd,
        schema_version: file.schema_version,
        row_count: row_count as usize,
        row_group_count,
        column_count,
    })
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

#[derive(Debug, Clone, Deserialize)]
struct JsonMoney {
    minor_units: i64,
    currency: String,
}

#[derive(Debug, Clone, Deserialize)]
struct FinanceJsonRecord {
    transaction_id: String,
    source: String,
    owner: String,
    occurred_at: String,
    posted_at: Option<String>,
    amount: JsonMoney,
    direction: String,
    merchant_name: Option<String>,
    category: Option<String>,
    account_id: Option<String>,
    card_id: Option<String>,
    raw_ref: String,
    tags: Vec<String>,
}

#[derive(Debug, Clone, Deserialize)]
struct CommerceJsonRecord {
    purchase_id: String,
    source: String,
    owner: String,
    purchased_at: String,
    merchant_name: String,
    order_id: Option<String>,
    item_name: String,
    brand: Option<String>,
    quantity: u64,
    unit_price: JsonMoney,
    total_price: JsonMoney,
    category: Option<String>,
    recurring_candidate: bool,
    raw_ref: String,
    tags: Vec<String>,
}

fn finance_record_batch(input: &str) -> Result<RecordBatch, StorageError> {
    let records: Vec<FinanceJsonRecord> = parse_jsonl(input, validate_finance_record)?;
    let schema = Arc::new(Schema::new(vec![
        Field::new("schema_version", DataType::Utf8, false),
        Field::new("transaction_id", DataType::Utf8, false),
        Field::new("source", DataType::Utf8, false),
        Field::new("owner", DataType::Utf8, false),
        Field::new("occurred_at", DataType::Utf8, false),
        Field::new("posted_at", DataType::Utf8, true),
        Field::new("amount_minor_units", DataType::Int64, false),
        Field::new("currency", DataType::Utf8, false),
        Field::new("direction", DataType::Utf8, false),
        Field::new("merchant_name", DataType::Utf8, true),
        Field::new("category", DataType::Utf8, true),
        Field::new("account_id", DataType::Utf8, true),
        Field::new("card_id", DataType::Utf8, true),
        Field::new("raw_ref", DataType::Utf8, false),
        Field::new("tag_count", DataType::UInt64, false),
    ]));
    RecordBatch::try_new(
        schema,
        vec![
            string_column(repeat_schema_version(records.len())),
            string_column(
                records
                    .iter()
                    .map(|record| record.transaction_id.as_str())
                    .collect(),
            ),
            string_column(
                records
                    .iter()
                    .map(|record| record.source.as_str())
                    .collect(),
            ),
            string_column(records.iter().map(|record| record.owner.as_str()).collect()),
            string_column(
                records
                    .iter()
                    .map(|record| record.occurred_at.as_str())
                    .collect(),
            ),
            optional_string_column(
                records
                    .iter()
                    .map(|record| record.posted_at.as_deref())
                    .collect(),
            ),
            int64_column(
                records
                    .iter()
                    .map(|record| record.amount.minor_units)
                    .collect(),
            ),
            string_column(
                records
                    .iter()
                    .map(|record| record.amount.currency.as_str())
                    .collect(),
            ),
            string_column(
                records
                    .iter()
                    .map(|record| record.direction.as_str())
                    .collect(),
            ),
            optional_string_column(
                records
                    .iter()
                    .map(|record| record.merchant_name.as_deref())
                    .collect(),
            ),
            optional_string_column(
                records
                    .iter()
                    .map(|record| record.category.as_deref())
                    .collect(),
            ),
            optional_string_column(
                records
                    .iter()
                    .map(|record| record.account_id.as_deref())
                    .collect(),
            ),
            optional_string_column(
                records
                    .iter()
                    .map(|record| record.card_id.as_deref())
                    .collect(),
            ),
            string_column(
                records
                    .iter()
                    .map(|record| record.raw_ref.as_str())
                    .collect(),
            ),
            uint64_column(
                records
                    .iter()
                    .map(|record| record.tags.len() as u64)
                    .collect(),
            ),
        ],
    )
    .map_err(|error| StorageError::Parquet {
        path: PathBuf::from("finance_transactions"),
        message: error.to_string(),
    })
}

fn commerce_record_batch(input: &str) -> Result<RecordBatch, StorageError> {
    let records: Vec<CommerceJsonRecord> = parse_jsonl(input, validate_commerce_record)?;
    let schema = Arc::new(Schema::new(vec![
        Field::new("schema_version", DataType::Utf8, false),
        Field::new("purchase_id", DataType::Utf8, false),
        Field::new("source", DataType::Utf8, false),
        Field::new("owner", DataType::Utf8, false),
        Field::new("purchased_at", DataType::Utf8, false),
        Field::new("merchant_name", DataType::Utf8, false),
        Field::new("order_id", DataType::Utf8, true),
        Field::new("item_name", DataType::Utf8, false),
        Field::new("brand", DataType::Utf8, true),
        Field::new("quantity", DataType::UInt64, false),
        Field::new("unit_minor_units", DataType::Int64, false),
        Field::new("total_minor_units", DataType::Int64, false),
        Field::new("currency", DataType::Utf8, false),
        Field::new("category", DataType::Utf8, true),
        Field::new("recurring_candidate", DataType::Boolean, false),
        Field::new("raw_ref", DataType::Utf8, false),
        Field::new("tag_count", DataType::UInt64, false),
    ]));
    RecordBatch::try_new(
        schema,
        vec![
            string_column(repeat_schema_version(records.len())),
            string_column(
                records
                    .iter()
                    .map(|record| record.purchase_id.as_str())
                    .collect(),
            ),
            string_column(
                records
                    .iter()
                    .map(|record| record.source.as_str())
                    .collect(),
            ),
            string_column(records.iter().map(|record| record.owner.as_str()).collect()),
            string_column(
                records
                    .iter()
                    .map(|record| record.purchased_at.as_str())
                    .collect(),
            ),
            string_column(
                records
                    .iter()
                    .map(|record| record.merchant_name.as_str())
                    .collect(),
            ),
            optional_string_column(
                records
                    .iter()
                    .map(|record| record.order_id.as_deref())
                    .collect(),
            ),
            string_column(
                records
                    .iter()
                    .map(|record| record.item_name.as_str())
                    .collect(),
            ),
            optional_string_column(
                records
                    .iter()
                    .map(|record| record.brand.as_deref())
                    .collect(),
            ),
            uint64_column(records.iter().map(|record| record.quantity).collect()),
            int64_column(
                records
                    .iter()
                    .map(|record| record.unit_price.minor_units)
                    .collect(),
            ),
            int64_column(
                records
                    .iter()
                    .map(|record| record.total_price.minor_units)
                    .collect(),
            ),
            string_column(
                records
                    .iter()
                    .map(|record| record.total_price.currency.as_str())
                    .collect(),
            ),
            optional_string_column(
                records
                    .iter()
                    .map(|record| record.category.as_deref())
                    .collect(),
            ),
            bool_column(
                records
                    .iter()
                    .map(|record| record.recurring_candidate)
                    .collect(),
            ),
            string_column(
                records
                    .iter()
                    .map(|record| record.raw_ref.as_str())
                    .collect(),
            ),
            uint64_column(
                records
                    .iter()
                    .map(|record| record.tags.len() as u64)
                    .collect(),
            ),
        ],
    )
    .map_err(|error| StorageError::Parquet {
        path: PathBuf::from("commerce_purchases"),
        message: error.to_string(),
    })
}

fn parse_jsonl<T>(
    input: &str,
    validate: fn(&T) -> Result<(), (&'static str, String)>,
) -> Result<Vec<T>, StorageError>
where
    T: DeserializeOwned,
{
    let mut items = Vec::new();
    for (index, line) in input.lines().enumerate() {
        let line_number = index + 1;
        let trimmed = line.trim();
        if trimmed.is_empty() {
            continue;
        }
        let item: T = serde_json::from_str(trimmed).map_err(|error| StorageError::Json {
            line: line_number,
            message: error.to_string(),
        })?;
        validate(&item).map_err(|(field, message)| StorageError::Validation {
            line: line_number,
            field,
            message,
        })?;
        items.push(item);
    }
    Ok(items)
}

fn validate_finance_record(record: &FinanceJsonRecord) -> Result<(), (&'static str, String)> {
    ensure_non_empty("transaction_id", &record.transaction_id)?;
    ensure_non_empty("source", &record.source)?;
    ensure_owner(&record.owner)?;
    ensure_non_empty("occurred_at", &record.occurred_at)?;
    ensure_positive_amount("amount", &record.amount)?;
    ensure_direction(&record.direction)?;
    ensure_non_empty("raw_ref", &record.raw_ref)?;
    if record.direction == "debit"
        && record
            .merchant_name
            .as_deref()
            .unwrap_or_default()
            .trim()
            .is_empty()
    {
        return Err(("merchant_name", "required for debit records".to_string()));
    }
    if record
        .account_id
        .as_deref()
        .unwrap_or_default()
        .trim()
        .is_empty()
        && record
            .card_id
            .as_deref()
            .unwrap_or_default()
            .trim()
            .is_empty()
    {
        return Err((
            "account_id",
            "account_id or card_id is required".to_string(),
        ));
    }
    Ok(())
}

fn validate_commerce_record(record: &CommerceJsonRecord) -> Result<(), (&'static str, String)> {
    ensure_non_empty("purchase_id", &record.purchase_id)?;
    ensure_non_empty("source", &record.source)?;
    ensure_owner(&record.owner)?;
    ensure_non_empty("purchased_at", &record.purchased_at)?;
    ensure_non_empty("merchant_name", &record.merchant_name)?;
    ensure_non_empty("item_name", &record.item_name)?;
    ensure_positive_amount("unit_price", &record.unit_price)?;
    ensure_positive_amount("total_price", &record.total_price)?;
    ensure_non_empty("raw_ref", &record.raw_ref)?;
    if record.quantity == 0 {
        return Err(("quantity", "must be greater than zero".to_string()));
    }
    if record.unit_price.currency != record.total_price.currency {
        return Err(("total_price", "currency must match unit price".to_string()));
    }
    let expected = record.unit_price.minor_units * record.quantity as i64;
    if expected != record.total_price.minor_units {
        return Err((
            "total_price",
            "must equal unit price times quantity".to_string(),
        ));
    }
    Ok(())
}

fn ensure_non_empty(field: &'static str, value: &str) -> Result<(), (&'static str, String)> {
    if value.trim().is_empty() {
        return Err((field, "must not be empty".to_string()));
    }
    Ok(())
}

fn ensure_owner(owner: &str) -> Result<(), (&'static str, String)> {
    match owner {
        "user" | "spouse" | "household" => Ok(()),
        _ => Err(("owner", "must be user, spouse, or household".to_string())),
    }
}

fn ensure_direction(direction: &str) -> Result<(), (&'static str, String)> {
    match direction {
        "debit" | "credit" => Ok(()),
        _ => Err(("direction", "must be debit or credit".to_string())),
    }
}

fn ensure_positive_amount(
    field: &'static str,
    amount: &JsonMoney,
) -> Result<(), (&'static str, String)> {
    if amount.currency.trim().is_empty() {
        return Err((field, "currency must not be empty".to_string()));
    }
    if amount.minor_units <= 0 {
        return Err((field, "minor units must be greater than zero".to_string()));
    }
    Ok(())
}

fn repeat_schema_version(count: usize) -> Vec<&'static str> {
    vec![SCHEMA_VERSION; count]
}

fn string_column(values: Vec<&str>) -> ArrayRef {
    Arc::new(StringArray::from(values))
}

fn optional_string_column(values: Vec<Option<&str>>) -> ArrayRef {
    Arc::new(StringArray::from(values))
}

fn int64_column(values: Vec<i64>) -> ArrayRef {
    Arc::new(Int64Array::from(values))
}

fn uint64_column(values: Vec<u64>) -> ArrayRef {
    Arc::new(UInt64Array::from(values))
}

fn bool_column(values: Vec<bool>) -> ArrayRef {
    Arc::new(BooleanArray::from(values))
}

fn validate_repo_relative_path(field: &'static str, path: &Path) -> Result<(), StorageError> {
    let rendered = path.to_string_lossy();
    if rendered.trim().is_empty() {
        return Err(StorageError::InvalidPath {
            field,
            message: "must not be empty".to_string(),
        });
    }
    if rendered.contains('\\') {
        return Err(StorageError::InvalidPath {
            field,
            message: "must use forward slashes".to_string(),
        });
    }
    if path.is_absolute() {
        return Err(StorageError::InvalidPath {
            field,
            message: "must be repo-relative".to_string(),
        });
    }
    for component in path.components() {
        match component {
            Component::Normal(segment) if !segment.to_string_lossy().trim().is_empty() => {}
            _ => {
                return Err(StorageError::InvalidPath {
                    field,
                    message: "must not contain traversal or empty segments".to_string(),
                });
            }
        }
    }
    Ok(())
}

fn path_to_repo_string(path: &Path) -> String {
    path.to_string_lossy().replace('\\', "/")
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::process;
    use std::time::{SystemTime, UNIX_EPOCH};

    #[test]
    fn default_manifest_covers_each_dataset_and_layer() {
        let manifest = default_manifest().expect("default manifest builds");

        assert_eq!(manifest.lake_root, DEFAULT_LAKE_ROOT);
        assert_eq!(manifest.schema_version, SCHEMA_VERSION);
        assert_eq!(manifest.files.len(), 8);
        assert!(manifest.files.iter().all(|file| {
            file.relative_path.starts_with("data/lake/")
                && !file.relative_path.contains("..")
                && !Path::new(&file.relative_path).is_absolute()
        }));
    }

    #[test]
    fn raw_files_are_jsonl_without_compression() {
        let file = plan_dataset(
            Path::new(DEFAULT_LAKE_ROOT),
            LakeLayer::Raw,
            DatasetKind::FinanceTransactions,
        )
        .expect("raw plan builds");

        assert_eq!(file.format, StorageFormat::Jsonl);
        assert_eq!(file.compression, Compression::None);
        assert_eq!(
            file.relative_path,
            "data/lake/raw/finance_transactions/finance_transactions.jsonl"
        );
    }

    #[test]
    fn curated_files_are_parquet_zstd_plans() {
        let file = plan_dataset(
            Path::new(DEFAULT_LAKE_ROOT),
            LakeLayer::Gold,
            DatasetKind::CommercePurchases,
        )
        .expect("gold plan builds");

        assert_eq!(file.format, StorageFormat::Parquet);
        assert_eq!(file.compression, Compression::Zstd);
        assert_eq!(
            file.relative_path,
            "data/lake/gold/commerce_purchases/commerce_purchases.parquet"
        );
    }

    #[test]
    fn lake_root_rejects_absolute_and_traversal_paths() {
        assert!(default_manifest().is_ok());
        assert!(manifest_for(Path::new("/tmp/lake"), &DatasetKind::all()).is_err());
        assert!(manifest_for(Path::new("../lake"), &DatasetKind::all()).is_err());
    }

    #[test]
    fn partitioned_paths_reject_nested_or_unsafe_segments() {
        assert!(partitioned_dataset_path(
            Path::new(DEFAULT_LAKE_ROOT),
            LakeLayer::Silver,
            DatasetKind::FinanceTransactions,
            "month=2026-06",
        )
        .is_ok());
        assert!(partitioned_dataset_path(
            Path::new(DEFAULT_LAKE_ROOT),
            LakeLayer::Silver,
            DatasetKind::FinanceTransactions,
            "../private",
        )
        .is_err());
        assert!(partitioned_dataset_path(
            Path::new(DEFAULT_LAKE_ROOT),
            LakeLayer::Silver,
            DatasetKind::FinanceTransactions,
            "month=2026-06/day=14",
        )
        .is_err());
        assert!(partitioned_dataset_path(
            Path::new(DEFAULT_LAKE_ROOT),
            LakeLayer::Silver,
            DatasetKind::FinanceTransactions,
            "month\\2026-06",
        )
        .is_err());
    }

    #[test]
    fn raw_jsonl_writer_materializes_only_raw_fixture_files() {
        let root = temp_root();
        let relative = write_raw_jsonl(
            &root,
            Path::new(DEFAULT_LAKE_ROOT),
            DatasetKind::CommercePurchases,
            "{\"purchase_id\":\"fixture-1\"}",
        )
        .expect("raw writer succeeds");

        assert_eq!(
            relative,
            Path::new("data/lake/raw/commerce_purchases/commerce_purchases.jsonl")
        );
        let written = fs::read_to_string(root.join(&relative)).expect("raw file readable");
        assert_eq!(written, "{\"purchase_id\":\"fixture-1\"}\n");
        let _ = fs::remove_dir_all(root);
    }

    #[test]
    fn curated_writer_materializes_finance_parquet_zstd_fixture() {
        let root = temp_root();
        let report = write_curated_parquet_from_jsonl(
            &root,
            Path::new(DEFAULT_LAKE_ROOT),
            LakeLayer::Bronze,
            DatasetKind::FinanceTransactions,
            include_str!("../../../fixtures/finance_transactions.jsonl"),
        )
        .expect("curated finance writer succeeds");

        assert_eq!(
            report.relative_path,
            "data/lake/bronze/finance_transactions/finance_transactions.parquet"
        );
        assert_eq!(report.format, StorageFormat::Parquet);
        assert_eq!(report.compression, Compression::Zstd);
        assert_eq!(report.row_count, 3);
        assert_parquet_file(&root.join(&report.relative_path), 3);
        let read_report = inspect_curated_parquet(
            &root,
            Path::new(DEFAULT_LAKE_ROOT),
            LakeLayer::Bronze,
            DatasetKind::FinanceTransactions,
        )
        .expect("curated finance reader succeeds");
        assert_eq!(read_report.row_count, 3);
        assert_eq!(read_report.row_group_count, 1);
        assert_eq!(read_report.column_count, 15);
        assert_eq!(read_report.compression, Compression::Zstd);
        let _ = fs::remove_dir_all(root);
    }

    #[test]
    fn curated_writer_materializes_commerce_parquet_zstd_fixture() {
        let root = temp_root();
        let report = write_curated_parquet_from_jsonl(
            &root,
            Path::new(DEFAULT_LAKE_ROOT),
            LakeLayer::Gold,
            DatasetKind::CommercePurchases,
            include_str!("../../../fixtures/commerce_purchases.jsonl"),
        )
        .expect("curated commerce writer succeeds");

        assert_eq!(
            report.relative_path,
            "data/lake/gold/commerce_purchases/commerce_purchases.parquet"
        );
        assert_eq!(report.format, StorageFormat::Parquet);
        assert_eq!(report.compression, Compression::Zstd);
        assert_eq!(report.row_count, 3);
        assert_parquet_file(&root.join(&report.relative_path), 3);
        let read_report = inspect_curated_parquet(
            &root,
            Path::new(DEFAULT_LAKE_ROOT),
            LakeLayer::Gold,
            DatasetKind::CommercePurchases,
        )
        .expect("curated commerce reader succeeds");
        assert_eq!(read_report.row_count, 3);
        assert_eq!(read_report.row_group_count, 1);
        assert_eq!(read_report.column_count, 17);
        assert_eq!(read_report.compression, Compression::Zstd);
        let _ = fs::remove_dir_all(root);
    }

    #[test]
    fn curated_writer_rejects_raw_layer() {
        let root = temp_root();
        let result = write_curated_parquet_from_jsonl(
            &root,
            Path::new(DEFAULT_LAKE_ROOT),
            LakeLayer::Raw,
            DatasetKind::FinanceTransactions,
            "{}",
        );

        assert!(matches!(
            result,
            Err(StorageError::InvalidPath { field: "layer", .. })
        ));
        let _ = fs::remove_dir_all(root);
    }

    #[test]
    fn curated_reader_rejects_raw_layer() {
        let root = temp_root();
        let result = inspect_curated_parquet(
            &root,
            Path::new(DEFAULT_LAKE_ROOT),
            LakeLayer::Raw,
            DatasetKind::FinanceTransactions,
        );

        assert!(matches!(
            result,
            Err(StorageError::InvalidPath { field: "layer", .. })
        ));
        let _ = fs::remove_dir_all(root);
    }

    fn assert_parquet_file(path: &Path, expected_rows: i64) {
        let bytes = fs::read(path).expect("parquet file readable");
        assert!(bytes.len() > 8);
        assert_eq!(&bytes[..4], b"PAR1");
        assert_eq!(&bytes[bytes.len() - 4..], b"PAR1");

        let file = File::open(path).expect("parquet file open");
        let reader = SerializedFileReader::new(file).expect("parquet metadata readable");
        let metadata = reader.metadata();
        assert_eq!(metadata.file_metadata().num_rows(), expected_rows);
        assert_eq!(metadata.num_row_groups(), 1);
        assert!(matches!(
            metadata.row_group(0).column(0).compression(),
            ParquetCompression::ZSTD(_)
        ));
    }

    fn temp_root() -> PathBuf {
        let nanos = SystemTime::now()
            .duration_since(UNIX_EPOCH)
            .expect("clock after epoch")
            .as_nanos();
        std::env::temp_dir().join(format!("mhj-storage-{}-{nanos}", process::id()))
    }
}
