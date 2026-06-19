import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/daemon_client.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

void expectCommandCatalog(JarvisSnapshot snapshot) {
  expect(
    snapshot.commands.map((command) => command.name),
    contains('open-youtube'),
  );
  expect(
    snapshot.commands.map((command) => command.name),
    containsAll([
      'open-netflix',
      'open-disney-plus',
      'open-tving',
      'open-wavve',
      'open-coupang-play',
    ]),
  );
  expect(
    snapshot.commands.map((command) => command.payload),
    contains('{"level":30}'),
  );
}

Future<void> expectDryRunPlan(
  DaemonSnapshotClient client,
  JarvisSnapshot snapshot,
) async {
  final command = snapshot.commands.singleWhere(
    (item) => item.name == 'volume-set',
  );
  expect(command.payloadFields, ['level']);
  final plan = await client.dryRun(command);
  expect(plan.name, 'volume_set');
  expect(plan.dryRun, isTrue);
  expect(plan.executeAllowed, isFalse);
  expect(plan.invocations.single.argv, contains('set volume output volume 30'));
}
