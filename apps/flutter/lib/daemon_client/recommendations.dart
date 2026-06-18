part of '../daemon_client.dart';

List<RecommendationInsight> _recommendations(Map<String, Object?> domain) {
  final recommendations = _object(domain['recommendations']);
  if (recommendations == null) {
    return JarvisSnapshot.sample.recommendations;
  }
  final rawItems = recommendations['items'];
  if (rawItems is! List<Object?>) {
    return JarvisSnapshot.sample.recommendations;
  }
  final items = <RecommendationInsight>[];
  for (final item in rawItems.whereType<Map<String, Object?>>()) {
    final insight = _recommendationInsight(item);
    if (insight != null) {
      items.add(insight);
    }
  }
  return items.isEmpty ? JarvisSnapshot.sample.recommendations : items;
}

RecommendationInsight? _recommendationInsight(Map<String, Object?> item) {
  final title = _string(item['title']);
  if (title == null || title.isEmpty) {
    return null;
  }
  return RecommendationInsight(
    kind: _string(item['kind']) ?? 'review',
    title: title,
    rationale: _string(item['rationale']) ?? '',
    score: _int(item['score']) ?? 0,
    currency: _string(item['currency']) ?? '',
    estimatedMonthlyMinorUnits:
        _int(item['estimated_monthly_minor_units']) ?? 0,
    evidenceCount: _int(item['evidence_count']) ?? 0,
  );
}
