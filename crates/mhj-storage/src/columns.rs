use crate::SCHEMA_VERSION;
use arrow_array::{ArrayRef, BooleanArray, Int64Array, StringArray, UInt64Array};
use std::sync::Arc;

pub(crate) fn repeat_schema_version(count: usize) -> Vec<&'static str> {
    vec![SCHEMA_VERSION; count]
}

pub(crate) fn string_column(values: Vec<&str>) -> ArrayRef {
    Arc::new(StringArray::from(values))
}

pub(crate) fn optional_string_column(values: Vec<Option<&str>>) -> ArrayRef {
    Arc::new(StringArray::from(values))
}

pub(crate) fn int64_column(values: Vec<i64>) -> ArrayRef {
    Arc::new(Int64Array::from(values))
}

pub(crate) fn uint64_column(values: Vec<u64>) -> ArrayRef {
    Arc::new(UInt64Array::from(values))
}

pub(crate) fn bool_column(values: Vec<bool>) -> ArrayRef {
    Arc::new(BooleanArray::from(values))
}
