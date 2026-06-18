part of '../main.dart';

class OptimizeView extends StatelessWidget {
  const OptimizeView({super.key, required this.recommendations});

  final List<RecommendationInsight> recommendations;

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: ListView.separated(
        padding: const EdgeInsets.all(16),
        itemCount: recommendations.length,
        separatorBuilder: (_, _) => const SizedBox(height: 8),
        itemBuilder: (context, index) {
          final recommendation = recommendations[index];
          return RecommendationTile(recommendation: recommendation);
        },
      ),
    );
  }
}
