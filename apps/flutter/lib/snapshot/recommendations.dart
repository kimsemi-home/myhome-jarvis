part of '../snapshot.dart';

@immutable
class RecommendationInsight {
  const RecommendationInsight({
    required this.kind,
    required this.title,
    required this.rationale,
    required this.score,
    required this.currency,
    required this.estimatedMonthlyMinorUnits,
    required this.evidenceCount,
  });

  final String kind;
  final String title;
  final String rationale;
  final int score;
  final String currency;
  final int estimatedMonthlyMinorUnits;
  final int evidenceCount;
}
