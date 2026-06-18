part of '../daemon_client.dart';

List<String> _linearItems(Map<String, Object?> linear) {
  final items = <String>[];
  final mode = _string(linear['mode']);
  if (mode != null && mode.isNotEmpty) {
    items.add('Mode: ${_title(mode)}');
  }
  final synced = _bool(linear['synced']);
  if (synced != null) {
    items.add('Synced: $synced');
  }
  final teams = linear['teams'];
  if (teams is List<Object?>) {
    for (final team in teams.whereType<Map<String, Object?>>()) {
      final name = _string(team['name']);
      if (name != null && name.isNotEmpty) {
        items.add('Team: $name');
      }
    }
  }
  final teamCount = _int(linear['team_count']);
  if (teamCount != null) {
    items.add('Teams: $teamCount');
  }
  final viewerConfigured = _bool(linear['viewer_configured']);
  if (viewerConfigured != null) {
    items.add('Viewer configured: $viewerConfigured');
  }
  final queuePath = _string(linear['queue_path']);
  if (queuePath != null && queuePath.isNotEmpty) {
    items.add('Queue: linear-offline-queue.jsonl');
  }
  return items.isEmpty ? const ['Offline queue'] : items;
}
