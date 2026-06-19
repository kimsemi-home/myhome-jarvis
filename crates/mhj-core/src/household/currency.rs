use std::collections::BTreeSet;

pub(super) fn scope_currency(currencies: &BTreeSet<&str>) -> String {
    match currencies.len() {
        0 => "KRW".to_string(),
        1 => currencies
            .iter()
            .next()
            .copied()
            .unwrap_or("KRW")
            .to_string(),
        _ => "mixed".to_string(),
    }
}
