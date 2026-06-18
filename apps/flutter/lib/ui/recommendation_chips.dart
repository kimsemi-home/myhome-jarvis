part of '../main.dart';

class RecommendationChips extends StatelessWidget {
  const RecommendationChips({super.key, required this.recommendation});

  final RecommendationInsight recommendation;

  @override
  Widget build(BuildContext context) {
    return Wrap(
      spacing: 8,
      runSpacing: 8,
      children: [
        Chip(
          label: Text(
            _moneyText(
              recommendation.estimatedMonthlyMinorUnits,
              recommendation.currency,
            ),
          ),
        ),
        Chip(label: Text('${recommendation.evidenceCount} evidence')),
        Chip(label: Text(recommendation.kind)),
      ],
    );
  }
}
