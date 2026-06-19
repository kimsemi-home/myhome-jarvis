part of '../snapshot.dart';

const _sampleRecommendationItems = [
  '81 - Compare recurring purchase: Bottled water 2L x 6',
  '67 - Review card-linked household spend',
  '61 - Review household subscriptions',
  '49 - Keep household cash buffer',
];

const _sampleRecommendations = [
  RecommendationInsight(
    kind: 'recurring_purchase_review',
    title: 'Compare recurring purchase: Bottled water 2L x 6',
    rationale: 'Coupang appears repeatedly in local purchase fixtures.',
    score: 81,
    currency: 'KRW',
    estimatedMonthlyMinorUnits: 11800,
    evidenceCount: 2,
  ),
  RecommendationInsight(
    kind: 'card_usage_review',
    title: 'Review card-linked household spend',
    rationale:
        'Card-linked debit fixtures exist; keep this as a review-only recommendation, not a card action.',
    score: 67,
    currency: 'KRW',
    estimatedMonthlyMinorUnits: 153200,
    evidenceCount: 2,
  ),
  RecommendationInsight(
    kind: 'subscription_review',
    title: 'Review household subscriptions',
    rationale:
        'Subscription-like debit fixtures exist; keep this as a review-only recommendation.',
    score: 61,
    currency: 'KRW',
    estimatedMonthlyMinorUnits: 65900,
    evidenceCount: 1,
  ),
  RecommendationInsight(
    kind: 'cash_buffer',
    title: 'Keep household cash buffer',
    rationale:
        'Fixture cashflow is positive; reserve surplus before recommendations become executable.',
    score: 49,
    currency: 'KRW',
    estimatedMonthlyMinorUnits: 4346800,
    evidenceCount: 3,
  ),
];
