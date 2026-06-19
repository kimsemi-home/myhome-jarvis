import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

import 'metric_value.dart';

void expectGovernanceMetrics(JarvisSnapshot snapshot) {
  expect(metricValue(snapshot, 'Connectors'), '2/2 fixture');
  expect(metricValue(snapshot, 'Agent Cluster'), '5 roles gated');
  expect(metricValue(snapshot, 'Authority Gate'), '3 debt');
  expect(metricValue(snapshot, 'Review Capacity'), '1 debt');
  expect(metricValue(snapshot, 'Learning'), '1 open');
  expect(metricValue(snapshot, 'Evidence Graph'), '2 links');
  expect(metricValue(snapshot, 'Confidence'), 'Medium');
  expect(metricValue(snapshot, 'Translation'), '1 open debt');
  expect(metricValue(snapshot, 'Control Plane'), '2 manifests');
  expect(metricValue(snapshot, 'Incidents'), '1 incident debt');
  expect(metricValue(snapshot, 'Evidence Quality'), '2 reassess');
}
