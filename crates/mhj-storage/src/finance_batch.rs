use crate::batch_error::record_batch_error;
use crate::finance_columns::finance_columns;
use crate::finance_schema::finance_schema;
use crate::parse::parse_jsonl;
use crate::records::FinanceJsonRecord;
use crate::validate_finance::validate_finance_record;
use crate::StorageError;
use arrow_array::RecordBatch;

pub(crate) fn finance_record_batch(input: &str) -> Result<RecordBatch, StorageError> {
    let records: Vec<FinanceJsonRecord> = parse_jsonl(input, validate_finance_record)?;
    RecordBatch::try_new(finance_schema(), finance_columns(&records))
        .map_err(|error| record_batch_error("finance_transactions", error))
}
