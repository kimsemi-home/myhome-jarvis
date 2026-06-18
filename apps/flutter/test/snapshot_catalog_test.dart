import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

import 'snapshot_generated_helpers.dart';

void main() {
  test('offline fallback connectors are generated catalog entries', () async {
    final generated = await generatedConnectors();
    final generatedKeys = {
      for (final connector in generated) requiredString(connector['key']),
    };
    final fallbackKeys = {
      for (final connector in JarvisSnapshot.sample.connectors) connector.key,
    };

    expect(generatedKeys, containsAll(fallbackKeys));
  });

  test(
    'offline fallback agent cluster signals are generated entries',
    () async {
      final generated = await generatedAgentClusterSignals();
      final generatedKeys = {
        for (final signal in generated) requiredString(signal['key']),
      };
      final fallbackKeys = {
        for (final signal in JarvisSnapshot.sample.agentClusterSignals)
          signal.key,
      };

      expect(generatedKeys, containsAll(fallbackKeys));
    },
  );
}
