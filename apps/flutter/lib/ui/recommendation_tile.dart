part of '../main.dart';

class RecommendationTile extends StatelessWidget {
  const RecommendationTile({super.key, required this.recommendation});

  final RecommendationInsight recommendation;

  @override
  Widget build(BuildContext context) {
    final colors = Theme.of(context).colorScheme;
    return DecoratedBox(
      decoration: BoxDecoration(
        border: Border.all(color: colors.outlineVariant),
        borderRadius: BorderRadius.circular(8),
      ),
      child: Padding(
        padding: const EdgeInsets.all(14),
        child: Row(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            RecommendationScoreBadge(score: recommendation.score),
            const SizedBox(width: 12),
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  RecommendationHeader(recommendation: recommendation),
                  if (recommendation.rationale.isNotEmpty) ...[
                    const SizedBox(height: 6),
                    Text(
                      recommendation.rationale,
                      style: Theme.of(context).textTheme.bodyMedium,
                    ),
                  ],
                  const SizedBox(height: 8),
                  RecommendationChips(recommendation: recommendation),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
