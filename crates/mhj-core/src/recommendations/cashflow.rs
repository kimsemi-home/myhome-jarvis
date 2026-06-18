use crate::finance::{summarize_cashflow, TransactionIr};
use crate::ValidationError;

use super::score::clamp_score;
use super::{Recommendation, RecommendationKind};

pub(crate) fn push_cash_buffer(
    recommendations: &mut Vec<Recommendation>,
    transactions: &[TransactionIr],
) -> Result<(), ValidationError> {
    let cashflow = summarize_cashflow(transactions, "KRW")?;
    if cashflow.net_minor_units <= 0 {
        return Ok(());
    }
    recommendations.push(Recommendation {
        kind: RecommendationKind::CashBuffer,
        title: "Keep household cash buffer".to_string(),
        rationale: "Fixture cashflow is positive; reserve surplus before recommendations become executable.".to_string(),
        score: clamp_score(45 + cashflow.net_minor_units / 1_000_000),
        currency: cashflow.currency,
        estimated_monthly_minor_units: cashflow.net_minor_units,
        evidence_count: transactions.len(),
    });
    Ok(())
}
