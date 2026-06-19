import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

void expectClusterModels(JarvisSnapshot snapshot) {
  expect(snapshot.connectors.map((connector) => connector.key), [
    'mydata',
    'commerce',
  ]);
  expect(snapshot.connectors.first.fixtureMode, isTrue);
  expect(
    snapshot.connectors.first.forbiddenOperations,
    contains('external_api_call'),
  );
  expect(snapshot.agentClusterSignals.map((signal) => signal.key), [
    'evidence_first',
    'authority_gated',
  ]);
  expect(snapshot.agentClusterSignals.first.status, 'active');
}
