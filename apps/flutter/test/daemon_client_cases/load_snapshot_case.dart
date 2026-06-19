import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/daemon_client.dart';

import '../daemon_client_support/cluster_assertions.dart';
import '../daemon_client_support/command_assertions.dart';
import '../daemon_client_support/dashboard_assertions.dart';
import '../daemon_client_support/governance_metrics_assertions.dart';
import '../daemon_client_support/recommendation_assertions.dart';
import '../daemon_client_support/runtime_metrics_assertions.dart';
import '../daemon_client_support/server.dart';
import '../daemon_client_support/storage_assertions.dart';

void runLoadSnapshotTest() {
  test('loads a snapshot and dry-run plan from daemon endpoints', () async {
    final server = await startDaemonFixtureServer();
    addTearDown(() => server.close(force: true));

    final client = DaemonSnapshotClient(baseUri: daemonUri(server));
    final snapshot = await client.load();

    expectRuntimeMetrics(snapshot);
    expectGovernanceMetrics(snapshot);
    expectCommandCatalog(snapshot);
    expectLinearAndStorage(snapshot);
    expectDomainDashboards(snapshot);
    expectRecommendations(snapshot);
    expectClusterModels(snapshot);
    await expectDryRunPlan(client, snapshot);
  });
}
