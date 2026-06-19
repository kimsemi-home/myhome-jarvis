import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

import 'metric_value.dart';

void expectRuntimeMetrics(JarvisSnapshot snapshot) {
  expect(snapshot.metrics.map((metric) => metric.value), contains('127.0.0.1'));
  expect(metricValue(snapshot, 'Network'), 'Local-only');
  expect(metricValue(snapshot, 'LAN Auth'), 'Configured');
  expect(snapshot.metrics.map((metric) => metric.value), contains('7'));
  expect(metricValue(snapshot, 'Quality'), 'Passing (3)');
  expect(metricValue(snapshot, 'Public Safety'), 'Clear');
  expect(metricValue(snapshot, 'Code Shape'), '81 debt');
  expect(metricValue(snapshot, 'Events'), '2');
  expect(metricValue(snapshot, 'Runtime'), '9 goroutines');
  expect(metricValue(snapshot, 'Heap'), '1.0 MiB');
  expect(metricValue(snapshot, 'Supervisor'), 'Reachable');
  expect(metricValue(snapshot, 'Command Audit'), '4');
  expect(metricValue(snapshot, 'Planner'), '5/6 done, 1 gated');
  expect(metricValue(snapshot, 'Planner Gate'), 'Linear Sync');
  expect(snapshot.metrics.map((metric) => metric.value), contains('Dirty'));
}
