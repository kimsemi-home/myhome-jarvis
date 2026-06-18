use crate::columns::{
    bool_column, int64_column, optional_string_column, string_column, uint64_column,
};
use arrow_array::ArrayRef;

pub(crate) fn strings<'a, T, F>(records: &'a [T], pick: F) -> ArrayRef
where
    F: Fn(&'a T) -> &'a str,
{
    string_column(records.iter().map(pick).collect())
}

pub(crate) fn optional_strings<'a, T, F>(records: &'a [T], pick: F) -> ArrayRef
where
    F: Fn(&'a T) -> Option<&'a str>,
{
    optional_string_column(records.iter().map(pick).collect())
}

pub(crate) fn i64s<T, F>(records: &[T], pick: F) -> ArrayRef
where
    F: Fn(&T) -> i64,
{
    int64_column(records.iter().map(pick).collect())
}

pub(crate) fn u64s<T, F>(records: &[T], pick: F) -> ArrayRef
where
    F: Fn(&T) -> u64,
{
    uint64_column(records.iter().map(pick).collect())
}

pub(crate) fn bools<T, F>(records: &[T], pick: F) -> ArrayRef
where
    F: Fn(&T) -> bool,
{
    bool_column(records.iter().map(pick).collect())
}
