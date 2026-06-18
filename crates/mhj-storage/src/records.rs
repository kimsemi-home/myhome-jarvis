use serde::Deserialize;

#[derive(Debug, Clone, Deserialize)]
pub(crate) struct JsonMoney {
    pub minor_units: i64,
    pub currency: String,
}

#[derive(Debug, Clone, Deserialize)]
pub(crate) struct FinanceJsonRecord {
    pub transaction_id: String,
    pub source: String,
    pub owner: String,
    pub occurred_at: String,
    pub posted_at: Option<String>,
    pub amount: JsonMoney,
    pub direction: String,
    pub merchant_name: Option<String>,
    pub category: Option<String>,
    pub account_id: Option<String>,
    pub card_id: Option<String>,
    pub raw_ref: String,
    pub tags: Vec<String>,
}

#[derive(Debug, Clone, Deserialize)]
pub(crate) struct CommerceJsonRecord {
    pub purchase_id: String,
    pub source: String,
    pub owner: String,
    pub purchased_at: String,
    pub merchant_name: String,
    pub order_id: Option<String>,
    pub item_name: String,
    pub brand: Option<String>,
    pub quantity: u64,
    pub unit_price: JsonMoney,
    pub total_price: JsonMoney,
    pub category: Option<String>,
    pub recurring_candidate: bool,
    pub raw_ref: String,
    pub tags: Vec<String>,
}
