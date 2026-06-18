part of '../daemon_client.dart';

List<ConnectorReadiness> _connectors(Map<String, Object?> status) {
  final rawConnectors = status['connectors'];
  if (rawConnectors is! List<Object?>) {
    return JarvisSnapshot.sample.connectors;
  }
  final connectors = <ConnectorReadiness>[];
  for (final item in rawConnectors.whereType<Map<String, Object?>>()) {
    final connector = _connector(item);
    if (connector != null) {
      connectors.add(connector);
    }
  }
  return connectors.isEmpty ? JarvisSnapshot.sample.connectors : connectors;
}

ConnectorReadiness? _connector(Map<String, Object?> item) {
  final key = _string(item['key']);
  final label = _string(item['label']);
  if (key == null || key.isEmpty || label == null || label.isEmpty) {
    return null;
  }
  return ConnectorReadiness(
    key: key,
    label: label,
    category: _string(item['category']) ?? '',
    status: _string(item['status']) ?? 'planned',
    fixtureMode: _bool(item['fixture_mode']) ?? true,
    dataClasses: _stringList(item['data_classes']),
    allowedOperations: _stringList(item['allowed_operations']),
    forbiddenOperations: _stringList(item['forbidden_operations']),
    nextStep: _string(item['next_step']) ?? '',
  );
}

List<AgentClusterSignal> _agentClusterSignals(Map<String, Object?> status) {
  final rawSignals = status['signals'];
  if (rawSignals is! List<Object?>) {
    return JarvisSnapshot.sample.agentClusterSignals;
  }
  final signals = <AgentClusterSignal>[];
  for (final item in rawSignals.whereType<Map<String, Object?>>()) {
    final signal = _agentClusterSignal(item);
    if (signal != null) {
      signals.add(signal);
    }
  }
  return signals.isEmpty ? JarvisSnapshot.sample.agentClusterSignals : signals;
}

AgentClusterSignal? _agentClusterSignal(Map<String, Object?> item) {
  final key = _string(item['key']);
  final label = _string(item['label']);
  if (key == null || key.isEmpty || label == null || label.isEmpty) {
    return null;
  }
  return AgentClusterSignal(
    key: key,
    label: label,
    status: _string(item['status']) ?? 'tracked',
    evidence: _string(item['evidence']) ?? '',
  );
}
