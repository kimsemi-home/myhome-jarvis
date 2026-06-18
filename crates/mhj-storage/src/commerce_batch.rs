use crate::batch_error::record_batch_error;
use crate::commerce_columns::commerce_columns;
use crate::commerce_schema::commerce_schema;
use crate::parse::parse_jsonl;
use crate::records::CommerceJsonRecord;
use crate::validate_commerce::validate_commerce_record;
use crate::StorageError;
use arrow_array::RecordBatch;

pub(crate) fn commerce_record_batch(input: &str) -> Result<RecordBatch, StorageError> {
    let records: Vec<CommerceJsonRecord> = parse_jsonl(input, validate_commerce_record)?;
    RecordBatch::try_new(commerce_schema(), commerce_columns(&records))
        .map_err(|error| record_batch_error("commerce_purchases", error))
}
