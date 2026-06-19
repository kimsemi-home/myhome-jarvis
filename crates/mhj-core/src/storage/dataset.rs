use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum DatasetKind {
    FinanceTransactions,
    CommercePurchases,
}

impl DatasetKind {
    pub fn slug(self) -> &'static str {
        match self {
            DatasetKind::FinanceTransactions => "finance_transactions",
            DatasetKind::CommercePurchases => "commerce_purchases",
        }
    }
}
