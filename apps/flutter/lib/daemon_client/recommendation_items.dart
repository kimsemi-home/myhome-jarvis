part of '../daemon_client.dart';

List<String> _recommendationItems(Map<String, Object?> domain) {
  final recommendations = _object(domain['recommendations']);
  if (recommendations == null) {
    return JarvisSnapshot.sample.recommendationItems;
  }
  final rawItems = recommendations['items'];
  if (rawItems is! List<Object?>) {
    return JarvisSnapshot.sample.recommendationItems;
  }
  final items = <String>[];
  for (final item in rawItems.whereType<Map<String, Object?>>()) {
    final title = _string(item['title']);
    if (title == null || title.isEmpty) {
      continue;
    }
    final score = _int(item['score']);
    items.add(score == null ? title : '$score - $title');
  }
  return items.isEmpty ? JarvisSnapshot.sample.recommendationItems : items;
}
