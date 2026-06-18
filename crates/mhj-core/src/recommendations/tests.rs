use super::*;

#[test]
fn fixture_recommendations_are_ranked_and_bounded() {
    let recommendations = fixture_recommendations().expect("recommendations score");
    assert_eq!(recommendations.len(), 4);
    assert!(recommendations
        .windows(2)
        .all(|window| window[0].score >= window[1].score));
    assert!(recommendations
        .iter()
        .all(|recommendation| recommendation.score <= 100));
    assert!(recommendations
        .iter()
        .any(|recommendation| recommendation.kind == RecommendationKind::SubscriptionReview));
    assert!(recommendations.iter().any(|recommendation| {
        recommendation.kind == RecommendationKind::CardUsageReview
            && recommendation.evidence_count == 2
            && recommendation.estimated_monthly_minor_units == 153_200
    }));
    assert!(recommendations.iter().any(|recommendation| {
        recommendation.kind == RecommendationKind::RecurringPurchaseReview
            && recommendation.evidence_count == 2
    }));
}
