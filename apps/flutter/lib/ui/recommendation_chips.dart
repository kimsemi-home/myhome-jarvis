part of '../main.dart';

class RecommendationChips extends StatelessWidget {
  const RecommendationChips({super.key, required this.recommendation});

  final RecommendationInsight recommendation;

  @override
  Widget build(BuildContext context) {
    return JarvisBadgeWrap(
      labels: [
        _moneyText(
          recommendation.estimatedMonthlyMinorUnits,
          recommendation.currency,
        ),
        '${recommendation.evidenceCount} evidence',
        recommendation.kind,
      ],
    );
  }
}
