use arrow_schema::{DataType, Field, Schema};
use std::sync::Arc;

pub(crate) fn commerce_schema() -> Arc<Schema> {
    Arc::new(Schema::new(vec![
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
    ]))
}
