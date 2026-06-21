part of '../daemon_client.dart';

Future<JarvisSnapshot> _loadSnapshotFromDaemon(
  DaemonSnapshotClient config,
) async {
  final client = HttpClient()..connectionTimeout = config.timeout;
  try {
    final objects = await _loadDaemonObjects(config, client);
    final commands = await _getArray(config, client, _commandsPath);
    return buildSnapshot(
      health: _statusObject(objects, 'health'),
      assistant: _statusObject(objects, 'assistant'),
      auth: _statusObject(objects, 'auth'),
      commands: commands,
      linear: _statusObject(objects, 'linear'),
      repo: _statusObject(objects, 'repo'),
      security: _statusObject(objects, 'security'),
      codeShape: _statusObject(objects, 'codeShape'),
      domain: _statusObject(objects, 'domain'),
      connectors: _statusObject(objects, 'connectors'),
      agentCluster: _statusObject(objects, 'agentCluster'),
      learning: _statusObject(objects, 'learning'),
      evidence: _statusObject(objects, 'evidence'),
      confidence: _statusObject(objects, 'confidence'),
      translation: _statusObject(objects, 'translation'),
      controlPlane: _statusObject(objects, 'controlPlane'),
      incidents: _statusObject(objects, 'incidents'),
      evidenceQuality: _statusObject(objects, 'evidenceQuality'),
      review: _statusObject(objects, 'review'),
      authority: _statusObject(objects, 'authority'),
      authorityReviewDecision: _statusObject(
        objects,
        'authorityReviewDecision',
      ),
      metrics: _statusObject(objects, 'metrics'),
      events: _statusObject(objects, 'events'),
      supervisor: _statusObject(objects, 'supervisor'),
      audit: _statusObject(objects, 'audit'),
      quality: _statusObject(objects, 'quality'),
      planner: _statusObject(objects, 'planner'),
    );
  } finally {
    client.close(force: true);
  }
}

Future<Map<String, Map<String, Object?>>> _loadDaemonObjects(
  DaemonSnapshotClient config,
  HttpClient client,
) async {
  final objects = <String, Map<String, Object?>>{};
  for (final endpoint in _objectEndpointPaths.entries) {
    objects[endpoint.key] = await _getObject(config, client, endpoint.value);
  }
  return objects;
}

Map<String, Object?> _statusObject(
  Map<String, Map<String, Object?>> objects,
  String key,
) {
  final value = objects[key];
  if (value == null) {
    throw StateError('missing daemon status object: $key');
  }
  return value;
}
