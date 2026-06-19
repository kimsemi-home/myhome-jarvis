import 'dart:convert';
import 'dart:io';

const _fixturePath = 'test/fixtures/daemon_client/endpoints.json';

Map<String, Object?>? _endpointCache;

Future<Object?> daemonFixturePayload(String path) async {
  final fixtures = await _daemonFixtures();
  return fixtures[path];
}

Future<Map<String, Object?>> _daemonFixtures() async {
  final cached = _endpointCache;
  if (cached != null) {
    return cached;
  }
  final decoded = jsonDecode(await File(_fixturePath).readAsString());
  if (decoded is! Map) {
    throw StateError('daemon fixture root must be a JSON object');
  }
  final fixtures = decoded.cast<String, Object?>();
  _endpointCache = fixtures;
  return fixtures;
}
