part of '../daemon_client.dart';

List<SystemMetric> _assistantMetrics(_StatusMaps s) {
  final blocked = _int(s.assistant['blocked_gate_count']) ?? 0;
  final state = _string(s.assistant['compact_state']) ?? 'unknown';
  final action = _string(s.assistant['next_safe_action']) ?? '';
  return [
    SystemMetric(
      label: 'Assistant Center',
      value: blocked > 0 ? '$blocked gated' : _titleWords(state),
      icon: Icons.assistant_direction_outlined,
    ),
    SystemMetric(
      label: 'Next Safe Action',
      value: _actionLabel(action),
      icon: Icons.playlist_add_check_circle_outlined,
    ),
    _assistantCostMetric(s.assistant),
  ];
}

SystemMetric _assistantCostMetric(Map<String, Object?> assistant) {
  final cost = _object(assistant['cost']) ?? const <String, Object?>{};
  final state = _string(cost['budget_state']) ?? 'unknown';
  final total = _int(cost['total_units']) ?? 0;
  return SystemMetric(
    label: 'Codex Cost',
    value: state == 'ok' ? '$total units' : _titleWords(state),
    icon: Icons.toll_outlined,
  );
}

String _actionLabel(String action) {
  switch (action) {
    case 'resolve_authority_gate':
      return 'Resolve authority';
    case 'assign_or_reduce_human_review':
      return 'Review queue';
    case 'review_codex_cost_budget':
      return 'Review cost';
    case 'continue_closed_loop_planning':
      return 'Continue loop';
    default:
      return _titleWords(action);
  }
}
