use crate::columns::*;
use crate::picked_columns::*;
use crate::records::FinanceJsonRecord;
use arrow_array::ArrayRef;

pub(crate) fn finance_columns(records: &[FinanceJsonRecord]) -> Vec<ArrayRef> {
    vec![
        string_column(repeat_schema_version(records.len())),
        strings(records, |record| record.transaction_id.as_str()),
        strings(records, |record| record.source.as_str()),
        strings(records, |record| record.owner.as_str()),
        strings(records, |record| record.occurred_at.as_str()),
        optional_strings(records, |record| record.posted_at.as_deref()),
        i64s(records, |record| record.amount.minor_units),
        strings(records, |record| record.amount.currency.as_str()),
        strings(records, |record| record.direction.as_str()),
        optional_strings(records, |record| record.merchant_name.as_deref()),
        optional_strings(records, |record| record.category.as_deref()),
        optional_strings(records, |record| record.account_id.as_deref()),
        optional_strings(records, |record| record.card_id.as_deref()),
        strings(records, |record| record.raw_ref.as_str()),
        u64s(records, |record| record.tags.len() as u64),
    ]
}
