use crate::columns::*;
use crate::picked_columns::*;
use crate::records::CommerceJsonRecord;
use arrow_array::ArrayRef;

pub(crate) fn commerce_columns(records: &[CommerceJsonRecord]) -> Vec<ArrayRef> {
    let mut columns = identity_columns(records);
    columns.extend(amount_columns(records));
    columns.extend(detail_columns(records));
    columns
}

fn identity_columns(records: &[CommerceJsonRecord]) -> Vec<ArrayRef> {
    vec![
        string_column(repeat_schema_version(records.len())),
        strings(records, |record| record.purchase_id.as_str()),
        strings(records, |record| record.source.as_str()),
        strings(records, |record| record.owner.as_str()),
        strings(records, |record| record.purchased_at.as_str()),
        strings(records, |record| record.merchant_name.as_str()),
        optional_strings(records, |record| record.order_id.as_deref()),
        strings(records, |record| record.item_name.as_str()),
        optional_strings(records, |record| record.brand.as_deref()),
    ]
}

fn amount_columns(records: &[CommerceJsonRecord]) -> Vec<ArrayRef> {
    vec![
        u64s(records, |record| record.quantity),
        i64s(records, |record| record.unit_price.minor_units),
        i64s(records, |record| record.total_price.minor_units),
        strings(records, |record| record.total_price.currency.as_str()),
    ]
}

fn detail_columns(records: &[CommerceJsonRecord]) -> Vec<ArrayRef> {
    vec![
        optional_strings(records, |record| record.category.as_deref()),
        bools(records, |record| record.recurring_candidate),
        strings(records, |record| record.raw_ref.as_str()),
        u64s(records, |record| record.tags.len() as u64),
    ]
}
