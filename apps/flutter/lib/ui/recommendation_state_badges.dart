part of '../main.dart';

class RecommendationStateBadges extends StatelessWidget {
  const RecommendationStateBadges({super.key, required this.recommendation});

  final RecommendationInsight recommendation;

  @override
  Widget build(BuildContext context) {
    return Wrap(
      spacing: 6,
      runSpacing: 6,
      children: [
        const JarvisBadge('review-only', tone: JarvisBadgeTone.outline),
        JarvisBadge(_scoreState, tone: _scoreTone),
        if (recommendation.evidenceCount > 0)
          const JarvisBadge('evidence-backed', tone: JarvisBadgeTone.primary),
      ],
    );
  }

  String get _scoreState {
    if (recommendation.score >= 80) {
      return 'warning';
    }
    if (recommendation.score >= 60) {
      return 'scored';
    }
    return 'watch';
  }

  JarvisBadgeTone get _scoreTone {
    if (recommendation.score >= 80) {
      return JarvisBadgeTone.secondary;
    }
    if (recommendation.score >= 60) {
      return JarvisBadgeTone.primary;
    }
    return JarvisBadgeTone.outline;
  }
}
