use arrow_schema::{DataType, Field, Schema};
use std::sync::Arc;

pub(crate) fn finance_schema() -> Arc<Schema> {
    Arc::new(Schema::new(vec![
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
    ]))
}
