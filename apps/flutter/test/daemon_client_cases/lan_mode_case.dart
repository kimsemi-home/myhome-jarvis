import 'package:flutter_test/flutter_test.dart';

import '../daemon_client_support/lan_snapshot.dart';
import '../daemon_client_support/metric_value.dart';

void runLanModeTest() {
  test('marks LAN daemon mode as token-gated', () {
    final snapshot = buildLanSnapshot();

    expect(metricValue(snapshot, 'Network'), 'LAN token-gated');
  });
}
