part of '../main.dart';

class RecommendationScoreBadge extends StatelessWidget {
  const RecommendationScoreBadge({super.key, required this.score});

  final int score;

  @override
  Widget build(BuildContext context) {
    final shad = ShadTheme.of(context);
    return SizedBox(
      width: 48,
      height: 48,
      child: DecoratedBox(
        decoration: BoxDecoration(
          shape: BoxShape.circle,
          color: shad.colorScheme.secondary,
        ),
        child: Center(
          child: Text(
            '$score',
            style: Theme.of(
              context,
            ).textTheme.titleSmall?.copyWith(color: shad.colorScheme.primary),
          ),
        ),
      ),
    );
  }
}

class RecommendationHeader extends StatelessWidget {
  const RecommendationHeader({super.key, required this.recommendation});

  final RecommendationInsight recommendation;

  @override
  Widget build(BuildContext context) {
    final shad = ShadTheme.of(context);
    return Row(
      children: [
        Icon(
          _recommendationIcon(recommendation.kind),
          size: 18,
          color: shad.colorScheme.primary,
        ),
        const SizedBox(width: 6),
        Expanded(
          child: Text(
            recommendation.title,
            style: Theme.of(context).textTheme.titleSmall,
          ),
        ),
      ],
    );
  }
}
