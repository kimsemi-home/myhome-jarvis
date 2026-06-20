part of '../daemon_client.dart';

List<SystemMetric> _systemMetrics(_StatusMaps status) {
  return [
    ..._assistantMetrics(status),
    ..._coreMetrics(status),
    ..._runtimeMetrics(status),
    ..._plannerMetrics(status),
    ..._governanceMetrics(status),
    ..._learningMetrics(status),
  ];
}
