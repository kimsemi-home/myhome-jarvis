import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/daemon_client.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

import '../daemon_client_support/server.dart';
import '../daemon_client_support/auth_header_server.dart';

void runAuthHeaderTest() {
  test('sends authorization header when configured', () async {
    final server = await startAuthHeaderCaptureServer('household-token');
    addTearDown(() => server.close(force: true));

    final client = DaemonSnapshotClient(
      baseUri: daemonUri(server),
      authToken: 'household-token',
    );
    final plan = await client.dryRun(
      const HomeCommand(
        name: 'volume-set',
        payload: '{"level":30}',
        icon: IconData(0),
      ),
    );

    expect(plan.name, 'volume_set');
  });
}
