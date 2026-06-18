part of '../daemon_client.dart';

List<SystemMetric> _plannerMetrics(_StatusMaps s) {
  final ready = _int(s.planner['ready_count']) ?? 0;
  final completed = _int(s.planner['completed_count']) ?? 0;
  final blockedExternal = _int(s.planner['blocked_external_write_count']) ?? 0;
  final tasks = _int(s.planner['task_count']) ?? 0;
  final gate = _plannerGate(s.planner['blocked_external_write_tasks']);
  return [
    SystemMetric(
      label: 'Planner',
      value: _plannerProgress(ready, completed, blockedExternal, tasks),
      icon: Icons.schema_outlined,
    ),
    if (gate != null)
      SystemMetric(
        label: 'Planner Gate',
        value: gate,
        icon: Icons.lock_outline,
      ),
  ];
}

String? _plannerGate(Object? tasks) {
  if (tasks is! List<Object?> || tasks.isEmpty) {
    return null;
  }
  final first = tasks.first;
  if (first is! Map<String, Object?>) {
    return null;
  }
  final id = _string(first['id']);
  if (id != null && id.isNotEmpty) {
    return _titleWords(id);
  }
  final title = _string(first['title']);
  if (title != null && title.isNotEmpty) {
    return title;
  }
  return null;
}
