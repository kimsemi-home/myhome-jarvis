part of '../daemon_client.dart';

List<SystemMetric> _runtimeMetrics(_StatusMaps s) {
  final requestCount = _int(s.metrics['requests']);
  final eventCount = _int(s.events['count']) ?? _int(s.metrics['event_count']);
  final goroutineCount = _int(s.metrics['goroutine_count']);
  final heapAllocBytes = _int(s.metrics['heap_alloc_bytes']);
  final supervisorStale = _bool(s.supervisor['stale']);
  final supervisorRecorded = _bool(s.supervisor['recorded']) ?? false;
  final auditCount = _int(s.audit['count']);
  return [
    SystemMetric(
      label: 'Requests',
      value: requestCount == null ? '0' : '$requestCount',
      icon: Icons.query_stats_outlined,
    ),
    SystemMetric(
      label: 'Events',
      value: eventCount == null ? '0' : '$eventCount',
      icon: Icons.receipt_long_outlined,
    ),
    if (goroutineCount != null)
      SystemMetric(
        label: 'Runtime',
        value: '$goroutineCount goroutines',
        icon: Icons.memory_outlined,
      ),
    if (heapAllocBytes != null)
      SystemMetric(
        label: 'Heap',
        value: _formatBytes(heapAllocBytes),
        icon: Icons.storage_outlined,
      ),
    SystemMetric(
      label: 'Supervisor',
      value: supervisorRecorded
          ? (supervisorStale == true ? 'Stale' : 'Reachable')
          : 'Unrecorded',
      icon: Icons.memory_outlined,
    ),
    SystemMetric(
      label: 'Command Audit',
      value: auditCount == null ? '0' : '$auditCount',
      icon: Icons.fact_check_outlined,
    ),
    SystemMetric(
      label: 'Repo',
      value: _bool(s.repo['worktree_clean']) == false ? 'Dirty' : 'Clean',
      icon: Icons.account_tree_outlined,
    ),
  ];
}
