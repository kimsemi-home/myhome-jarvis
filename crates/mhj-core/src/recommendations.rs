mod card;
mod cashflow;
mod fixture;
mod model;
mod recurring;
mod score;
mod subscription;

#[cfg(test)]
mod tests;

pub use fixture::fixture_recommendations;
pub use model::{Recommendation, RecommendationKind};
pub use score::score_recommendations;
